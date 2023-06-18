common:
	protoc --go_out=./internal/proto --go_opt=paths=source_relative api/shortener.proto

server_interface:
	protoc --go-grpc_out=./internal/proto --go-grpc_opt=paths=source_relative api/shortener.proto

swagger_gen:
	swag init -g ../app/app.go  -d ./,../internal/handler/,../internal/model/


