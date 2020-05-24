# 服务与注册发现篇
## 在启动项目前，需要有一个docker环境用来启动一个单机的consul
### 假设你已经有了一个docker环境，启动consul的命令是：docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
### 使用本机ip地址全局替换 192.168.1.5
### 在命令窗口下执行go run server.go -name userservice -p 8081
### 127.0.0.1:8500/ui 就可以看到consul的后台管理了，并且能看到上一步我们注册的userservice服务
