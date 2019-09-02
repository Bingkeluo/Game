package net

import (
	"fmt"
	"net"

	"zinx/04-config/zinx/iface"
)
//connection方法是实现通道的作用：将从客户端获取的请求放在这里进行传递给路由router，再由路由进行具体的分派实现，
//我们获取来自server包，将原本的通道传递数据放在了现在的connection中
//定义connection结构体
type Connection struct {
	conn *net.TCPConn
	connID uint32
	isClosed bool

	//用户注册的处理函数
	//callback iface.CallBackFunc
	//将原本的回调函数改写成了router
	router iface.IRouter//这是一个接口，可以通过接口进行函数的具体实现
}
//新建一个通道，将从server中客户端获取的信息，在这里通过路由进行赋值，返回一个结构体，通过结构体的绑定具体的方法实现，
func NewConnection(conn *net.TCPConn,cid uint32,router iface.IRouter) iface.IConnection{
	return &Connection{
		conn:  conn,
		connID:  cid,
		isClosed: false,
		//callback: callback,
		router: router,
	}
}
//绑定Start方法实现
//start方法主要实现，读取客户端传来的数据，
func (c *Connection)Start(){
	go func() {
		for {
			buf := make([]byte, 4096)

			n, err := c.conn.Read(buf)
			if err != nil {
				fmt.Println("tcpConn.Read err:", err)
				return
			}

			fmt.Println("Server <==== Client, len:", n, ", buf :", string(buf[:n]))
			//将客户端传来的数据，传递给封装的Request方法，这样从Request方法中可以找到所需求不同的数据，data，len，conn
			req:= NewRequest(c,buf,uint32(n))

			//刚刚将传递给请求Request中封装的数据，通过req直接传递给router，router中封装了三个方法PreHandle，Handle，以及PostHandle方法，
			// 这样就能将用户实现的具体业务从框架中剥离出来，实现框架和业务具体实现的分离
			//这些框架中封装的方法又被server_main.go所继承，进行重写，由于就近原则，就可以在server_main.go中实现具体的方法
			//c.callback(req)
			c.router.PreHandle(req)
			c.router.Handle(req)
			c.router.PostHandle(req)
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
	_=c.conn.Close()
}
func (c *Connection)GetConnId() uint32{
	return c.connID
}
//原生的链接
func (c *Connection)GetTcpConn() *net.TCPConn {
	return c.conn
}