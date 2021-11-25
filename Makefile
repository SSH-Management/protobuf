CLIENT_PROTOFILES := $(shell find proto/client -iname "*.proto")
SERVER_PROTOFILES := $(shell find proto/server -iname "*.proto")

COMMON_PROTOFILES :=  $(shell find proto/common -iname "*.proto")
GOOGLE_PROOTFILES := $(shell find proto/3rd-party/google -iname "*.proto")


.PHONY: protoc-go
protoc-go:
	@mkdir -p client
	@mkdir -p server
	@mkdir -p common
	@mkdir -p google/errdetails
	@mkdir -p google/tag

	@protoc -I proto/3rd-party/google \
		--go_out=google \
		--go-grpc_out=paths=source_relative:google \
		--go-tag_out=paths=source_relative:google \
		$(GOOGLE_PROOTFILES)
	@rm -rf google/github.com

	@protoc -I proto/common -I proto/3rd-party \
		--go_out=common \
		--go-grpc_out=paths=source_relative:common \
		--go-tag_out=paths=source_relative:common \
		$(COMMON_PROTOFILES)
	@rm -rf common/google.golang.org
	@rm -rf common/github.com

	@protoc -I proto/client -I proto/3rd-party -I proto/common \
		--go_out=client \
		--go-grpc_out=paths=source_relative:client \
		--go-tag_out=paths=source_relative:client \
		$(CLIENT_PROTOFILES)
	@rm -rf client/github.com
	@rm -rf client/google.golang.org

	@protoc -I proto/server -I proto/3rd-party -I proto/common \
		--go_out=server \
		--go-grpc_out=paths=source_relative:server \
		--go-tag_out=paths=source_relative:server \
		$(SERVER_PROTOFILES)
	@rm -rf server/google.golang.org
	@rm -rf server/github.com
