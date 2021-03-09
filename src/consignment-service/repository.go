package main

import (
	pb "github.com/consingment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)
const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "consignments"
)
//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() ([]*pb.Consignment,error)                                  // 获取仓库中所有的货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
	mgoSession  *mgo.Session
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
    err:=repo.Collection().Insert(consignment)
    if err!=nil{
    	return nil,err
	}
	return consignment, nil
}

func (repo *Repository) GetAll() ([]*pb.Consignment,error) {
	var cons []*pb.Consignment
	// Find() 一般用来执行查询，如果想执行 select * 则直接传入 nil 即可
	// 通过 .All() 将查询结果绑定到 cons 变量上
	// 对应的 .One() 则只取第一行记录
	err := repo.Collection().Find(nil).All(&cons)
	return cons, err
}
// 关闭连接
func (repo *Repository) Close() {
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
func (repo *Repository) Collection()*mgo.Collection  {
   return repo.mgoSession.DB(DB_NAME).C(CON_COLLECTION)
}
