module vess-service

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/vess-service/proto/vess v0.0.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
)

replace github.com/vess-service/proto/vess => ./proto/vess
