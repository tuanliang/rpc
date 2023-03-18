package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/tuanliang/rpc/pbrpc/codec/client"
	"github.com/tuanliang/rpc/pbrpc/service"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 建立socket连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端采用json格式来编码
	client := rpc.NewClientWithCodec(client.NewClientCodec(conn))

	return &HelloServiceClient{
		client: client,
	}, nil

}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request *service.Request, response *service.Response) error {
	// 执行rpc方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var resp service.Response
	if err := c.Hello(&service.Request{Value: "shiyi"}, &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp.Value)
}
