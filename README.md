# edu-parser-rpc
基于GRPC与go-micro的教育资源解析RPC服务。

## protobuf 编译
```shell script
$ cd proto
$ protoc --micro_out=paths=source_relative:. --go_out=paths=source_relative:. ./basepb/*.proto
$ protoc --micro_out=paths=source_relative:. --go_out=paths=source_relative:. ./wordpb/*.proto
$ protoc --micro_out=paths=source_relative:. --go_out=paths=source_relative:. ./excelpb/*.proto
```

## 配置文件
```shell script
$ cp settings.yaml.example settings.yaml
```