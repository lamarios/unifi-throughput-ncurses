build:

	 CGO_CFLAGS_ALLOW=".*" CGO_LDFLAGS_ALLOW=".*" go build -ldflags="-s -w -X main.VERSION=$VERSION" -o unifi-throughput

install:
	cp unifi-throughput /usr/bin/unifi-throughput