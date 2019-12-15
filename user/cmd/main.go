package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"quick_find_backend/common"
	"quick_find_backend/user/conf"
	"quick_find_backend/user/dao"
	services "quick_find_backend/user/grpc"
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

	//init dao
	dao,err:=dao.New(conf.Conf)
    if err!=nil{
    	panic(err)
	}
	srv.Init()

	services.RegisterUserServicesHandler(srv.Server(),services.NewUserService(dao))

	if err := srv.Run(); err != nil {
		panic(err)
	}
}