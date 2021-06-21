if [ ! -f "$GOPATH/bin/protoc-go-inject-tag" ];then go get github.com/favadi/protoc-go-inject-tag; fi

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/product.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/product.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/common.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/common.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/api.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/api.pb.go
