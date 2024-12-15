#!/usr/bin/env bash

#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
# hello.proto

PROTO_DIR="../proto"
OUTPUT_DIR="../pb"

cd "$PROTO_DIR"

# 循环遍历目录中的所有 .proto 文件
for proto_file in *.proto; do
    # 检查文件是否存在
    if [ -f "$proto_file" ]; then
        echo "Compiling $proto_file..."
        protoc --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" -I="$PROTO_DIR" "$proto_file"
    else
        echo "No .proto files found in directory."
        exit
    fi
done

echo "proto file generate success..."
