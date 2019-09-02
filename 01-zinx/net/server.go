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
go func() {
	for {
		//监听
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("listener.AcceptTCP err:", err)
			return
		}

		go func() {
			for {
				buf := make([]byte, 4096)

				n, err := tcpConn.Read(buf)
				if err != nil {
					fmt.Println("tcpConn.Read err:", err)
					return
				}

				fmt.Println("Server <==== Client, len:", n, ", buf :", string(buf[:n]))
				writeBackInfo := strings.ToUpper(string(buf[:n]))

				n, err = tcpConn.Write([]byte(writeBackInfo))
				if err != nil {
					fmt.Println("tcpConn.Write err:", err)
					return
				}

				fmt.Println("Server ====> Client, len:", n, ", buf :", writeBackInfo)
			}

		}()
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