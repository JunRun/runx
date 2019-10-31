/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-31 14:23
 */
package rImpl

import (
	"errors"
	"fmt"
	"github.com/runx/server/rIterface"
	"strconv"
	"sync"
)

type ConnManger struct {
	connMap map[uint64]rIterface.IConnection
	lock    sync.RWMutex
}

func NewConnManger() *ConnManger {
	c := &ConnManger{
		connMap: make(map[uint64]rIterface.IConnection),
	}
	return c
}

//添加链接
func (c *ConnManger) AddConn(conn rIterface.IConnection) {
	c.lock.Lock()
	c.connMap[conn.GetConnID()] = conn
	c.lock.Unlock()
	fmt.Println("Conn add ID:", conn.GetConnID())
}

//删除连接
func (c *ConnManger) RemoveConn(conn rIterface.IConnection) {
	c.lock.Lock()
	delete(c.connMap, conn.GetConnID())
	c.lock.Unlock()
	fmt.Println("Conn quit ID:", conn.GetConnID())
}
func (c *ConnManger) Len() int {
	return len(c.connMap)
}

func (c *ConnManger) Get(ConnId uint64) (rIterface.IConnection, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if v, ok := c.connMap[ConnId]; ok {
		return v, nil
	} else {
		return nil, errors.New("the ConnID is not exist :ID ==" + strconv.Itoa(int(ConnId)))
	}
}

//清空连接
func (c *ConnManger) Clear() {
	c.lock.Lock()
	for connId, conn := range c.connMap {
		conn.Stop()
		delete(c.connMap, connId)
	}
	c.lock.Unlock()
	fmt.Println("clear all connection")
}
