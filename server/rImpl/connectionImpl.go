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

	Router rIterface.IRouter
}

//初始化链接模块
func NewConnection(Conn *net.TCPConn, ConnID uint64, Router rIterface.IRouter) *Connection {
	c := &Connection{
		Conn:     Conn,
		ConnID:   ConnID,
		Closed:   true,
		ExitChan: make(chan bool, 1),
		Router:   Router,
	}
	return c
}

func (c *Connection) Start() {
	fmt.Println("Connection  start ConnID =  :", c.ConnID)
	go c.StartReader()
	//todo 读写业务分离
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
func (c *Connection) StartReader() {
	fmt.Println("Connection StartReader ConnID = ", c.ConnID)
	defer fmt.Println("CoonID = ", c.ConnID, "RemoteAddress is exit", c.RemoteAddress().String())
	defer c.Stop()
	for {
		bytes := make([]byte, 512)
		_, err := c.Conn.Read(bytes)
		if err != nil {
			fmt.Println("conn read err", err)
			continue
		}
		req := Request{
			conn: c,
			data: bytes,
		}
		go func(req *Request) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

//发送数据的方法
func (c *Connection) Send(data []byte) error {
	return nil
}

//获取链接当前链接对象的的套接字
func (c *Connection) GetTcpNetConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint64 {
	return c.ConnID
}

func (c *Connection) RemoteAddress() net.Addr {
	return c.Conn.RemoteAddr()
}
