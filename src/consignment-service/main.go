package main

import (
	pb "github.com/consingment-service/proto/consignment"
	vessPb "github.com/vess-service/proto/vess"
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
	vessClient vessPb.VesselServiceClient
}

func (s *service) CreateConsignment(c context.Context, consignment *pb.Consignment, response *pb.Response) error {
	log.Printf("service creat consignment.....")
	vReq := &vessPb.Specification{
		Capacity:  int32(len(consignment.Containers)),
		MaxWeight: consignment.Weight,
	}
	vResp,verr:=s.vessClient.FindAvailable(c,vReq)
	if verr!=nil{
		return verr
	}
	// 货物被承运
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	consignment.VesselId = vResp.Vessel.Id
	respConsignment, err := s.repo.Create(consignment)
	log.Print( respConsignment )
	if err != nil {
		return err
	}
	response.Consignment = respConsignment
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
	vessCient := vessPb.NewVesselServiceClient("go.micro.srv.vessel",server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: repo,vessClient: vessCient})
	log.Printf("service reg successfully.....")
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("service closed .....")
}