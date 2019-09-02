package net

import "zinx/03-Request/iface"

type Request struct {
	conn iface.IConnection
	data []byte
	len  uint32
}

func NewRequest(conn iface.IConnection,data []byte,len uint32) iface.IRequest{
	return &Request{
		conn: conn,
		data: data,
		len: nil,
	}
}

func (req *Request)GetConnection()iface.IConnection{
	return req.conn
}

func (req Request)GetData()[]byte{
	return req.data
}

func (req Request)GetLen() uint32{
	return req.len
}