# 代码生成

```shell
# 生成service pb编译文件  
protoc -I=./service/pb/ --go_out=./service/ --go_opt=module="github.com/tuanliang/rpc/pbrpc/service" hello.proto
```

```shell
# 生成codec pb编译文件
protoc -I=./codec/pb/ --go_out=./codec/pb/ --go_opt=module="github.com/tuanliang/rpc/pbrpc/codec/pb" header.proto
```