## 前言
* 会尽量写好注释，run起代码的细节以及各个组件使用的文档，有不清楚的地方大家请issue，希望对大家的学习有一点点帮助就好
* 学习这条路总是漫长且孤独的，聪明人也要下笨功夫，大家一起加油！莫愁前路无知己，天涯谁人不识君

## 服务与注册发现篇
* 下载项目到本地：git clone git@github.com:lightTrace/seed-oil.git
* 在启动项目前，需要有一个docker环境用来启动一个单机的consul
* 假设你已经有了一个docker环境，启动consul的命令是：docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
* 使用本机ip地址全局替换 /userserver/util/consul.go文件下的 reg.Address 变量
* 在/userserver目录下开启一个命令窗口，执行go run server.go -name userservice -p 8081
* 127.0.0.1:8500/ui 就可以看到consul的后台管理了，并且能看到上一步我们注册的userservice服务
* 在/userserver目录下再开启一个命令窗口，执行go run server.go -name userservice -p 8082，这里我们开启两个服务，client客户端准备使用轮询的方式访问这两个节点提供的相同的服务，服务名都是userservice
* 在/client目录下直接运行go run client.go，发现在调用注册好的userservice服务，8081，8082两个端口依次提供服务

## 服务限流篇
* [服务限流文档地址](https://github.com/lightTrace/seed-oil/tree/master/tool/rate)

## 服务熔断篇
* 取消 /client/logic/getuserinfo.go的time.Sleep(time.Second * 3),可以模拟熔断降级
