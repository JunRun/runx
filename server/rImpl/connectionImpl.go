package rImpl

import (
	"errors"
	"fmt"
	"github.com/runx/server/rIterface"
	"io"
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
		pack := NewDataPack()
		//读取头部数据
		bytes := make([]byte, pack.GetHeadLen())
		if _, err := io.ReadFull(c.GetTcpNetConnection(), bytes); err != nil {
			fmt.Println("read msg head err")
			break
		}

		headMessage, err := pack.UnPackData(bytes)
		if err != nil {
			fmt.Println(err)
			break
		}

		data := make([]byte, headMessage.GetMessageLen())
		if headMessage.GetMessageLen() > 0 {
			//读取数据流
			if _, err := io.ReadFull(c.GetTcpNetConnection(), data); err != nil {
				fmt.Println("read data err")
				break
			}
		}
		headMessage.SetData(data)
		//创建请求体
		req := Request{
			conn:    c,
			message: headMessage,
		}
		//执行router方法
		go func(req *Request) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

//发送数据的方法
func (c *Connection) SendMsg(Id uint64, data []byte) error {
	if c.Closed == false {
		return errors.New("channel was closed")
	}
	m := &Message{
		MessageId:     Id,
		MessageLength: uint64(len(data)),
		Data:          data,
	}
	pack := DataPack{}
	bytes, err := pack.PackData(m)
	if err != nil {
		return err
	}
	_, er := c.GetTcpNetConnection().Write(bytes)
	if er != nil {
		return er
	}
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
