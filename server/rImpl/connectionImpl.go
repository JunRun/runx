package rImpl

import (
	"fmt"
	"github.com/runx/server/rIterface"
	"net"
)

type Connection struct {
	//当前的tcp 链接
	Conn *net.TCPConn

	//tcp Id
	ConnID uint64

	//该链接 当前状态 true 链接中 ，false :关闭
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

func (c *Connection) Start() {

}

func (c *Connection) Stop() {
	fmt.Println("Connection stop ConnID :", c.ConnID)
	if c.Closed == false {
		return
	}
	c.Closed = false
	close(c.ExitChan)
	return
}

//获取链接当前链接对象的的套接字
func (c *Connection) GetTcpNetConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

//发送数据的方法
func (c *Connection) Send(data []byte) error {
	return nil
}

func (c *Connection) RemoteAddress() net.Addr {
	return c.Conn.RemoteAddr()
}
