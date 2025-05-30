build:
	go build -o bin/fin-protoc ./cmd/
	GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -o lib/libpacketdsl.so ./cmd/
run: build
	./bin/fin-protoc format -f internal/parser/testdata/need_format.dsl
test:
	go test -v ./...