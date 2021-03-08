module client

go 1.16

require (
	github.com/consingment-service/proto/consignment v0.0.0-incompatible
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/consingment-service/proto/consignment => ../consignment-service/proto/consignment
