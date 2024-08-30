#!/bin/bash

# Directories
PROTO_DIR=./api/proto
OUT_DIR=./api/proto
THIRD_PARTY_DIR=./third_party

echo "Generating Go files from protobuf definitions..."

protoc --proto_path=$PROTO_DIR --proto_path=$THIRD_PARTY_DIR \
  --go_out=$OUT_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=$OUT_DIR --grpc-gateway_opt=paths=source_relative \
  $PROTO_DIR/event/*.proto \
  $PROTO_DIR/booking/*.proto \
  $PROTO_DIR/gateway/*.proto

echo "Protobuf generation completed!"
