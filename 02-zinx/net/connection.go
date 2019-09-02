package net

import (
	"fmt"
	"net"
	"strings"
	"../iface"
)

//定义connection结构体
type Connection struct {
	conn *net.TCPConn
	connID uint32
	isClosed bool

	//用户注册的处理函数
	callback iface.CallBackFunc
}

func NewConnection(conn *net.TCPConn,cid uint32,callback iface.CallBackFunc) iface.IConnection{
	return &Connection{
		conn:  conn,
		connID:  cid,
		isClosed: false,
		callback: callback,
	}
}

func (c *Connection)Start(){
	go func() {
		for {
			buf := make([]byte, 4096)

			n, err := tcpConn.Read(buf)
			if err != nil {
				fmt.Println("tcpConn.Read err:", err)
				return
			}

			fmt.Println("Server <==== Client, len:", n, ", buf :", string(buf[:n]))

			c.callback(c,buf)

		}

	}()
}
func (c *Connection)Send(data []byte) (int,error){

	n, err := c.conn.Write(data)
	if err != nil {
		fmt.Println("tcpConn.Write err:", err)
		return -1,nil
	}
	return n,nil
}
func (c *Connection)Stop(){
	if c.isClosed {
		return
	}

	_:=c.conn.Close()
}
func (c *Connection)GetConnId() uint32{
	return c.connID
}
//原生的链接
func (c *Connection)GetTcpConn() *net.TCPConn {
	return c.conn
}