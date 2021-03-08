module consignment-service

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/consingment-service/proto/consignment v0.0.0-incompatible
	github.com/google/uuid v1.1.2 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
)

replace github.com/consingment-service/proto/consignment => ./proto/consignment
