package main

import (
	"github.com/asim/go-micro/v3"
	pb "github.com/vess-service/proto/vess"
	"log"
	"os"
)

const (
	DEFAULT_HOST = "localhost:27017"
)

func main() {

	// 获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == ""{
		dbHost = DEFAULT_HOST
	}
	session, err := CreateSession(dbHost)
	// 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	repo := &VesselRepository{mgoSession: session}
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	// 将实现服务端的 API 注册到服务端
	pb.RegisterVesselServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
