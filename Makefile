build:
	go build -o bin/fin-protoc ./
run: build
	./bin/fin-protoc format -f internal/parser/testdata/need_format.dsl
test:
	go test -v ./...