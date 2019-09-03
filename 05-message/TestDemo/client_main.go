package main


import (
	"fmt"
	"net"
	"time"
)

func main() {
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
		time.Sleep(time.Second)
		fmt.Println(data)
	}
}
