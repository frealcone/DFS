default: build

build: protoc
	go build -o $(GOPATH)/bin/DFS/ ./cmd/*

run: build
	$(GOPATH)/bin/DFS/*

test:
	go test ./test/...

protoc:
	protoc --proto_path="./proto" --go_out="./" --go-grpc_out="./" ./proto/*.proto