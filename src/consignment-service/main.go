package main

import (
	pb "github.com/consingment-service/proto/consignment"
	"context"
	"log"
	"github.com/asim/go-micro/v3"
)

//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment                                   // 获取仓库中所有的货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//
// 定义微服务
//
type service struct {
	repo Repository
}

func (s *service) CreateConsignment(c context.Context, consignment *pb.Consignment, response *pb.Response) error {
	log.Printf("service creat consignment.....")
	consignment, err := s.repo.Create(consignment)
	log.Print( consignment )
	if err != nil {
		return err
	}
	response.Consignment = consignment
	response.Created = true
	return nil
}

func (s *service) GetAllConsignment(c context.Context, request *pb.Request, response *pb.Response) error {
	log.Printf("service getall.....")
	allConsignments := s.repo.GetAll()
	response.Consignments = allConsignments
	return nil
}



func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
    log.Printf("service init.....")
	// 解析命令行参数
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})
	log.Printf("service reg successfully.....")
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("service closed .....")
}