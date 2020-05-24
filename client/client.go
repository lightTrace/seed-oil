package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"gokit/client/service"
	"io"
	"net/url"
	"os"
)

func main() {

	config :=  api.DefaultConfig()
	config.Address = "192.168.1.5:8500"
	apiClient,err := api.NewClient(config)
	if err != nil{

	}
	client := consul.NewClient(apiClient)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	tags := []string{"primary"}
	//可实时查询服务实例的状态信息
	instance := consul.NewInstancer(client,logger,"userservice",tags,true)

	f := func(serviceUrl string)(endpoint.Endpoint,io.Closer,error) {
		tart,_ := url.Parse("http://" + serviceUrl)
		return httptransport.NewClient("GET",tart,service.GetUserInfoRequest,service.GetUserInfoResponse).Endpoint(),nil,nil
	}
	endPointer := sd.NewEndpointer(instance,f,logger)
	endpoints,err := endPointer.Endpoints()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("服务有",len(endpoints),"条")
	getUserInfo := endpoints[0]

	ctx := context.Background()

	res,err := getUserInfo(ctx,service.UserRequest{Uid:101})
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	userInfo := res.(service.UserResponse)
	fmt.Println(userInfo.Result)
}