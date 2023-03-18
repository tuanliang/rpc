# Hello World Grpc

```shell
cd /e/go_project/rpc/grpc/simple/server
# 生成service pb编译文件
protoc -I=. --go_out=. --go_opt=module="github.com/tuanliang/rpc/grpc/simple/server" pb/hello.proto


# 补充rpc 接口定义protobuf的代码生成
protoc -I=. --go_out=. --go_opt=module="github.com/tuanliang/rpc/grpc/simple/server" pb/hello.proto --go-grpc_out=. --go-grpc_opt=module="github.com/tuanliang/rpc/grpc/simple/server" pb/hello.proto
```