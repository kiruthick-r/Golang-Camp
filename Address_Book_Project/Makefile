create:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/service.proto
	protoc -I . --openapiv2_out . --openapiv2_opt logtostderr=true proto/service.proto
	protoc -I . --go_opt=paths=source_relative --go_out=. --validate_out="lang=go:.." proto/service.proto

run:
	go run server/main.go
