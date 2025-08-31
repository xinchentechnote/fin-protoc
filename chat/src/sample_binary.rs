// 手动实现 SampleBinary 结构及其所有变体的序列化/反序列化逻辑

use bytes::{Buf, BufMut, Bytes, BytesMut};

pub trait BinaryCodec: Sized {
    fn encode(&self, buf: &mut BytesMut);
    fn decode(buf: &mut Bytes) -> Option<Self>;
}

#[derive(Debug, PartialEq)]
pub struct Logon {
    pub user_name: String,
    pub password: String,
    pub client_id: u64,
    pub heartbeat_interval: u16,
}

#[derive(Debug, PartialEq)]
pub struct Logout {
    pub user_name: String,
    pub client_id: u64,
}

#[derive(Debug, PartialEq)]
pub struct Heartbeat;

#[derive(Debug, PartialEq)]
pub struct SubOrder {
    pub cl_ord_id: String,
    pub price: u64,
    pub qty: u32,
}

#[derive(Debug, PartialEq)]
pub struct RiskControlRequest {
    pub unique_order_id: String,
    pub cl_ord_id: String,
    pub market_id: String,
    pub security_id: String,
    pub side: char,
    pub order_type: char,
    pub price: u64,
    pub qty: u32,
    pub extra_info: Vec<String>,
    pub sub_orders: Vec<SubOrder>,
}

#[derive(Debug, PartialEq)]
pub struct Detail {
    pub rule_name: String,
    pub code: u16,
}

#[derive(Debug, PartialEq)]
pub struct RiskControlResponse {
    pub unique_order_id: String,
    pub status: i32,
    pub msg: String,
    pub detail: Vec<Detail>,
}

#[derive(Debug, PartialEq)]
pub enum SampleBinary {
    Logon(Logon),
    Logout(Logout),
    Heartbeat(Heartbeat),
    RiskControlRequest(RiskControlRequest),
    RiskControlResponse(RiskControlResponse),
}

// Utility Functions
fn put_string(buf: &mut BytesMut, s: &str) {
    buf.put_u16(s.len() as u16);
    buf.extend_from_slice(s.as_bytes());
}

fn get_string(buf: &mut Bytes) -> Option<String> {
    if buf.remaining() < 2 {
        return None;
    }
    let len = buf.get_u16() as usize;
    if buf.remaining() < len {
        return None;
    }
    Some(String::from_utf8(buf.copy_to_bytes(len).to_vec()).ok()?)
}

fn put_char(buf: &mut BytesMut, c: char) {
    buf.put_u8(c as u8);
}

fn get_char(buf: &mut Bytes) -> Option<char> {
    if buf.remaining() < 1 {
        return None;
    }
    Some(buf.get_u8() as char)
}

impl BinaryCodec for Logon {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.user_name);
        put_string(buf, &self.password);
        buf.put_u64(self.client_id);
        buf.put_u16(self.heartbeat_interval);
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        Some(Self {
            user_name: get_string(buf)?,
            password: get_string(buf)?,
            client_id: buf.get_u64(),
            heartbeat_interval: buf.get_u16(),
        })
    }
}

impl BinaryCodec for Logout {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.user_name);
        buf.put_u64(self.client_id);
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        Some(Self {
            user_name: get_string(buf)?,
            client_id: buf.get_u64(),
        })
    }
}

impl BinaryCodec for Heartbeat {
    fn encode(&self, _buf: &mut BytesMut) {}
    fn decode(_buf: &mut Bytes) -> Option<Self> {
        Some(Heartbeat)
    }
}

impl BinaryCodec for SubOrder {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.cl_ord_id);
        buf.put_u64(self.price);
        buf.put_u32(self.qty);
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        Some(Self {
            cl_ord_id: get_string(buf)?,
            price: buf.get_u64(),
            qty: buf.get_u32(),
        })
    }
}

impl BinaryCodec for RiskControlRequest {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.unique_order_id);
        put_string(buf, &self.cl_ord_id);
        put_string(buf, &self.market_id);
        put_string(buf, &self.security_id);
        put_char(buf, self.side);
        put_char(buf, self.order_type);
        buf.put_u64(self.price);
        buf.put_u32(self.qty);

        buf.put_u16(self.extra_info.len() as u16);
        for s in &self.extra_info {
            put_string(buf, s);
        }

        buf.put_u16(self.sub_orders.len() as u16);
        for s in &self.sub_orders {
            s.encode(buf);
        }
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        let unique_order_id = get_string(buf)?;
        let cl_ord_id = get_string(buf)?;
        let market_id = get_string(buf)?;
        let security_id = get_string(buf)?;
        let side = get_char(buf)?;
        let order_type = get_char(buf)?;
        let price = buf.get_u64();
        let qty = buf.get_u32();

        let n = buf.get_u16();
        let mut extra_info = Vec::with_capacity(n as usize);
        for _ in 0..n {
            extra_info.push(get_string(buf)?);
        }

        let m = buf.get_u16();
        let mut sub_orders = Vec::with_capacity(m as usize);
        for _ in 0..m {
            sub_orders.push(SubOrder::decode(buf)?);
        }

        Some(Self {
            unique_order_id,
            cl_ord_id,
            market_id,
            security_id,
            side,
            order_type,
            price,
            qty,
            extra_info,
            sub_orders,
        })
    }
}

impl BinaryCodec for Detail {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.rule_name);
        buf.put_u16(self.code);
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        Some(Self {
            rule_name: get_string(buf)?,
            code: buf.get_u16(),
        })
    }
}

impl BinaryCodec for RiskControlResponse {
    fn encode(&self, buf: &mut BytesMut) {
        put_string(buf, &self.unique_order_id);
        buf.put_i32(self.status);
        put_string(buf, &self.msg);
        buf.put_u16(self.detail.len() as u16);
        for d in &self.detail {
            d.encode(buf);
        }
    }

    fn decode(buf: &mut Bytes) -> Option<Self> {
        let unique_order_id = get_string(buf)?;
        let status = buf.get_i32();
        let msg = get_string(buf)?;

        let n = buf.get_u16();
        let mut detail = Vec::with_capacity(n as usize);
        for _ in 0..n {
            detail.push(Detail::decode(buf)?);
        }

        Some(Self {
            unique_order_id,
            status,
            msg,
            detail,
        })
    }
}

impl SampleBinary {
    pub fn to_bytes(&self) -> BytesMut {
        let mut buf = BytesMut::new();
        match self {
            SampleBinary::Logon(inner) => {
                buf.put_u16(1);
                let mut body = BytesMut::new();
                inner.encode(&mut body);
                buf.put_u16(body.len() as u16);
                buf.extend_from_slice(&body);
            }
            SampleBinary::Logout(inner) => {
                buf.put_u16(2);
                let mut body = BytesMut::new();
                inner.encode(&mut body);
                buf.put_u16(body.len() as u16);
                buf.extend_from_slice(&body);
            }
            SampleBinary::Heartbeat(inner) => {
                buf.put_u16(3);
                let mut body = BytesMut::new();
                inner.encode(&mut body);
                buf.put_u16(body.len() as u16);
                buf.extend_from_slice(&body);
            }
            SampleBinary::RiskControlRequest(inner) => {
                buf.put_u16(4);
                let mut body = BytesMut::new();
                inner.encode(&mut body);
                buf.put_u16(body.len() as u16);
                buf.extend_from_slice(&body);
            }
            SampleBinary::RiskControlResponse(inner) => {
                buf.put_u16(5);
                let mut body = BytesMut::new();
                inner.encode(&mut body);
                buf.put_u16(body.len() as u16);
                buf.extend_from_slice(&body);
            }
        }
        buf
    }

    pub fn from_bytes(mut buf: &mut Bytes) -> Option<Self> {
        if buf.remaining() < 4 {
            return None;
        }
        let msg_type = buf.get_u16();
        let len = buf.get_u16() as usize;
        if buf.remaining() < len {
            return None;
        }
        let mut body = buf.split_to(len);

        match msg_type {
            1 => Some(SampleBinary::Logon(Logon::decode(&mut body)?)),
            2 => Some(SampleBinary::Logout(Logout::decode(&mut body)?)),
            3 => Some(SampleBinary::Heartbeat(Heartbeat::decode(&mut body)?)),
            4 => Some(SampleBinary::RiskControlRequest(
                RiskControlRequest::decode(&mut body)?,
            )),
            5 => Some(SampleBinary::RiskControlResponse(
                RiskControlResponse::decode(&mut body)?,
            )),
            _ => None,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use bytes::BytesMut;

    fn roundtrip(msg: SampleBinary) {
        let mut buf = msg.to_bytes();
        let recovered =
            SampleBinary::from_bytes(&mut buf.freeze()).expect("Deserialization failed");
        assert_eq!(recovered, msg);
    }

    #[test]
    fn test_logon() {
        let msg = SampleBinary::Logon(Logon {
            user_name: "alice".into(),
            password: "secret".into(),
            client_id: 123456789,
            heartbeat_interval: 30,
        });
        roundtrip(msg);
    }

    #[test]
    fn test_logout() {
        let msg = SampleBinary::Logout(Logout {
            user_name: "bob".into(),
            client_id: 987654321,
        });
        roundtrip(msg);
    }

    #[test]
    fn test_heartbeat() {
        let msg = SampleBinary::Heartbeat(Heartbeat);
        roundtrip(msg);
    }

    #[test]
    fn test_risk_control_request() {
        let msg = SampleBinary::RiskControlRequest(RiskControlRequest {
            unique_order_id: "OID123".into(),
            cl_ord_id: "CLORD001".into(),
            market_id: "SH1".into(),
            security_id: "600519.SH".into(),
            side: '1',
            order_type: '2',
            price: 123456,
            qty: 1000,
            extra_info: vec!["note1".into(), "note2".into()],
            sub_orders: vec![
                SubOrder {
                    cl_ord_id: "sub001".into(),
                    price: 123450,
                    qty: 500,
                },
                SubOrder {
                    cl_ord_id: "sub002".into(),
                    price: 123460,
                    qty: 600,
                },
            ],
        });
        roundtrip(msg);
    }

    #[test]
    fn test_risk_control_response() {
        let msg = SampleBinary::RiskControlResponse(RiskControlResponse {
            unique_order_id: "OID456".into(),
            status: 0,
            msg: "OK".into(),
            detail: vec![
                Detail {
                    rule_name: "price_limit".into(),
                    code: 1001,
                },
                Detail {
                    rule_name: "qty_limit".into(),
                    code: 1002,
                },
            ],
        });
        roundtrip(msg);
    }

    #[test]
    fn test_invalid_type() {
        let mut buf = BytesMut::new();
        buf.put_u16(99);
        buf.put_u16(0);
        assert!(SampleBinary::from_bytes(&mut buf.freeze()).is_none());
    }

    #[test]
    fn test_incomplete_header() {
        let mut buf = BytesMut::from(&[0x00u8][..]);
        assert!(SampleBinary::from_bytes(&mut buf.freeze()).is_none());
    }

    #[test]
    fn test_incomplete_body() {
        let mut buf = BytesMut::new();
        buf.put_u16(1);
        buf.put_u16(10);
        buf.extend_from_slice(&[1, 2, 3]);
        assert!(SampleBinary::from_bytes(&mut buf.freeze()).is_none());
    }
}
