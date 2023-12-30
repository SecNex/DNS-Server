#!/bin/bash

# Build the project
BUILD_DIR=../build
NAME="secnex-dns-server"

rm -rf $BUILD_DIR

# Create the build directory if it doesn't exist
if [ ! -d "$BUILD_DIR" ]; then
    mkdir -p $BUILD_DIR
    mkdir -p $BUILD_DIR/linux/{amd64,386,arm64,arm}
    mkdir -p $BUILD_DIR/mac/{amd64,386,arm64,arm}
    mkdir -p $BUILD_DIR/windows/{amd64,386,arm64,arm}
fi

# Build go binaries
echo "Building go binaries..."

echo "Building for linux (amd64)..."
FILE_NAME=$NAME-linux_amd64
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/linux/amd64/$FILE_NAME
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for linux (386)..."
FILE_NAME=$NAME-linux_386
GOOS=linux GOARCH=386 go build -o $BUILD_DIR/linux/386/$FILE_NAME
GOOS=linux GOARCH=386 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for linux (arm64)..."
FILE_NAME=$NAME-linux_arm64
GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/linux/arm64/$FILE_NAME
GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for linux (arm)..."
FILE_NAME=$NAME-linux_arm
GOOS=linux GOARCH=arm go build -o $BUILD_DIR/linux/arm/$FILE_NAME
GOOS=linux GOARCH=arm go build -o $BUILD_DIR/$FILE_NAME

echo "Building for mac (amd64)..."
FILE_NAME=$NAME-mac_amd64
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/mac/amd64/$FILE_NAME
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for mac (arm64)..."
FILE_NAME=$NAME-mac_arm64
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/mac/arm64/$FILE_NAME
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for windows (amd64)..."
FILE_NAME=$NAME-windows_amd64.exe
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/windows/amd64/$FILE_NAME
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/$FILE_NAME

echo "Building for windows (386)..."
FILE_NAME=$NAME-windows_386.exe
GOOS=windows GOARCH=386 go build -o $BUILD_DIR/windows/386/$FILE_NAME
GOOS=windows GOARCH=386 go build -o $BUILD_DIR/$FILE_NAME

echo "Build complete!"