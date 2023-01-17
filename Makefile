proto_c:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	 proto/calculator.proto

calculator_server:
	cd ./server && go run server.go

calculator_client:
	cd ./server && go run client.go