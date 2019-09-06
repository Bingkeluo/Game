package net

import (
	"fmt"
	"sync"
	"zinx/v10-hookFunc/zinx/iface"
)
//建立管理结构体
type ConnManager struct {
	//用map创建管理所有连接的集合
	connS map[uint32]iface.IConnection
	connLock sync.RWMutex//创建读写锁
}

func NewConnManager() iface.IConnManager{
	return &ConnManager{
		connS:  make(map[uint32]iface.IConnection),
	}
}
//增加连接
func (cm *ConnManager)AddConn( conn iface.IConnection){
	cid:=conn.GetConnId()
	//建立新连接
	fmt.Println("增加新连接:",cid)

	//多个goroutine操作map时，一定要加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//判断是否含有连接
	if _,ok:=cm.connS[cid];ok{
		fmt.Println("当前连接已存在，无需添加",cid)
		return
	}

	cm.connS[cid]=conn
	fmt.Println("dukdudk")
}

//删除连接o
func (cm *ConnManager)Remove(cid uint32){
	fmt.Println("删除连接：cid:",cid)
	//加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	//删除id
	delete(cm.connS,cid)
}
//获取连接,给定id，返回句柄
func (cm *ConnManager)GetConn(cid uint32)iface.IConnection{
	fmt.Println("获取连接：",cid)
	cm.connLock.RLock()
	defer  cm.connLock.RUnlock()

	return cm.connS[cid]
}
//获取连接总数
func (cm *ConnManager)GetConnCount( )uint32{
	fmt.Println("获取当前连接数量")
	return uint32(len(cm.connS))

}
//清除所有连接
func (cm *ConnManager)ClearConn( ){
	fmt.Println("清楚所有连接")
	cm.connLock.Lock()
	defer  cm.connLock.Unlock()

	//删除所有连接
	for cid,conn:=range cm.connS  {
		//先停再删除
		conn.Stop()
		delete(cm.connS,cid)
	}
}

