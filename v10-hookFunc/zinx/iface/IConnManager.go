package iface

type IConnManager interface {
	AddConn( IConnection) //增加连接
	Remove(uint32) //删除连接
	GetConn(uint32) IConnection //给定cid，返回连接句柄
	GetConnCount() uint32 //获取当前所有的连接的总数
	ClearConn() //清除所有的连接
}