package net

import "zinx/05-message/zinx/iface"
//将客户端的请求进行封装，这样职责更加分明，
//将请求分为
type Request struct {
	conn iface.IConnection
	data []byte
	len  uint32
}
//将从connection中获取的数据赋值给结构体，再将结构体作为接收者绑定给具体的实现方法，但此时的方法接收者是有值的了
func NewRequest(conn iface.IConnection,data []byte,len uint32) iface.IRequest{
	return &Request{
		conn: conn,
		data: data,
		len: len,
	}
}

//写这些方法的好处是，通过调用这些方法就能直接实现获取相应的数据
func (req *Request)GetConnection()iface.IConnection{
	return req.conn
}

func (req Request)GetData()[]byte{
	return req.data
}

func (req Request)GetLen() uint32{
	return req.len
}