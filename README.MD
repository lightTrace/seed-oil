## 服务与注册发现篇
* 下载项目到本地：git clone git@github.com:lightTrace/seed-oil.git
* 在启动项目前，需要有一个docker环境用来启动一个单机的consul
* 假设你已经有了一个docker环境，启动consul的命令是：docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
* 使用本机ip地址全局替换 192.168.1.5
* 在gokit/userserver目录下开启一个命令窗口，执行go run server.go -name userservice -p 8081
* 127.0.0.1:8500/ui 就可以看到consul的后台管理了，并且能看到上一步我们注册的userservice服务
* 在gokit/userserver目录下再开启一个命令窗口，执行go run server.go -name userservice -p 8082，这里我们开启两个服务，client客户端准备使用轮询的方式访问这两个节点提供的相同的服务，服务名都是userservice
* 在gokit/client目录下直接运行go run client.go，发现在依次调用两个端口提供的服务

## 服务限流篇
* [服务限流文档地址](https://github.com/lightTrace/seed-oil/tree/master/tool/rate)
