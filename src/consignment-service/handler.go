package main

import (
	"context"
	"log"
	pb "github.com/consingment-service/proto/consignment"
	vessPb "github.com/vess-service/proto/vess"
)

//
// 定义微服务
//
type service struct {
	repo IRepository
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
	allConsignments,err := s.repo.GetAll()
	if err!=nil{
		return err
	}
	response.Consignments = allConsignments
	return nil
}
