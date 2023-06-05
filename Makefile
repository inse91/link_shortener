#proto13:
#	protoc --proto_path=./proto --go_out=./proto .proto
#
#ver:
#	protoc --version
#
#proto_work:
#	protoc --proto_path=api/proto --go_out=. --go-grpc_out=:.
#
#proto_new:
#	protoc -I . .proto --go-grpc_out=:.
#
#proto_from_course:
#	protoc -I api/proto --go_out=plugins=grpc:api/proto api/proto/shortener.proto
#
#proto_from_course1:
#	protoc -I api/proto --go-grpc_out=api/proto api/proto/shortener.proto
#
#protoFOO:
#	protoc --proto_path=~/Users/andrey/Developer/GOLANG/link_shortener \
#	--go-grpc_out=~/Users/andrey/Developer/GOLANG/link_shortener \
#	shortener.proto

common:
	protoc --go_out=./internal/proto --go_opt=paths=source_relative api/shortener.proto

server_interface:
	protoc --go-grpc_out=./internal/proto --go-grpc_opt=paths=source_relative api/shortener.proto

