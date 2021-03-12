package main

import (
	"context"
	 microclient "github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/cmd"

	"log"
	"os"

	pb "github.com/user-service/proto/user"
)






func main() {

	cmd.Init()
	// 创建 user-service 微服务的客户端
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// 暂时将用户信息写死在代码中
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	resp, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}
	log.Println("created: ", resp.User.Id)

	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("call GetAll error: %v", err)
	}
	for _, u := range allResp.Users {
		log.Printf("%v\n", u)
	}

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	log.Println("token: ", authResp.Token)

	// 直接退出即可
	os.Exit(0)
}