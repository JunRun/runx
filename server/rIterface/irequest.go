package rIterface

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte

	GetHeadLen() uint64

	GetMessageID() uint64
}
