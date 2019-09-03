package net

import (
	"fmt"
	"io"
	"net"

	"zinx/v6-router/zinx/iface"
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
	//router iface.IRouter//这是一个接口，可以通过接口进行函数的具体实现
	msghandler iface.IMsgHanler//这里面包含了路由和函数的集合
}
//新建一个通道，将从server中客户端获取的信息，在这里通过路由进行赋值，返回一个结构体，通过结构体的绑定具体的方法实现，
func NewConnection(conn *net.TCPConn,cid uint32,msghandler iface.IMsgHanler) iface.IConnection{
	return &Connection{
		conn:  conn,
		connID:  cid,
		isClosed: false,
		//callback: callback,
		msghandler: msghandler,
	}
}
//绑定Start方法实现
//start方法主要实现，读取客户端传来的数据，
func (c *Connection)Start(){
	go func() {
		for {

			//实现拆包
			//读取
			//设置headbuffer
			headBuffer:=make([]byte,8)

			//先读取，再拆分,将获取的数据包读取到headbuffer
			n,err:=io.ReadFull(c.conn,headBuffer)
			if err != nil {
				fmt.Println("io.ReadFull err",err)
				return
			}
			fmt.Printf("读取数据头的长度:%d\n", n)
			dp:=NewDataPack{}
			//将数据头进行拆分
			headmsg,err:=dp.UnPack(headBuffer)
			if err != nil {
				fmt.Println("dp.UnPack err",err)
				return
			}
			fmt.Printf("数据头拆包后的数据详情：%v\n",headmsg)
			//获取数据长度
			datalen:=headmsg.GetDataLen()
			if datalen==0 {
				fmt.Printf("数据长度为0，无需读取, msgid:%d\n", headmsg.GetMsgId())
				continue
			}

			databuffer:=make([]byte,datalen)

			n,err=io.ReadFull(c.conn,databuffer)

			fmt.Printf("Server <===== Client, data:%s,cnt:%d, msgid:%d\n", databuffer, n, headmsg.GetMsgId())

			if err != nil {
				fmt.Println("dp.UnPack err",err)
				return
			}

			headmsg.SetData(databuffer)
			//将客户端传来的数据，传递给封装的Request方法，这样从Request方法中可以找到所需求不同的数据，data，len，conn
			req:= NewRequest(c,headmsg)

			//刚刚将传递给请求Request中封装的数据，通过req直接传递给router，router中封装了三个方法PreHandle，Handle，以及PostHandle方法，
			// 这样就能将用户实现的具体业务从框架中剥离出来，实现框架和业务具体实现的分离
			//这些框架中封装的方法又被server_main.go所继承，进行重写，由于就近原则，就可以在server_main.go中实现具体的方法
			//c.callback(req)
			//c.router.PreHandle(req)
			//c.router.Handle(req)
			//c.router.PostHandle(req)
			//多路由
			c.msghandler.DoMsgRouter(req)
		}

	}()
}
func (c *Connection)Send(data []byte,msgid uint32) (int,error){

	message:=NewMessage(data,uint32(len(data)),msgid)

	dp:=NewDataPack{}
	sendData,err:=dp.Pack(message)

	if err != nil {
		fmt.Println("dp.Pack err",err)
		return -1,err
	}

	n, err := c.conn.Write(sendData)
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