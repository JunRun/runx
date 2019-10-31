/**
 *
 * @Description: 连接管理模块
 * @Version: 1.0.0
 * @Date: 2019-10-31 14:20
 */
package rIterface

type IConnManger interface {
	AddConn(conn IConnection)
	RemoveConn(conn IConnection)
	//根据连接ID获取 获取连接
	Get(id uint64) (IConnection, error)
	//获取总连接的个数
	Len() int
	//清空链接
	Clear()
}
