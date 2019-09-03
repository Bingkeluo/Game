package net

import "zinx/v6-router/zinx/iface"
//将客户端的请求进行封装，这样职责更加分明，
//将请求分为
type Request struct {
	conn iface.IConnection
	message iface.IMessage//定义成接口类型可以通过接口调用Get方法找到对应的数据
}
//将从connection中获取的数据赋值给结构体，再将结构体作为接收者绑定给具体的实现方法，但此时的方法接收者是有值的了
func NewRequest(conn iface.IConnection,message iface.IMessage) iface.IRequest{
	return &Request{
		conn: conn,
		message:message,
	}
}

//写这些方法的好处是，通过调用这些方法就能直接实现获取相应的数据
func (req *Request)GetConnection()iface.IConnection{
	return req.conn
}

func (req Request)GetMessage()iface.IMessage{
	return req.message
}
