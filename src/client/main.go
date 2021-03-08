package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asim/go-micro/v3/cmd"
	"github.com/asim/go-micro/v3/client"
	pb "github.com/consingment-service/proto/consignment"
	"io/ioutil"
	"log"
	"os"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	cmd.Init()
	// 创建微服务的客户端，简化了手动 Dial 连接服务端的步骤
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)
	log.Printf("client create.....")
	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	log.Printf("client ready to pare file.....")
	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 调用 RPC
	// 将货物存储到我们自己的仓库里

	log.Print(consignment)
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	log.Print(resp)
	log.Printf("created: %t", resp.Created)

    resp1,err:=client.GetAllConsignment(context.Background(),&pb.Request{})
	log.Print(resp1)
	if err != nil {
		log.Fatalf("get All consignment error: %v", err)
	}
    log.Printf("All consignment count is %d",len(resp1.Consignments))
	// 新货物是否托运成功

}