generate_pb:
	protoc --proto_path=./proto --go_out ./pb --go_opt=module=github.com/liuxiaobleach/kafka-demo/pb \
	--go-grpc_out=./pb --go-grpc_opt=require_unimplemented_servers=false,module=github.com/liuxiaobleach/kafka-demo/pb \
	--grpc-gateway_out=./pb --grpc-gateway_opt=module=github.com/liuxiaobleach/kafka-demo/pb \
	./proto/*.proto
