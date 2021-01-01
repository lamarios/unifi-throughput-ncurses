arch := $(shell go env GOARCH)
os := $(shell go env GOOS)
version := 1.5

build:
	 CGO_CFLAGS_ALLOW=".*" CGO_LDFLAGS_ALLOW=".*" go mod download
	 go build -ldflags="-s -w -X main.VERSION=$(version)" -o unifi-throughput

package:
	tar -czf unifi-throughput-$(version)-$(os)-$(arch).tar.gz unifi-throughput

install:
	cp unifi-throughput /usr/bin/unifi-throughput