.PHONY: gen.%

path=

gen.grpc:
	protoc --go_out=. --go-grpc_out=. $(path)

