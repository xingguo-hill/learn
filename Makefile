pb:
	rm -f x/grpc/helloworld/*.go
	protoc -I=. -I=../googleapis/ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. x/grpc/helloworld/*.proto
	rm -f x/grpc/helloworldgw/*.go
	protoc -I=. -I=../googleapis/ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. x/grpc/helloworldgw/*.proto

pbclean:
	rm -f x/grpc/helloworld/*.go
	rm -f x/grpc/helloworldgw/*.go