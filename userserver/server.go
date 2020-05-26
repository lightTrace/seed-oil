package main

import (
	"flag"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	logtool "gokit/tool/log"
	ratetool "gokit/tool/rate"
	"gokit/userserver/service"
	"gokit/userserver/util"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	name := flag.String("name", "", "服务名")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定服务端口")
	}
	util.SetServiceInfo(*name, *port)

	//go-kit的日志模块
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger, "go-kit", "1.0")
		logger = kitlog.With(logger, "time", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	user := service.UserService{}
	//每秒钟只能接受一个请求，但是可以容忍瞬间提高的5个请求，超过的请求会报429
	limit := rate.NewLimiter(1, 5)
	//使用无耦合的限流中间件和日志中间件去包装handler
	endpoint := ratetool.RateLimit(limit)(logtool.UserServiceLogger(logger)(service.GenUserEndpoint(&user)))

	//自定义error的解码
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(util.MyErrorEncoder),
	}
	userHandler := httptransport.NewServer(endpoint, service.DecodeUserRequest, service.EncodeResponse, options...)

	r := mux.NewRouter()
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(userHandler)
		r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}
	errChan := make(chan error)
	go (func() {
		//注册服务
		util.RegisterService()
		if err := http.ListenAndServe(":"+strconv.Itoa(*port), r); err != nil {
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		//监听系统的退出信号
		signalC := make(chan os.Signal)
		signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signalC)
	})()
	//一旦有异常信号进来就不阻塞退出服务
	getErr := <-errChan
	//优雅退出服务
	util.UnRegisterService()
	log.Println(getErr)
}
