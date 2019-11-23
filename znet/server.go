package znet

import (
	"fmt"
	"net"
	"zinxText/ziface"
)

//IServer的接口实现，定义一个个Server的服务器模块

type Server struct {
	//服务器名称
	Name string
	//服务器绑定的IP版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port string
	//当前的server添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter
}

/*
//定义当前客户端连接的所绑定的handle api（目前这个handle是写死的，)
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error  {
	//回显的业务
	fmt.Println("[Conn Handle CallBackToClient...]")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallbackToClient error")
	}
	return nil
}
 */


//启动服务器
func (s *Server) Start()  {
	fmt.Println("[Start] Server Listenner at IP:%s, Port %d, is starting\n ", s.IP, s.Port)
	go func() {
		// 1.获取一个TCP的Address
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP,s.Port))
		if err != nil {
			fmt.Println("resolve tcp aaddt err:", err)
			return
		}
		//2.监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listent", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succ", s.Name, "succ, Listening...")
		var cid uint32
		cid = 0
		//3.阻塞的等待客户端连接，处理客户端连接业务
		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//将处理该连接的业务方法和conn进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid ++

			//启动当前的连接业务处理
			go dealConn.Start()
			/*
			//已经与客户端建立连接，做一些业务，做一个最基本的最大512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
			 */
		}
	}()
}
//停止服务器
func (s *Server) Stop()  {
	//TODO 将一些服务器的资源，状态或者一些已经开辟的链接信息，进行停止或修改
}
//运行服务器
func (s *Server) Serve()  {
	//启动server的服务功能
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态
	select {

	}
}
//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")
}

/**
初始化Server模块的方法
 */
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "8999",
		Router: nil,
	}
	return s
}