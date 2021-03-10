package main

import (
	"github.com/asim/go-micro/v3"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	pb "github.com/user-service/proto/user"
)
const (
	DEFAULT_HOST = "localhost:27017"
)
const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	)`;
func main(){


	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	log.Printf("service init.....")
	// 解析命令行参数
	server.Init()

	db, err := NewConnection()
	if err!=nil{
		log.Print(err)
	}
	defer db.Close()
    db.CreateTable(pb.User{})
	repo1 := UserRepository{db: db}


	pb.RegisterUserServiceHandler(server.Server(), &Handler{repo: &repo1})
	log.Printf("service reg successfully.....")
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("service closed .....")
}
