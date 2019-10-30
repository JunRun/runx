/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-29 16:58
 */
package rImpl

type Message struct {
	MessageId     uint64
	MessageLength uint64
	Data          []byte
}

func (m *Message) GetMessageId() uint64 {
	return m.MessageId
}
func (m *Message) GetMessageLen() uint64 {
	return m.MessageLength
}
func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMessageId(id uint64) {
	m.MessageId = id
}
func (m *Message) SetMessageLen(len uint64) {
	m.MessageLength = len
}
func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}
