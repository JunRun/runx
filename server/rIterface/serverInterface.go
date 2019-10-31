package rIterface

type ServerIF interface {
	//start server
	Start()
	//stop server
	Stop()
	//添加路由
	AddRouter(id uint64, router IRouter)
	//获取 链接管理模块
	GetConnMan() IConnManger

	//设置开始hook 方法
	SetStartFunc(func(connection IConnection))

	//设置接收的hook方法
	SetStopFunc(func(connection IConnection))

	//调用开始的Hook方法
	CallStartFunc(connection IConnection)
	//调用结束的hook方法
	CallStopFunc(connection IConnection)
}
