module storage-application

go 1.16

replace github.com/spf13/afero => github.com/spf13/afero v1.5.1

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/infobloxopen/atlas-app-toolkit v1.2.0
	github.com/infobloxopen/atlas-pubsub v0.4.0
	github.com/infobloxopen/protoc-gen-gorm v1.1.0
	github.com/jinzhu/gorm v1.9.16
	github.com/prometheus/client_golang v1.12.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.12.0
	google.golang.org/genproto v0.0.0-20220628213854-d9e0b6570c03
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)
