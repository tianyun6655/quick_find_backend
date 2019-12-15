package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"quick_find_backend/common"
	"quick_find_backend/gateway/http"
)

func main(){

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	common.InitLogger()
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("greeter.client"))

	service.Init()


	//初始化http
    http.Init(service,":8085")

}
