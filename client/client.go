package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/prometheus/common/log"
	"gokit/client/logic"
	"time"
)

func main() {
	//熔断配置
	hystrixConfig := hystrix.CommandConfig{
		Timeout:                2000,                 //熔断器的超时时间
		MaxConcurrentRequests:  5,                    //最大并发
		RequestVolumeThreshold: 3,                    //失败几次开始计算err占比
		ErrorPercentThreshold:  20,                   //出现错误的占比超过20%的时候打开熔断器
		SleepWindow:            int(time.Second * 5), //5秒后熔断器尝试开启半开状态，半开状态下熔断器尝试访问正常服务来恢复服务，不能让其一直熔断

	}
	hystrix.ConfigureCommand("getUserInfo", hystrixConfig)
	for i := 0; i < 10; i++ {
		err := hystrix.Do("getUserInfo", func() error {
			result, err := logic.GetUserInfo()
			fmt.Println(result)
			return err
		}, func(e error) error {
			fmt.Println("降级用户")
			return e
		})
		if err != nil {
			log.Error(err)
		}
	}
}
