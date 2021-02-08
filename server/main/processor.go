package main

import (
	"char_room/common/message"
	"char_room/server/process"
	"char_room/server/utils"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//看看是否能接收到客户端发送的群发消息
	fmt.Println("mes=", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
	//处理注册
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息.
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.OppositeType:
		//创建一个oppositeProcess实例完成私聊消息
		smsProcess := &process.SmsProcess{}
		smsProcess.SendOppositeMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {

	//循环读客户端发送的信息
	for {
		//这里将读取数据包，直接封装成一个函数readpkg（），返回message，err
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
			Buf:  [8096]byte{},
		}
		mes, err := tf.Readpkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常关闭...")
				return err
			} else {
				fmt.Println("readpkg err=", err)
				return err
			}
		}
		//fmt.Println("mes=",mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
