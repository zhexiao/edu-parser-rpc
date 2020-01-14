package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/zhexiao/edu-parser-proto/wordpb"
	"github.com/zhexiao/edu-parser-rpc/word"
	"github.com/zhexiao/office-parser/utils"
	"time"
)

func init() {
	//读取配置文件
	yamlSettings := ReadYaml()

	//初始化七牛的配置
	utils.OfficeParserQiniuCfg = &utils.Qiniu{
		//七牛key
		AccessKey: yamlSettings.Qiniu.AccessKey,
		//七牛secret
		SecretKey: yamlSettings.Qiniu.SecretKey,
		//七牛存储的bucket
		Bucket: yamlSettings.Qiniu.Bucket,
		//所属区域
		Zone: yamlSettings.Qiniu.Zone,
		//访问的域名地址
		Domain: yamlSettings.Qiniu.Domain,
		//路径
		UploadPrefix: yamlSettings.Qiniu.UploadPrefix,
	}

	//初始化WMF的配置
	utils.WmfConfiguration = &utils.CT_WmfCfg{
		Uri: yamlSettings.Wmf.Uri,
	}
}

func main() {
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://0.0.0.0:2379"}
		options.Secure = false
	})

	service := micro.NewService(
		micro.Name("word.srv"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
	)

	_ = wordpb.RegisterWordSrvHandler(service.Server(), new(word.CT_Word))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
