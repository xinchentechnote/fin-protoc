build:
	go build -o bin/fin-protoc ./cmd/
# run: build
# 	./bin/gt-auto --casePath pkg/testcase/testdata/test_case.csv --config pkg/config/testdata/gw-auto.toml
test:
	go test -v ./...