package iface

import "net"

//1.start()===>读写方法
//2.send()===>向conn发送数据
//3.Stop()
//4.GetConnId ===>每个连接都有自己的Id
//5.GetTcpConn()===> *netTcpConn

type IConnection interface {
	Start()
	Send([]byte ,uint32) (int,error)//增加id字段
	Stop()
	GetConnId() uint32
	GetTcpConn() *net.TCPConn //原生的链接
}
//定义一个回调函数由用户指定，处理用户指定的业务
//路由
//type CallBackFunc func(request IRequest)
