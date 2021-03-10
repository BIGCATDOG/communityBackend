module user-client

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/micro/cli/v2 v2.1.2
	github.com/user-service/proto/user v0.0.0
)

replace github.com/user-service/proto/user => ./proto/user
