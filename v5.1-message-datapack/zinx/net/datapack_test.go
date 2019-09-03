package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//1. 每个测试文件需要以_test.go结尾,创建包名
//2. 每个测试文件需要引用testing包
//3. 每个测试的函数需要以Test开头。
func TestDataPack(t *testing.T) {
fmt.Println("TestDataPack called")

	//测试data pack 与unpack函数

	//server
	go func() {

		lister,err:=net.Listen("tcp","127.0.0.1:8888")
		if err != nil {
			t.Error("listen err",err)
			return
		}
		defer lister.Close()

		conn,err:=lister.Accept()
		if err != nil {
			t.Error("listen err",err)
			return
		}
		defer conn.Close()

		for{
			//两次读取
			//1.读取数据头
			//a.创建，数据头的缓存包
			headBuff:=make([]byte,8)

			//2. 拆包, 把数据的长度(N)，与消息的类型解析出来
			//c.Read()
			//a. 只能一次性读取，没有指定长度, 如果有网络延迟，可能没有读取我们需要的长度的数据就返回
			//b. 解决办法，使用io.ReadFull来读取数据，这个函数可以读取指定的buf长度的数据，如果未读取完毕，则不返回
			//b.按照headbuff获取数据头，差分出长度和id
			n, err :=io.ReadFull(conn,headBuff)
			if err != nil {
				t.Errorf("io.ReadFull err:%v\n",err)
				return
			}

			fmt.Printf("读取数据头的长度:%d\n", n)

			//拆包
			dp:=NewDataPack{}

			headmsg,err:=dp.UnPack(headBuff)
			if err != nil {
				t.Errorf("dp.UnPack err:%v\n",err)
				return
			}
			fmt.Printf("数据头拆包后的数据详情：%v\n",headmsg)

			//校验真实数据长度
			dataLen:=headmsg.GetDataLen()

			if dataLen==0 {
				fmt.Printf("数据长度为0，无需读取, msgid:%d\n", headmsg.GetMsgId())
				continue
			}

			//根据得到的数据长度拆分数据
			databuffer:=make([]byte,dataLen)

			//这里可以根据数据头第一次读取获取到数据长度，读取直接可以得到数据，不需要再进行拆包，每次读取数据只需要进行一次拆包。
			n, err =io.ReadFull(conn,databuffer)
			if err != nil {
				t.Errorf("io.ReadFull err:%v\n",err)
				return
			}
			fmt.Printf("Server <===== Client, data:%s,cnt:%d, msgid:%d\n", databuffer, n, headmsg.GetMsgId())

		}

	}()


	//client
	go func() {
		//封包，发送
		//把多个包黏在一起，一起发送
		//1. 准备数据（封包）
		data1:=[]byte("你好")
		data2:=[]byte("hello world")
		data3:=[]byte("国庆节不回家")

		//a. 创建message
		datamsg1:=NewMessage(data1,uint32(len(data1)),0)
		datamsg2:=NewMessage(data2,uint32(len(data2)),1)
		datamsg3:=NewMessage(data3,uint32(len(data3)),2)
		//b. 对message进行封包
		//创建封包接收者
		dp:=NewDataPack{}
		msg1,err:=dp.Pack(datamsg1)
		msg2,err:=dp.Pack(datamsg2)
		msg3,err:=dp.Pack(datamsg3)

		//将三个消息的字节流拼接到一起，一次性发送给服务器
		//切片追加到切片
		msgInfo:= append(msg1, msg2...)
		msgInfo=append(msgInfo,msg3...)


		//2. 发送
		conn,err:=net.Dial("tcp","127.0.0.1:8888")
		if err != nil {
		t.Errorf("dail err:%v\n",err)
			return
		}
		defer conn.Close()
		cnt,err:=conn.Write(msgInfo)
		if err != nil {
			t.Errorf("dail err:%v\n",err)
			return
		}

		fmt.Println("Client ====> Server cnt:", cnt)
	}()


	select {}
}
