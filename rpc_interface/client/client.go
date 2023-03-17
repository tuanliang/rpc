package main

import (
	"fmt"
	"net/rpc"

	"github.com/tuanliang/rpc/rpc_interface/service"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 建立socket连接
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: client,
	}, nil

}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	// 执行rpc方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var resp string
	if err := c.Hello("shiyi", &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
