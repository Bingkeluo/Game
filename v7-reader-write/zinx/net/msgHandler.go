package net

import (
	"fmt"
	"zinx/v6-router/zinx/iface"
)

type MsgHandler struct {
	msghandle map[uint32]iface.IRouter
}

func NewMsgHandler()* MsgHandler{
	return &MsgHandler{
		msghandle:make(map[uint32]iface.IRouter),
	}
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

