/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-30 10:21
 */
package rIterface

//采用TVL  格式 存储消息。解决Tcp 粘包问题
type IDataPack interface {
	//获取包的长度
	GetHeadLen() uint64

	//封包方法
	PackData(m IMessage) ([]byte, error)

	//拆包方法
	UnPackData([]byte) (IMessage, error)
}
