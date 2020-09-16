package rImpl

import (
	"errors"
	"fmt"
	"io"
	"net"
	"runx/server/rIterface"
	"runx/util"
)

type Connection struct {
	TcpServer rIterface.ServerIF
	//当前的tcp 链接
	Conn *net.TCPConn

	//tcp Id
	ConnID uint64

	//该链接 当前状态 true 链接中 ，false :关闭
	Closed bool

	// 通知 链接退出 的通道
	ExitChan chan bool
	// 读，写 协程传递数据的通道
	MesCh chan []byte
	// 多路由方法
	Ms rIterface.IMsgHandler
}

//初始化链接模块
func NewConnection(server rIterface.ServerIF, Conn *net.TCPConn, ConnID uint64, handler rIterface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer: server,
		Conn:      Conn,
		ConnID:    ConnID,
		Closed:    true,
		MesCh:     make(chan []byte),
		ExitChan:  make(chan bool, 1),
		Ms:        handler,
	}
	c.TcpServer.GetConnMan().AddConn(c)
	return c
}

func (c *Connection) Start() {
	fmt.Println("Connection  start ConnID =  :", c.ConnID)
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallStartFunc(c)
}

func (c *Connection) Stop() {
	fmt.Println("Connection stop ConnID :", c.ConnID)
	if c.Closed == false {
		return
	}
	c.TcpServer.CallStopFunc(c)
	c.Conn.Close()
	c.Closed = false
	c.TcpServer.GetConnMan().RemoveConn(c)
	close(c.ExitChan)
	close(c.MesCh)
	return
}
func (c *Connection) StartWriter() {
	fmt.Println("[writer]")
	for {
		select {
		case data := <-c.MesCh:
			if _, err := c.GetTcpNetConnection().Write(data); err != nil {
				fmt.Println("writer err,connID: ", c.ConnID, err)
			}
		case <-c.ExitChan:
			fmt.Println("writer exit")
			return
		}
	}
}
func (c *Connection) StartReader() {
	fmt.Println("Connection StartReader ConnID = ", c.ConnID)
	defer fmt.Println("Reader exit, CoonID = ", c.ConnID, "RemoteAddress is exit", c.RemoteAddress().String())
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
		if util.Config.GetInt("server.work-pool-size") > 0 {
			c.Ms.SendMsgToTask(&req)
		} else {
			go c.Ms.DoMsgHandler(&req)

		}
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
	c.MesCh <- bytes
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
