# GO TCP长连接业务逻辑
本项目简单实现TCP的客户端和服务端连接逻辑，client端在连接断开后会自动重试，仅供参考学习

# 下载/安装
克隆项目到本地

# server端启动
cd tcpserver && go run server/tcpServerMain.go

# client端启动
cd tcpserver && go run client/tcpclientMain.go
