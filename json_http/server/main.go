package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/tuanliang/rpc/json_http/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

// service handler
type HelloService struct {
}

func NewRPCReadWriteCloser(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

// request收到name，response相应hello，name
func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello,%s", request)
	return nil
}

func (s *HelloService) Calc(req *service.CalcRequest, response *int) error {
	*response = req.A + req.B
	return nil
}

func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SERVICE_NAME, new(HelloService))

	// 通过jsonrpc这个path，来处理所有的请求
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriteCloser(w, r)))
	})

	http.ListenAndServe(":1234", nil)
}
