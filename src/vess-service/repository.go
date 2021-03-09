package main

import (
	"errors"
	pb "github.com/vess-service/proto/vess"
	"gopkg.in/mgo.v2"
)
const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "vessel"
)
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	CreatVessel(vessel *pb.Vessel)error
}

type VesselRepository struct {
	mgoSession *mgo.Session
}

// 接口实现
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一条容量、载重都符合的货轮
	var vessels[]*pb.Vessel
	repo.Collection().Find(nil).All(vessels)
	for _, v := range vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("No vessel can't be use")
}

func (repo *VesselRepository) CreatVessel(vessel *pb.Vessel)error  {
	return repo.Collection().Insert(vessel)
}

// 关闭连接
func (repo *VesselRepository) Close() {
	// Close() 会在每次查询结束的时候关闭会话
	// Mgo 会在启动的时候生成一个 "主" 会话
	// 你可以使用 Copy() 直接从主会话复制出新会话来执行，即每个查询都会有自己的数据库会话
	// 同时每个会话都有自己连接到数据库的 socket 及错误处理，这么做既安全又高效
	// 如果只使用一个连接到数据库的主 socket 来执行查询，那很多请求处理都会阻塞
	// Mgo 因此能在不使用锁的情况下完美处理并发请求
	// 不过弊端就是，每次查询结束之后，必须确保数据库会话要手动 Close
	// 否则将建立过多无用的连接，白白浪费数据库资源
	repo.mgoSession.Close()
}
func (repo *VesselRepository) Collection()*mgo.Collection  {
	return repo.mgoSession.DB(DB_NAME).C(CON_COLLECTION)
}