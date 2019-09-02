package iface

/*
IRequest 接⼝：
实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request⾥
*/

type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetLen() uint32
}