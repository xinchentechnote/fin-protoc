./generate.sh
make build
./bin/fin-protoc format -f chat/proto/chat.dsl > chat/proto/chat_format.dsl