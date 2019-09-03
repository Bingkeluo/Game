package net

import (
	"fmt"
	"zinx/v5.1-message-datapack/zinx/iface"
)
//封装一个原生的Router结构体
type Router struct {

}
//将原生的结构体绑定三个方法，这三个方法都是为了后面在server_main.go中进行具体的实现所封装的，可以利用继承将三个方法进行重写，或者直接调用
//绑定三个方法 //处理业务之前做准备
func (r Router)PreHandle(req iface.IRequest){
	fmt.Println("PreHandle called")
}
//真正的处理业务
func (r Router)Handle(req iface.IRequest){
	fmt.Println("Handle called")
}
//处理业务之后做清理
func (r Router)PostHandle(req iface.IRequest){
	fmt.Println("PostHandle called")
}