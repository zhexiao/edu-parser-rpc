package main

import (
	"github.com/micro/go-micro"
	"github.com/zhexiao/edu-parser-rpc/base"
	"github.com/zhexiao/edu-parser-rpc/proto/wordpb"
	"github.com/zhexiao/edu-parser-rpc/word/srv"
	"github.com/zhexiao/office-parser/utils"
	"log"
	"time"
)

func init() {
	//读取配置文件
	yamlSettings := base.ReadYaml("./settings.yaml")

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
	//注册etcd
	//reg := etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{"http://0.0.0.0:2379"}
	//	options.Secure = false
	//})

	service := micro.NewService(
		micro.Name("word.srv"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		//micro.Registry(reg),
	)

	_ = wordpb.RegisterWordSrvHandler(service.Server(), new(srv.CT_Word))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
