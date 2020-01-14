package word

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/zhexiao/edu-parser-proto/basepb"
	"github.com/zhexiao/edu-parser-proto/wordpb"
	"io/ioutil"
	"log"
	"testing"
)

func TestWordPaper(t *testing.T) {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://0.0.0.0:2379"}
	})

	service := micro.NewService(
		micro.Name("word.srv"),
		micro.Registry(reg),
	)

	fileBytes, err := ioutil.ReadFile("../_testdoc/paper.docx")
	if err != nil {
		log.Println(err)
	}

	client := wordpb.NewWordSrvService("word.srv", service.Client())
	resp, err := client.ParserPaper(context.TODO(), &basepb.Request{
		Body: fileBytes,
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(resp.Data)
}

func TestWordQuestion(t *testing.T) {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://0.0.0.0:2379"}
	})

	service := micro.NewService(
		micro.Name("word.srv"),
		micro.Registry(reg),
	)

	fileBytes, err := ioutil.ReadFile("../_testdoc/question-fill.docx")
	if err != nil {
		log.Println(err)
	}

	client := wordpb.NewWordSrvService("word.srv", service.Client())
	resp, err := client.ParserQuestion(context.TODO(), &basepb.Request{
		Body: fileBytes,
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(resp.Data)
}
