package main

import (
	"char_room/server/model"
	"fmt"
	"net"
)

/*
func readpkg(conn net.Conn)(mes message.Message,err error)  {
	buf:=make([]byte,8096)
	fmt.Println("读取客户端发送的数据...")

	n,err:=conn.Read(buf[0:4])
	if n!=4||err!=nil {
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[0:4]读到的长度转成一个uint32类型
	var pkglen uint32
	pkglen=binary.BigEndian.Uint32(buf[0:4])

	//根据pkglen读取消息内容
	//从conn里读取pkglen个字节到缓存buf里面去
	n,err=conn.Read(buf[:pkglen])
	if n!=int(pkglen)||err!=nil {
		fmt.Println("conn.read fail err=",err)
		return
	}

//把pkg反序列化成 -> message.Message
	err=json.Unmarshal(buf[:pkglen],&mes)
	if err!=nil {
		fmt.Println("Json.Unmarsha err=",err)
		return
	}

	return

}

func writePkg(conn net.Conn,data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen=uint32(len(data))


	var buf []byte
	buf=make([]byte,4)
	binary.BigEndian.PutUint32(buf,pkgLen)

	//发送长度
	n,err:=conn.Write(buf)
	if n!=4||err!=nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//发送data本身
	n,err=conn.Write(data)
	if n!=int(pkgLen)||err!=nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	return

}
*/

/*
//编写一个函数serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn,mes *message.Message) (err error) {
	//1.先从mes中取出mes.data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err=json.Unmarshal([]byte(mes.Data),&loginMes)
	if err!=nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//如果用户id=100，密码=123456，认为合法，否则不合法
	if loginMes.UserId==100&&loginMes.UserPwd=="123456" {
			//合法
		loginResMes.Code = 200
	}else {
		//不合法
		loginResMes.Code = 500 //500状态码表示用户不存在
		loginResMes.Error="饿货！用户不存在！来点士力架吧..."
	}

	//3.将loginResMes序列化
	data,err:=json.Marshal(loginResMes)
	if err!=nil {
		fmt.Println("json.marshal fail",err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data=string(data)

	//5.对resMes进行序列化，准备发送
	data,err=json.Marshal(resMes)
	if err!=nil {
		fmt.Println("json.Marshal fail",err)
		return
	}

//6.发送data，我们将其封装到writePkg函数
   err=writePkg(conn,data)
   return
}

*/

////编写一个ServerProcessMes函数
////功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录
//		err=serverProcessLogin(conn,mes)
//		case message.RegisterMesType:
//		//处理注册
//		default:
//		fmt.Println("消息类型不存在，无法处理...")
//	}
//	return
//}

//处理和客户端的通讯
func process1(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()

	//这里调用总控
	processor := &Processor{Conn: conn}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误=", err)
	}
}

//这里编写一个函数完成对userDao的初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool) //pool本身就是一个全局变量
}

func main() {
	//但服务器启动时，就去初始化redis
	initPool("localhost:6379", 16, 0, 300)
	initUserDao()
	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen,Accept err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯...
		go process1(conn)
	}
}
