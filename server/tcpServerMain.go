package main

import (
	"net"
	"fmt"
	"strconv"
)

var(
	ip string = ""
	port int = 9901
	j int = 0
)
var tcpListen *net.TCPListener
func main()  {
	createTcpConn()
	server()
}

func createTcpConn()  {
	var err error
	tcpListen,err = net.ListenTCP("tcp",&net.TCPAddr{net.ParseIP(""),port,""})
	if err != nil{
		fmt.Printf("net.ListenTCP error |err=%v\n",err)
		panic(111)
	}
	fmt.Println("初始化连接 success")
}

func server()  {
	for {
		conn,err := tcpListen.AcceptTCP()
		if err != nil{
			fmt.Printf("tcpListen.AcceptTCP err |err=%v\n",err)
			continue
		}
		fmt.Printf("tcpListen.AcceptTCP 获取数据信息: = %+v\n",conn)
		defer conn.Close()
		go func() {
			data := make([]byte,128)
			for {
				i,err := conn.Read(data)
				fmt.Printf("接收客户端发送数据：=%v,\n",string(data[0:i]))
				if err != nil{
					fmt.Printf("接收客户端发送数据 error |err=%v\n",err)
					break
				}
				j++
				conn.Write([]byte("my name is Wang -- " + strconv.Itoa(j)))
			}
		}()

	}
}