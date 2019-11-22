package znet

import (
	"net"
	"zinxText/ziface"
)

/**
链接模块
 */
type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的处理业务的方法API
	handleAPI ziface.HandleFunc

	//告知当前链接已经退出.停止 channel
	ExitChan chan bool
}

//初始化链接模块的方法
func NewConnection()  {
	
}
