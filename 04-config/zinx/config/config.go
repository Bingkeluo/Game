package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//定义一个配置文件结构
type config struct {
	IP string
	PORT uint32
	Name string
	Version string
}
//在init函数中加载LoadConfig函数,init函数会先于main函数执行，而且init函数不能显式调用
func init(){
	err:=LoadConfig()
	if err != nil {
	fmt.Println("zinx配置文件失败。。。。:", err)
	os.Exit(-1)
	}
	fmt.Println("++++++++++++++配置文件信息如下:++++++++++++")
	fmt.Printf("%v\n", GlobalConfig)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")
}



//定义一个全局的配置文件结构，用于接收从配置文件中读取的数据
var GlobalConfig config

//2. 加载配置文件
func LoadConfig() error {
	//1.读取配置文件
	//基于server_main.go目录进行寻找
	fmt.Println("开始读取配置文件。。。")
	configInfo,err:=ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		return err
	}
	//2. 反序列为Config结构
	//3.配置文件全局唯一，需要定义一个GlobalConfig字段，赋值解析出来的数据
	err=json.Unmarshal(configInfo,&GlobalConfig)
	if err != nil {
		return err
	}
	fmt.Println("读取配置文件成功。。。")
	return nil
}


