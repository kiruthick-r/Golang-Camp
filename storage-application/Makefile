#include Makefile.vars
#include Makefile.common

create:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/service.proto
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true pkg/pb/service.proto
	protoc -I . --openapiv2_out . --openapiv2_opt logtostderr=true pkg/pb/service.proto
	protoc -I . --go_opt=paths=source_relative --go_out=. --validate_out="lang=go:.." pkg/pb/service.proto
	protoc --gorm_out=. pkg/pb/service.proto

run:
	go run cmd/server/main.go
