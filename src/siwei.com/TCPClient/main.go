package main

import "github.com/astaxie/beego"
//import "net"
import (
	"fmt"
	"runtime"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//this.Ctx.WriteString("hello world")
	go sendMessageToServer(this)
	//fmt.Println("test")
	runtime.Gosched()
}

func main()  {

	beego.Router("/",&MainController{})
	beego.Run()
}

func checkError(err error,info string) (res bool) {

	if(err != nil){
		fmt.Println(info+"  " + err.Error())
		return false
	}
	return true
}

func sendMessageToServer(this *MainController) () {
	fmt.Println("send start")
	this.Ctx.WriteString("test")
	/*service := "192.168.1.106:5018"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if(err != nil) {
		fmt.Println(err);
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err!= nil {
		fmt.Println(err)
	}
	_, err = conn.Write([]byte("test"))

	if err!= nil {
		fmt.Println(err)
	}*/
}