#!/bin/sh

VERSION=$1
ARCH=$(go env GOARCH)
OS=$(go env GOOS)

echo "Building unifi-throughput $VERSION  os:$OS arch:$ARCH"



go build -ldflags="-s -w -X main.VERSION=$VERSION" -o unifi-throughput
#Compressing

tar -czf unifi-throughput-$VERSION-$OS-$ARCH.tar.gz unifi-throughput