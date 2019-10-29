package rImpl

import "github.com/runx/server/rIterface"

type Request struct {
	conn rIterface.IConnection

	data []byte
}

func (r *Request) GetConnection() rIterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
