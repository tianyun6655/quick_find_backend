package http

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"immouse_backend/immouse_utils/http_client"
	"testing"
)

func TestMain(m *testing.M) {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("greeter.client"))

	service.Init()

	//初始化http
	Init(service, ":8085")

}
func TestRegister(t *testing.T) {
	type RegisterRequest struct {
		Phone string `json:"phone"`
		Code  int    `json:"code"`
	}
	para := &RegisterRequest{
		Phone: "15701381436",
		Code:  1234,
	}
	paraRequest, _ := json.Marshal(para)
	responst, _, err := http_client.SendHTTPReuqest(
		http_client.POST,
		"http://127.0.0.1:8081/quick_find/user/register",
		paraRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response: ", string(responst))
}
