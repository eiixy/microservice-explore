.PHONY: build go.%

GO_VERSION=1.16.7

go.proxy:
	export GO111MODULE=on
    export GOPROXY=https://goproxy.cn

build:
	go build -o ./bin/ ./...