package iface

type IServer interface {
	//方法
	//1.启动
	Start()
	//2.停止
	Stop()
	//3.服务
	Server()

	AddRouter(uint32,IRouter)//这里添加的是msgid，和用户要加的路由

	GetConnMAgr() IConnManager

	RegisterStartHookFunc(func(connection IConnection))

	RegisterStopHookFunc(func(connection IConnection))

	//7. 提供调用钩子函数的方法
	CallStartHookFunc(IConnection)
	CallStopHookFunc(IConnection)
}
