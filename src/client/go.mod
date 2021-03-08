module client

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/consingment-service/proto/consignment v0.0.0-incompatible
)

replace github.com/consingment-service/proto/consignment => ../consignment-service/proto/consignment
