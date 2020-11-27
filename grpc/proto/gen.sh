if [ ! -f "$GOPATH/bin/protoc-go-inject-tag" ];then go get github.com/favadi/protoc-go-inject-tag; fi

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/pbs --go_out=plugins=grpc:$(pwd)/grpc/pbs $(pwd)/grpc/pbs/product.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pbs/product.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/pbs --go_out=plugins=grpc:$(pwd)/grpc/pbs $(pwd)/grpc/pbs/common.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pbs/common.pb.go

protoc -I . -I=$(pwd)/grpc/imports  --proto_path=$(pwd)/grpc/pbs --go_out=plugins=grpc:$(pwd)/grpc/pbs $(pwd)/grpc/pbs/api.proto
protoc-go-inject-tag -input=$(pwd)/grpc/pbs/api.pb.go