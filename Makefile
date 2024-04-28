

generate:
	protoc --proto_path=internal\api\proto\ --go_out=internal\api\grpc --go-grpc_out=internal\api\grpc internal\api\proto\system_stats.proto