package net

import (
	"fmt"
	"zinx/v8-workpool/zinx/config"
	"zinx/v8-workpool/zinx/iface"
)

type MsgHandler struct {
	msghandle map[uint32]iface.IRouter
	workpool int
	taskQueue []chan iface.IRequest
}

func NewMsgHandler()* MsgHandler{

	wookpoolsize:=config.GlobalConfig.WorkerSize
	return &MsgHandler{

		msghandle:make(map[uint32]iface.IRouter),
		workpool: wookpoolsize,
		taskQueue: make([]chan iface.IRequest,wookpoolsize),

	}
}



//建立 连接池 消息队列，给每一个消息队列分配空间，并监听任务。
func (mh *MsgHandler)StartWorkPool(){
		fmt.Println("StartWorkPool start...")
	for i := 0; i <mh.workpool; i++ {
fmt.Println("第",i,"个工作池启动")
		//给每一个工作池，创建消息队列的容量
		mh.taskQueue[i]=make(chan iface.IRequest,config.GlobalConfig.TaskQueSize)

		//并监听任务
		go func( i int ) {
			for{
				req:=<-mh.taskQueue[i]
				mh.DoMsgRouter(req)
			}

		}(i)

	}
}

//提供一个方法，向任务队列发送请求
func (mh *MsgHandler)SendReqToQueue(req iface.IRequest){
	//每一个链接分配一个worker，
	//同一个worker可以服务多个链接
	//1. 先获取连接cid
	cid:=req.GetConnection().GetConnId()
	fmt.Println("创建链接+++++++++",cid)
	//获取当前连接所分配的workerid
		workid:=int(cid) % mh.workpool
	fmt.Println("添加cid:", cid, " 的请求到workerid:", workid)
	//将请求放入对应的worker的消息队列中
	mh.taskQueue[workid]<-req
}


//添加新的路由
func (mh *MsgHandler)AddRouter(msgid uint32,router iface.IRouter){

	//根据参数msgid确定已存id是否含有需要添加的id，如果有直接返回，没有添加
		_,ok:=mh.msghandle[msgid]

	if ok {
		fmt.Println("要输入的id已经存在：",msgid)
		return
	}

		mh.msghandle[msgid]=router
}


//执行路由操作

func (mh *MsgHandler)DoMsgRouter(req iface.IRequest){
	//根据已知的id判断是否存在对应的router，存在执行，不存在直接返回

	msgid:=req.GetMessage().GetMsgId()
	router,ok:=mh.msghandle[msgid]
	if !ok {
		fmt.Println("要输入的id不存在：",msgid)
		return
	}
	router.PreHandle(req)
	router.Handle(req)
	router.PostHandle(req)


}

