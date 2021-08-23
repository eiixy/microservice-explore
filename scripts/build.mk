.PHONY: build build.% grpcui grpcui.%
ELK_VERSION=7.12.0

GRPCUI_PROTO=
GRPCUI_SERVER=localhost:9000

build:

build.exe:


# grpc ui
grpcui:
	grpcui -plaintext -proto $(GRPCUI_PROTO) $(GRPCUI_SERVER)

grpcui.account:
	grpcui -plaintext -proto ./account/api/account/v1/account.proto 127.0.0.1:9000

grpcui.install:
	go get github.com/fullstorydev/grpcui/...
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
