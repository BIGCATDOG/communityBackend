package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	pb "github.com/consingment-service/proto/consignment"
	vessPb "github.com/vess-service/proto/vess"
	"log"
	"os"
)

const (
	DEFAULT_HOST = "localhost:27017"
)

func main() {

	// 获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}
	session, err := CreateSession(dbHost)
	// 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	CreateSession("")
	log.Printf("service init.....")
	// 解析命令行参数
	server.Init()
	repo := Repository{mgoSession: session}
	vessCient := vessPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	vessCient.Create(context.Background(), &vessPb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500})
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: &repo, vessClient: vessCient})
	log.Printf("service reg successfully.....")
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("service closed .....")
}
