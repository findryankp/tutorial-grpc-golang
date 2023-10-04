# Generate PB from proto
pb:
	protoc --go_out=./model/ --go-grpc_out=./model/ ./proto/*.proto