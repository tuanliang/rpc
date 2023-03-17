package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立socket连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 执行rpc方法
	var resp string
	err = client.Call("HelloService.Hello", "shiyi", &resp)
	if err != nil {
		panic(err)
	}

	// 打印调用相应
	fmt.Println(resp)
}
