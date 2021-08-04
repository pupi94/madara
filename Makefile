# install:
# 	go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
# 	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.0
# 	brew install protobuf@3.14.0

v1:
	protoc -I=./grpc/proto -I=./grpc/imports --go_out=./grpc/pb/v1 --go-grpc_out=./grpc/pb/v1 ./grpc/proto/v1/common.proto
	protoc-go-inject-tag -input=./grpc/pb/v1/common.pb.go

	protoc -I=./grpc/proto -I=./grpc/imports --go_out=./grpc/pb/v1 --go-grpc_out=./grpc/pb/v1 ./grpc/proto/v1/product.proto
	protoc-go-inject-tag -input=./grpc/pb/v1/product.pb.go

	protoc -I=./grpc/proto -I=./grpc/imports --go_out=./grpc/pb/v1 --go-grpc_out=./grpc/pb/v1 --go-grpc_opt require_unimplemented_servers=false ./grpc/proto/v1/api.proto
	protoc-go-inject-tag -input=./grpc/pb/v1/api.pb.go
