package rImpl

import "github.com/runx/server/rIterface"

type Request struct {
	conn rIterface.IConnection

	message rIterface.IMessage
}

func (r *Request) GetConnection() rIterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.message.GetData()
}

func (r *Request) GetMessageID() uint64 {
	return r.message.GetMessageId()
}

func (r *Request) GetHeadLen() uint64 {
	return r.message.GetMessageLen()
}
