.PHONY: all

all:
	protoc  --go_out=. \
		--go_opt=module=github.com/zRich/zFusion \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/zRich/zFusion \
		./protos/common/*.proto

	protoc  --proto_path=. --go_out=. \
		--go_opt=module=github.com/zRich/zFusion \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/zRich/zFusion \
		./protos/peer/*.proto

