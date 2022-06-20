# Golang-Camp

1. Install Golang
2. Add GOPATH to environment variables
3. Get and install protoc-gen-go, protoc-gen-go-grpc, protoc-gen-grpc-gateway, protoc-gen-openapiv2 and protoc-gen-validate
4. Create .proto file
5. Create server file (main.go)
6. Generates service.pb.go & service.pb.gw.go 
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

    Generates service.pb.gw.go
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/service.proto

    Generates service.swagger.json
	protoc -I . --openapiv2_out . --openapiv2_opt logtostderr=true proto/service.proto

    Generates service,pb.validate.go
	protoc -I . --go_opt=paths=source_relative --go_out=. --validate_out="lang=go:.." proto/service.proto
7. Install Docker
8. Install TablePlus Database GUI 
9. Create docker-compose.yml file
10. Create db.go file
11. Create Dockerfile
12. Get GORM
13. Run ( docker compose up -d )
14. Service and Database in Docker Container starts 
15. Test with Postman
