.PHONY: protoc-go
protoc-go: protoc-go-common protoc-go-client protoc-go-server

COMMON_PROTOFILES := $(shell find proto/common -iname "*.proto")
.PHONY: protoc-go-common
protoc-go-common:
	@mkdir -p common

	@protoc -I proto/common -I proto/3rd-party \
		--go_out=common \
		--go-grpc_out=paths=source_relative:common \
		--go-tag_out=paths=source_relative:common \
		$(COMMON_PROTOFILES)
	@rm -rf common/google.golang.org
	@rm -rf common/github.com

SERVER_PROTOFILES := $(shell find proto/server -iname "*.proto")

.PHONY: protoc-go-common
protoc-go-server:
	@mkdir -p server
	@protoc -I proto/server -I proto/3rd-party -I proto/common \
		--go_out=server \
		--go-grpc_out=paths=source_relative:server \
		--go-tag_out=paths=source_relative:server \
		$(SERVER_PROTOFILES)
	@rm -rf server/google.golang.org
	@rm -rf server/github.com

CLIENT_PROTOFILES := $(shell find proto/client -iname "*.proto")

.PHONY: protoc-go-client
protoc-go-client: protoc-go-common
	@mkdir -p client
	@protoc -I proto/client -I proto/3rd-party -I proto/common \
		--go_out=client \
		--go-grpc_out=paths=source_relative:client \
		--go-tag_out=paths=source_relative:client \
		$(CLIENT_PROTOFILES)
	@rm -rf client/github.com
	@rm -rf client/google.golang.org
