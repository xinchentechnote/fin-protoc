use bytes::{Buf, BufMut, BytesMut};
use serde::{Deserialize, Serialize};
use serde_json;

#[derive(Debug, Serialize, Deserialize)]
pub struct SimpleMessage {
    pub msg_type: u16,
    pub json_body: serde_json::Value,
}

impl SimpleMessage {
    /// 将结构体编码为二进制协议格式
    pub fn to_bytes(&self) -> BytesMut {
        let json_str = self.json_body.to_string();
        let json_bytes = json_str.as_bytes();
        let body_len = json_bytes.len() as u16;

        let mut buf = BytesMut::with_capacity(4 + json_bytes.len());
        buf.put_u16(self.msg_type);
        buf.put_u16(body_len);
        buf.put(json_bytes);

        buf
    }

    /// 从 Bytes 中解析为结构体
    pub fn from_bytes(mut buf: &mut BytesMut) -> Option<Self> {
        if buf.len() < 4 {
            return None;
        }

        let msg_type = buf.get_u16();
        let body_len = buf.get_u16();

        if buf.len() < body_len as usize {
            return None;
        }

        let json_bytes = buf.split_to(body_len as usize);
        let json_str = String::from_utf8_lossy(&json_bytes);

        let json_body = serde_json::from_str::<serde_json::Value>(&json_str).ok()?;

        Some(Self {
            msg_type,
            json_body,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json::json;

    #[test]
    fn test_serialize_deserialize() {
        let original = SimpleMessage {
            msg_type: 42,
            json_body: json!({
                "user": "alice",
                "score": 95
            }),
        };

        // Serialize to bytes
        let mut bytes = original.to_bytes();
        assert!(bytes.len() >= 4);

        // Deserialize from bytes
        let parsed = SimpleMessage::from_bytes(&mut bytes).expect("Deserialization failed");

        assert_eq!(parsed.msg_type, original.msg_type);
        assert_eq!(parsed.json_body, original.json_body);
    }

    #[test]
    fn test_invalid_input() {
        let mut data = BytesMut::from(&[0x00, 0x01][..]);
        let result = SimpleMessage::from_bytes(&mut data);
        assert!(result.is_none());

        // Correct header but not enough body bytes
        let mut data = BytesMut::from(&[0x00, 0x01, 0x00, 0x05, 0x7b, 0x22][..]); // {"... but incomplete
        let result = SimpleMessage::from_bytes(&mut data);
        assert!(result.is_none());

        // Valid header but invalid JSON string
        let mut buf = BytesMut::new();
        buf.put_u16(1); // msgType
        buf.put_u16(4); // bodyLen
        buf.extend_from_slice(b"bad{"); // malformed JSON
        let result = SimpleMessage::from_bytes(&mut buf);
        assert!(result.is_none());
    }
}
