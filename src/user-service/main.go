package main

import (
	"github.com/asim/go-micro/v3"
	pb "github.com/user-service/proto/user"
	"log"
)

func main(){
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	log.Printf("service init.....")
	// 解析命令行参数
	server.Init()
	repo1 := UserRepository{}


	pb.RegisterUserServiceHandler(server.Server(), &Handler{repo: &repo1})
	log.Printf("service reg successfully.....")
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("service closed .....")
}
