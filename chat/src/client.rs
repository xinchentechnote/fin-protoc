mod message;

use message::SimpleMessage;
use serde_json::json;
use tokio::{
    io::{AsyncReadExt, AsyncWriteExt},
    net::TcpStream,
};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let mut stream = TcpStream::connect("127.0.0.1:8080").await?;
    println!("Connected to server");

    let msg = SimpleMessage {
        msg_type: 1,
        json_body: json!({
            "hello": "world",
            "count": 42
        }),
    };

    let buf = msg.to_bytes();
    stream.write_all(&buf).await?;
    println!("Sent message");

    // 持续读取服务器响应
    let mut buf = vec![0; 1024];
    loop {
        match stream.read_buf(&mut buf).await {
            Ok(0) => {
                println!("Connection closed by server");
                break;
            }
            Ok(n) => {
                println!("Received {} bytes: {:?}", n, &buf[..n]);
            }
            Err(e) => {
                eprintln!("Failed to read from socket: {}", e);
                break;
            }
        }
    }

    Ok(())
}
