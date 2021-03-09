package main

import (
	"context"
	pb "github.com/vess-service/proto/vess"
)



// 定义货船服务
type service struct {
	repo Repository
}

// 实现服务端
func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	// 调用内部方法查找
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}

func (s *service) Create(ctx context.Context,vessel *pb.Vessel,resp *pb.Response)error  {
	return s.repo.CreatVessel(vessel)
}
