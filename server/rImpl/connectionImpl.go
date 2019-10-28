package rImpl

import (
	"github.com/runx/server/rIterface"
	"net"
)

type Connection struct {
	//当前的tcp 链接
	Conn *net.TCPConn

	//tcp Id
	ConnID uint64

	//该链接 当前状态
	Closed bool

	// 通知 链接退出 的通道
	ExitChan chan bool

	//当前链接所绑定的 业务处理方法
	HandelApi rIterface.HandleFunc
}

//初始化链接模块
func NewConnection(Conn *net.TCPConn, ConnID uint64, HandelApi rIterface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      Conn,
		ConnID:    ConnID,
		Closed:    true,
		ExitChan:  make(chan bool, 1),
		HandelApi: HandelApi,
	}
	return c
}
