if [ ! -f "$GOPATH/bin/protoc-go-inject-tag" ];then go get github.com/favadi/protoc-go-inject-tag; fi

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/product.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/product.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/common.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/common.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/proto/v1 --go_out=plugins=grpc:$(pwd)/grpc/pb/v1 $(pwd)/grpc/proto/v1/api.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pb/v1/api.pb.go

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=$(pwd)/grpc/pb/v1 --go-grpc_opt=paths=source_relative $(pwd)/grpc/proto/v1/api.proto