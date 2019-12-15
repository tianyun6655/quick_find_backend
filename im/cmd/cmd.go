package main

import (
	"github.com/micro/go-micro"
	"quick_find_backend/common"
	"quick_find_backend/im/conf"
	dao2 "quick_find_backend/im/dao"
	services "quick_find_backend/im/grpc"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"

)

func main(){
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := micro.NewService(
		micro.Registry(reg),
		micro.Name("quick.find.services.user"),
		micro.Version("latest"),
	)

	common.InitLogger()

	//init config file
	if err:=conf.Init();err!=nil{
		panic(err)
	}
	srv.Init()
	dao,_:=dao2.New(conf.Conf)

	services.RegisterIMServicesHandler(srv.Server(),services.NewImServices(dao))

	srv.Run()


}
