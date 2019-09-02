package iface

type IServer interface {
	//方法
	//1.启动
	Start()
	//2.停止
	Stop()
	//3.服务
	Server()

	AddRouter(IRouter)
}
