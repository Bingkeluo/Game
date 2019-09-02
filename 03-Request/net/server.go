package net

import (
	"../iface"
	"fmt"
	"net"
	"strings"
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
}

func userBussiness(req iface.IRequest){
	conn:=req.GetConnection()
	data:=req.GetData()
	//用户业务处理
	writeBackInfo := strings.ToUpper(string(data))

	//将回写的操作写到一个方法，
	n,err:=conn.Send([]byte(writeBackInfo))
	if err != nil {
		fmt.Println("conn.Send err:", err)
		return
	}
	fmt.Println("Server ====> Client, len:", n, ", buf :", writeBackInfo)
}

func NewSerVer (name string) iface.IServer{
	return &Server{
		IP:      "0.0.0.0",//监听所有的端口
		Port:    8848,
		Name:    name,
		Version: "tcp4",
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
	tcpaddr, err:=net.ResolveTCPAddr(s.Version,address)
	if err != nil {
		fmt.Println("Server Start err:",err)
		return
	}

	listener,err:=net.ListenTCP(s.Version,tcpaddr)
	if err != nil {
		fmt.Println("net.ListenTCP err:",err)
		return
	}
	var cid uint32
	cid=0
	go func() {
		for {
			//监听
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				return
			}
				//1.得到tcpconn，封装自己的Connection
				myconnecttion:=NewConnection(tcpConn,cid,userBussiness)
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
//方法
//1.启动start
//2.停止Stop
//3.服务Serve