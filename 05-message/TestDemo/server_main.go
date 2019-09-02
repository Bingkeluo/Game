package main


import (
	"../../05-message/zinx/net"
	"fmt"
	"strings"

	//"strings"
	"zinx/05-message/zinx/iface"
)
//具体业务应该由使用框架的人传入
//继承框架中的Router结构体，对其绑定的结构体实现方法重写。
type TestRouter struct {
	net.Router
}

//用户重写三个函数实现自己的业务
func (r *TestRouter)PreHandle(req iface.IRequest){
	fmt.Println("用户实现自己的PreHandle")
}

func (r *TestRouter)Handle(req iface.IRequest){
	fmt.Println("用户实现自己的Handle")
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

func (r *TestRouter)PostHandle(req iface.IRequest){
	fmt.Println("用户实现自己的PostHandle")
}
func main() {
	server:=net.NewSerVer("zinx v1.0")
	//方法重写，子类方法会将父类方法进行覆盖。
	//将用户实现的具体方法传递给server，进行赋值s.Router=router，将子类 的&TestRouter{}方法赋值给了父类s.Router,实现了方法的重写
	server.AddRouter(&TestRouter{})
	server.Server()
}
