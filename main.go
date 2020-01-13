package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "github.com/zhexiao/edu-parser-proto/paper"
	"github.com/zhexiao/edu-parser-rpc/paper"
	"time"
)

func main() {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://0.0.0.0:2379"}
		options.Secure = false
	})

	service := micro.NewService(
		micro.Name("paper.wordPaper.srv"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
	)

	_ = pb.RegisterWordPaperHandler(service.Server(), new(paper.CT_WordPaper))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
