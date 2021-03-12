if [ ! -f "$GOPATH/bin/protoc-go-inject-tag" ];then go get github.com/favadi/protoc-go-inject-tag; fi

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto --go_out=plugins=grpc:$(pwd)/grpc/pb $(pwd)/grpc/proto/product.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/product.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto --go_out=plugins=grpc:$(pwd)/grpc/pb $(pwd)/grpc/proto/common.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/common.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto --go_out=plugins=grpc:$(pwd)/grpc/pb $(pwd)/grpc/proto/api.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/api.pb.go