package rIterface

type ServerIF interface {
	//start server
	Start()
	//stop server
	Stop()
	//

	AddRouter(router IRouter)
}
