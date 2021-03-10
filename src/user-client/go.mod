module "user-client"

require (
	github.com/user-service/proto/user v0.0.0
)

replace (
	github.com/user-service/proto/user  => ../user-service/proto/user
)