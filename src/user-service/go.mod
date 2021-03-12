module user-service

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/user-service/proto/user v0.0.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace github.com/user-service/proto/user => ./proto/user
