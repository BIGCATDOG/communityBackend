module user-client

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/user-service/proto/user v0.0.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
)

replace github.com/user-service/proto/user => ../user-service/proto/user
