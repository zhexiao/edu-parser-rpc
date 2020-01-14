# edu-parser-rpc
基于GRPC与go-micro的教育资源解析RPC服务。

[![Build Status](https://travis-ci.org/zhexiao/edu-parser-rpc.svg?branch=master)](https://travis-ci.org/zhexiao/edu-parser-rpc)
[![codecov](https://codecov.io/gh/zhexiao/edu-parser-rpc/branch/master/graph/badge.svg)](https://codecov.io/gh/zhexiao/edu-parser-rpc)
![go](https://img.shields.io/badge/go->%3D1.13-blue)


## 基本需求
1. golang >= 1.13
2. protoc >= 3.11.2
3. go-micro >= 1.18

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