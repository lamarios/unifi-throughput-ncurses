#!/bin/sh

VERSION=$1
ARCH=$(go env GOARCH)
OS=$(go env GOOS)


echo "Downloading dependencies"
export CGO_CFLAGS_ALLOW=".*"
export CGO_LDFLAGS_ALLOW=".*"
go mod download

echo "Building unifi-throughput $VERSION  os:$OS arch:$ARCH"
go clean
go build -ldflags="-s -w -X main.VERSION=$VERSION" -o unifi-throughput
#Compressing

tar -czf unifi-throughput-$VERSION-$OS-$ARCH.tar.gz unifi-throughput