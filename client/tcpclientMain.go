package main

import (
	"net"
	"fmt"
)

var (
	ip string = "127.0.0.1"
	port int = 9901
)
var conn net.Conn

func main()  {
	createClient()
	client()
}

func createClient()  {
	var err error
	defer func() {
		if err := recover();err != nil{
			fmt.Printf("createClient.defer error|err=%v\n",err)
		}
	}()

	conn,err = net.Dial("tcp",fmt.Sprintf("%s:%d",ip,port))

	if err != nil {
		fmt.Printf("net.Dial 失败 err=%v\n",err)
		return
	}
	fmt.Printf("创建客户端连接成功 success....\n")

}

func client() {
	sms := make([]byte,128)
	for{
		_,err := fmt.Scan(&sms)
		if err != nil{
			fmt.Printf("数据输入异常:%v\n",err)
		}
		fmt.Printf("数据信息：=%v\n",string(sms))
		clientSendMoreTwice(sms,0)
		fmt.Printf("nextTime\n")
	}
}


func clientSendMoreTwice(sms []byte, num int)  {
	if num > 2 {
		fmt.Printf("最多重试2次，已超过次数,停止重试,num=%v\n",num)
		return
	}

	fmt.Printf("conn=%+v\n",conn)
	if conn == nil{
		retryConnect()
		clientSendMoreTwice(sms,num+1)
		return
	}
	n,err := conn.Write(sms)
	if n == 0 && err != nil {
		//链接报错，从事机制
		retryConnect()
		clientSendMoreTwice(sms,num+1)
		return
	}

	fmt.Printf("conn.write: n=%v,err=%v\n",n,err)
	buf := make([]byte,128)
	c,err := conn.Read(buf)
	if err != nil{
		fmt.Printf("读取服务端返回数据失败:=%v\n",err)
		clientSendMoreTwice(sms,num+1)
		return
	}
	fmt.Printf("服务端返回信息=%v\n",string(buf[0:c]))
}

func retryConnect(){
	fmt.Printf("重新启动客户端\n")
	createClient()
}
