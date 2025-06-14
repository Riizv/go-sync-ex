#!/bin/bash

OUTPUT_DIR="build"
OUTPUT_BIN="$OUTPUT_DIR/myapp"
SOURCE_DIR="./cmd/sysinfo"

if [ ! -d "$OUTPUT_DIR" ]; then
    echo "Creating directory $OUTPUT_DIR..."
    mkdir -p "$OUTPUT_DIR"
fi

echo "Building go application..."
go build -o "$OUTPUT_BIN" "$SOURCE_DIR"

if [ $? -eq 0 ]; then
    echo "Application was built as $OUTPUT_BIN"
else
    echo "Error durning building application"
    exit 1
fi