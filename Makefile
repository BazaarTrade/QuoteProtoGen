gen:
	protoc --go_out=./ --go-grpc_out=./ proto/quote.proto