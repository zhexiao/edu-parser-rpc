package main

import (
	"github.com/micro/go-micro"
	"github.com/zhexiao/edu-parser-rpc/excel/srv"
	"github.com/zhexiao/edu-parser-rpc/proto/excelpb"
	"github.com/zhexiao/office-parser/utils"
	"log"
	"time"
)

func init() {
	//初始化七牛的配置
	utils.OfficeParserQiniuCfg = &utils.Qiniu{}

	//初始化WMF的配置
	utils.WmfConfiguration = &utils.CT_WmfCfg{}
}

func main() {
	//注册etcd
	//reg := etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{"http://0.0.0.0:2379"}
	//	options.Secure = false
	//})

	service := micro.NewService(
		micro.Name("excel.srv"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		//micro.Registry(reg),
	)

	_ = excelpb.RegisterExcelSrvHandler(service.Server(), new(srv.CT_Excel))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
