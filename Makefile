pb:
	rm -f x/grpc/helloworld/*\.pb\.*
	#protoc -I=. -I=../googleapis/ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. x/grpc/helloworld/*.proto
	rm -f x/grpc/helloworldgw/*\.pb\.*
	#protoc -I=. -I=../googleapis/ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. x/grpc/helloworldgw/*.proto

swagger:
	rm -f x/grpc/helloworld/*.swagger.json
	#protoc -I=. -I=../googleapis/  --swagger_out=./ x/grpc/helloworld/helloworld.proto
	rm -f x/grpc/helloworldgw/*.swagger.json
	#protoc -I=. -I=../googleapis/  --swagger_out=./ x/grpc/helloworldgw/helloworld_gw.proto
pbclean:
	rm -f x/grpc/helloworld/*.go
	rm -f x/grpc/helloworldgw/*.go