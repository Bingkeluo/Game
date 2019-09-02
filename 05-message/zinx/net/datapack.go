package net

import (
	"bytes"
	"encoding/binary"
	"zinx/05-message/zinx/iface"
)

//负责封包和拆包

type DataPack struct {
}

//封包函数
func (dp *DataPack)Pack(msg iface.IMessage)([]byte,error){
	//打包过程：Message  ====》 字节流
	//TODO
	//1. 先获取消息的长度，内容，id
	data:=msg.GetData()
	len:=msg.GetDataLen()
	msgId:=msg.GetMsgId()

	var buff bytes.Buffer

	//2. 写消息头
	//a. 写长度
	err:=binary.Write(&buff,binary.LittleEndian,len)
	if err != nil {
		return nil,err
	}
	//b. 写msgid
	err=binary.Write(&buff,binary.LittleEndian,msgId)
	if err != nil {
		return nil,err
	}
	//3. 写消息体
	err=binary.Write(&buff,binary.LittleEndian,data)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

//拆包函数

//在connection中使用时，会读取两次
//1. 第一次会读取固定8字节的长度，然后调用Unpack
//2. 第二次会读取真实长度的数据
func (dp *DataPack)UnPack(data []byte)(iface.IMessage,error){
	//进来的data就是8字节
	//创建一个reader，读取data
		reader:=bytes.NewReader(data)

	//定义一个空的message结构，用于接收反序列化的数据
	var message Message
	//拆包主要是解析出两个内容：真实传递数据的长度，消息id
	//
	//	//1. 读取数据头, 获得数据长度
	//	//func Read(r io.Reader, order ByteOrder, data interface{}) error
	//	//err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &message.len) //错误做法:
	err:=binary.Read(reader,binary.LittleEndian,&message.len)
	if err != nil {
		return nil,err
	}
	//2. 读取数据头，获得消息类型
	err=binary.Read(reader,binary.LittleEndian,&message.msgid)
	if err != nil {
		return nil,err
	}
	return &message,nil
}
