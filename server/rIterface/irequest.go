package rIterface

type IRequest interface {
	GetConnection() *IConnection

	GetData() []byte
}
