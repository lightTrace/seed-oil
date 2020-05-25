### 使用go自带的rate包，利用go kit的middleware中间件进行简单的限流，ps：这一层限流本身应该是网关做的事，这里为了学习，所以写了出来
### 启动8081的端口服务后，连续快速点击访问刷新http://localhost:8081/user/1 该接口，可以发现页面出现too many request
