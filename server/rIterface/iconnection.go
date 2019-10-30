package rIterface

import "net"

type IConnection interface {
	Start()

	Stop()
	//获取链接当前链接对象的的套接字
	GetTcpNetConnection() *net.TCPConn

	GetConnID() uint64
	//发送数据的方法
	SendMsg(Id uint64, data []byte) error

	RemoteAddress() net.Addr
}
