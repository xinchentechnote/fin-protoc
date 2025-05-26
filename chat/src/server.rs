mod message;

use bytes::BytesMut;
use message::SimpleMessage;
use tokio::{io::AsyncReadExt, net::TcpListener};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:8080").await?;
    println!("Server listening on 127.0.0.1:8080");

    loop {
        let (mut socket, addr) = listener.accept().await?;
        println!("Accepted connection from {}", addr);

        tokio::spawn(async move {
            let mut buf = BytesMut::with_capacity(1024);

            loop {
                let mut temp_buf = [0u8; 1024];
                let n = match socket.read(&mut temp_buf).await {
                    Ok(0) => {
                        println!("Connection closed");
                        return;
                    }
                    Ok(n) => n,
                    Err(e) => {
                        eprintln!("Failed to read: {}", e);
                        return;
                    }
                };

                buf.extend_from_slice(&temp_buf[..n]);

                while buf.len() >= 4 {
                    let mut peek_buf = buf.clone();
                    if let Some(msg) = SimpleMessage::from_bytes(&mut peek_buf) {
                        println!("Received message: {:?}", msg);
                        buf = peek_buf;
                    } else {
                        break;
                    }
                }
            }
        });
    }
}
