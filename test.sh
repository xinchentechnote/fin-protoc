./generate.sh
make build
./bin/fin-protoc format -f internal/parser/testdata/test_packet.dsl 