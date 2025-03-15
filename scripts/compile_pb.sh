#!/usr/bin/env bash

#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
# hello.proto

PROTO_DIR="../proto"
OUTPUT_DIR="../pb"

cd "$PROTO_DIR"

for proto_file in *.proto; do
    if [ -f "$proto_file" ]; then
        echo "Compiling $proto_file..."
        protoc --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" -I="$PROTO_DIR" "$proto_file"
    else
        echo "No .proto files found in directory."
        exit
    fi
done

echo "proto file generate success..."
