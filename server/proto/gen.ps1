protoc --go_out=. --go-grpc_out=. ./proto/test.proto
protoc --go_out=../client --go-grpc_out=../client ./proto/test.proto
