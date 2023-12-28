#!/bin/bash

cd src

# Build the project
BUILD_DIR=../build

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
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/linux/amd64/ ./...
echo "Building for linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/linux/arm64/ ./...
echo "Building for linux (arm)..."
GOOS=linux GOARCH=arm go build -o $BUILD_DIR/linux/arm/ ./...

echo "Building for mac (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/mac/amd64/ ./...
echo "Building for mac (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/mac/arm64/ ./

echo "Building for windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/windows/amd64/ ./...
echo "Building for windows (386)..."
GOOS=windows GOARCH=386 go build -o $BUILD_DIR/windows/386/ ./...
echo "Building for windows (arm64)..."
GOOS=windows GOARCH=arm64 go build -o $BUILD_DIR/windows/arm64/ ./...
echo "Building for windows (arm)..."
GOOS=windows GOARCH=arm go build -o $BUILD_DIR/windows/arm/ ./...

echo "Build complete!"