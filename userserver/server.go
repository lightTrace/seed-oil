package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gokit/userserver/service"
	"gokit/userserver/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	name := flag.String("name","","服务名")
	port := flag.Int("p",0,"服务端口")
	flag.Parse()
	if *name == ""{
		log.Fatal("请指定服务名")
	}
	if *port == 0{
		log.Fatal("请指定服务端口")
	}
	util.SetServiceInfo(*name,*port)

	user := service.UserService{}
	endpoint := service.GenUserEndpoint(&user)
	userHandler := httptransport.NewServer(endpoint,service.DecodeUserRequest,service.EncodeResponse)

	http.Handle("/user",userHandler)

	r := mux.NewRouter()
	{
		r.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(userHandler)
		r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type","application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}
	errChan := make(chan error)
	go(func() {
		//注册服务
		util.RegisterService()
		if err := http.ListenAndServe(":" + strconv.Itoa(*port), r);err != nil{
			log.Println(err)
			errChan <- err
		}
	})()

	go(func() {
		signalC := make(chan os.Signal)
		signal.Notify(signalC,syscall.SIGINT,syscall.SIGTERM)
		errChan <-fmt.Errorf("%s",<-signalC)
	})()
	getErr := <- errChan
	util.UnRegisterService()
	log.Println(getErr)
}
