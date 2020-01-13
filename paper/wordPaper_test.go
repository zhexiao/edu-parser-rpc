package paper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb1 "github.com/zhexiao/edu-parser-proto/common"
	pb "github.com/zhexiao/edu-parser-proto/paper"
	"testing"
)

func TestWordPaper(t *testing.T) {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://0.0.0.0:2379"}
	})

	service := micro.NewService(
		micro.Name("paper.wordPaper.srv"),
		micro.Registry(reg),
	)

	client := pb.NewWordPaperService("paper.wordPaper.srv", service.Client())
	resp, err := client.Parser(context.TODO(), &pb1.Request{
		Body: []byte("肖哲"),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Data)

}
