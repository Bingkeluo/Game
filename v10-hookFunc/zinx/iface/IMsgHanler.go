package iface

type IMsgHanler interface {
	AddRouter(uint32,IRouter)
	DoMsgRouter(IRequest)
	StartWorkPool()
	SendReqToQueue(IRequest)
}
