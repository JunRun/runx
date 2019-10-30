/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-29 16:55
 */
package rIterface

type IMessage interface {
	GetMessageId() uint64
	GetMessageLen() uint64
	GetData() []byte

	SetMessageId(id uint64)
	SetMessageLen(len uint64)
	SetData(bytes []byte)
}
