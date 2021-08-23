.PHONY: build go.%

GO_VERSION=1.16.7

include scripts/gen.mk
include scripts/elk.mk
include scripts/build.mk
include scripts/kafka.mk
include scripts/supervisor.mk

go.install:
	wget https://golang.google.cn/dl/go$(GO_VERSION).linux-amd64.tar.gz
	rm -rf /usr/local/go && tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
	source /etc/profile
	@$(MAKE) go.proxy

go.proxy:
	export GO111MODULE=on
    export GOPROXY=https://goproxy.cn

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

# Makefile
## go.install: install golang
## go.proxy: set gomod and goproxy

# build.mk
## build: build all linux
## build.exe: build all windows

