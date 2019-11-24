package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	fmt.Println("Client start ....")

	time.Sleep(1 * time.Second)

	//1 直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp","127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}

	for {
		//2 连接调用Write 写数据
		_, err := conn.Write([]byte("Hello Zinx V0.2"))
		if err != nil {
			fmt.Println("write conn err",err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		//CPU阻塞
		time.Sleep(1 * time.Second)
	}
}
