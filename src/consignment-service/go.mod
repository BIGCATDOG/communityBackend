module consignment-service

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/consingment-service/proto/consignment v0.0.0-incompatible
	github.com/google/uuid v1.1.2 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/vess-service/proto/vess v0.0.0-incompatible
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/consingment-service/proto/consignment => ./proto/consignment

replace github.com/vess-service/proto/vess => ../vess-service/proto/vess
