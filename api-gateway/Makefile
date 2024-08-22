gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/user.proto \
	protos/message.proto

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/http/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)