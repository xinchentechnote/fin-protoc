./generate.sh
make build
make test
./bin/fin-protoc format -f chat/proto/chat.dsl > chat/proto/chat_format.dsl