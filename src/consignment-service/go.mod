module consignment-service

go 1.16

require (
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
	github.com/consingment-service/proto/consignment v0.0.0-incompatible
)
replace github.com/consingment-service/proto/consignment => ./proto/consignment