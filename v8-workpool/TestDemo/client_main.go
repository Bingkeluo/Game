package main

import (
	"fmt"
	"io"
	"net"
	"time"
	net2 "zinx/v7-reader-write/zinx/net"
)

func main1() {
	conn,err:=net.Dial("tcp",":8848")
	if err != nil {
		fmt.Println("net.Dial err:",err)
		return
	}

	data:=[]byte("hello world")
//追加？
	//write
	for{
		n,err:=conn.Write(data)
		if err != nil {
			fmt.Println("conn.Write err:",err)
			continue
		}
		fmt.Println("conn write :",n,"Data:",data)

		n,err=conn.Read(data)

		if err != nil {
			fmt.Println("conn.Read err:",err)
			continue
		}
		time.Sleep(time.Second*10)
		fmt.Println(data)
	}
}

func main() {
	//封包，发送
	//把多个包黏在一起，一起发送
	//1. 准备数据（封包）
	data1:=[]byte("你好")
	data2:=[]byte("hello world")
	data3:=[]byte("国庆节不回家")

	//a. 创建message
	datamsg1:=net2.NewMessage(data1,uint32(len(data1)),0)
	datamsg2:=net2.NewMessage(data2,uint32(len(data2)),1)
	datamsg3:=net2.NewMessage(data3,uint32(len(data3)),2)
	//b. 对message进行封包
	//创建封包接收者
	dp:=net2.NewDataPack{}
	msg1,err:=dp.Pack(datamsg1)
	msg2,err:=dp.Pack(datamsg2)
	msg3,err:=dp.Pack(datamsg3)

	//将三个消息的字节流拼接到一起，一次性发送给服务器
	//切片追加到切片
	msgInfo:= append(msg1, msg2...)
	msgInfo=append(msgInfo,msg3...)


	//2. 发送
	conn,err:=net.Dial("tcp","127.0.0.1:8848")
	if err != nil {
		fmt.Printf("dail err:%v\n",err)
		return
	}

			go func() {
				for {

					cnt, err := conn.Write(msgInfo)
					if err != nil {
						fmt.Printf("dail err:%v\n", err)
						return
					}
					fmt.Printf("dail err:%v\n", err)

					fmt.Println("Client ====> Server cnt:", cnt)
					//??????为啥
					//time.Sleep(1 * time.Second)
				}
			}()


				for{
					dp:=net2.NewDataPack{}
					headBuffer:=make([]byte,dp.GetDataHeadLen())

					_,err=io.ReadFull(conn,headBuffer)
					if err!=nil{
						fmt.Println("io.ReadFull err",err)
						return
					}

					message,err:=dp.UnPack(headBuffer)

					if err!=nil{
						fmt.Println("dp.UnPack err",err)
						return
					}
					//数据长度
					datalen:=message.GetDataLen()

					if datalen==0 {
						fmt.Println("数据长度为0")
						continue
					}

					data:=make([]byte,datalen)
					n,err:=io.ReadFull(conn,data)
					if err!=nil{
						fmt.Println("io.ReadFull 2 err",err)
						return
					}
					fmt.Printf("datalen:%d,data:%s",n,data)
				}

	
}