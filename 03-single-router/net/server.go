package net

import (
	"zinx/03-single-router/iface"
	"fmt"
	"net"
)

type Server struct {
	//属性
	//1.Ip
	IP string
	//2.Port
	Port uint32
	//3.Name
	Name string
	//4.Version
	Version string//tcp4，tcp6
	//5.Router
	//来自封装的接口，目的将从web端获取的路由，传递给connection
	Router iface.IRouter

}

func userBussiness(req iface.IRequest){

}

func NewSerVer (name string) iface.IServer{
	return &Server{
		IP:      "0.0.0.0",//监听所有的端口
		Port:    8848,
		Name:    name,
		Version: "tcp4",
		//创建一个空的路由，即使用户不传递过来自定义的实现方法，程序也能正常运行
		Router:&Router{},
	}
}

func (s *Server)Start(){

	fmt.Println("[Server Start]...")
	//TODO
	//socket监听

	//old
	//l := net.Listen("tcp", ":8888")
	//c := l.Accept()
	//c.Read()

	address:=fmt.Sprintf("%s:%d",s.IP,s.Port)
	//func ResolveIPAddr(network, address string) (*IPAddr, error) {
	//初始化tcpaddr
	tcpaddr, err:=net.ResolveTCPAddr(s.Version,address)
	if err != nil {
		fmt.Println("Server Start err:",err)
		return
	}
	//创建监听者
	listener,err:=net.ListenTCP(s.Version,tcpaddr)
	if err != nil {
		fmt.Println("net.ListenTCP err:",err)
		return
	}
	var cid uint32
	cid=0
	go func() {
		for {
			//监听，创建通道
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				return
			}
			//1.得到tcpconn，封装自己的Connection通道，将请求细分
			myconnecttion:=NewConnection(tcpConn,cid,s.Router)
			//参数tcpconn：将原生的通道传递过去，可以封装创建属于自己的通道，
			//参数cid：绑定指定的id，这样可以区分请求，
			//参数s.router：通过绑定的接收者s，将其中的路由传递给我们封装的通道。
			cid++
			//2. 启动conn.start，
			//server只负责管理连接，具体的业务处理，由conn负责
			go myconnecttion.Start()

		}
	}()

}

func (s *Server)Stop(){
	fmt.Println("[Server Stop]...")
}

func (s *Server)Server(){
	fmt.Println("[Server Server]...")
	s.Start()
	for {;}
}

func (s *Server)AddRouter(router iface.IRouter){
	s.Router=router
}
//方法
//1.启动start
//2.停止Stop
//3.服务Serve