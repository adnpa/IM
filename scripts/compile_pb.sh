#!/usr/bin/env bash

#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
# hello.proto

PROJECT_ROOT="/home/hz/Code/go/IM"
PROTO_DIR="$PROJECT_ROOT/proto"
OUTPUT_DIR="$PROJECT_ROOT/internal/api"

cd "$PROTO_DIR"
# protoc --proto_path=proto --go_out=../internal/api hello.proto
# --go-grpc_out="$OUTPUT_DIR"
for proto_file in *.proto; do
    if [ -f "$proto_file" ]; then
        echo "Compiling $proto_file..."
        protoc --proto_path="$PROTO_DIR" --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" "$proto_file" 
    else
        echo "No .proto files found in directory."
        exit
    fi
done

echo "proto file generate success..."
