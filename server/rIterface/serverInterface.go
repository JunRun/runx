package rIterface

type ServerIF interface {
	//start server
	Start()
	//stop server
	Stop()

	AddRouter(id uint64, router IRouter)
}
