#!/bin/sh

OUT_DIR="generated"
if [ ! -d "$OUT_DIR" ]; then
  # 如果目录不存在，创建它
  mkdir "$OUT_DIR"
fi

protoc \
    -I $GOPATH/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.0.4 \
    --go_opt=paths=source_relative \
    --go_out=${OUT_DIR} \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=${OUT_DIR} \
    --proto_path=./proto \
    --validate_out=paths=source_relative,lang=go:${OUT_DIR} \
    proto/**/*.proto
    # --go-gin_opt=plugin=service \