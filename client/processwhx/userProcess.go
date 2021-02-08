package processwhx

import (
	"char_room/common/message"
	"char_room/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	//暂时不需要声明字段...
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	////下一个就要开始定协议..
	//fmt.Printf("userId=%d userPwd=%v\n",userId,userPwd)
	//return nil

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.data就是要发送的消息
	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf []byte
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, pkgLen)

	//发送长度
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))

	//time.Sleep(time.Second*20)
	//fmt.Println("休眠了20秒..")
	//这里需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	mes, err = tf.Readpkg()
	if err != nil {
		fmt.Println("readPkg(conn)err=", err)
	}
	//mes就是
	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登陆成功！")
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//显示当前在线用户列表
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			//如果是不显示自己在线，可以增加以下代码
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端的onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1.显示我们的登录成功菜单[循环]...
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userid int, userPwd string, userName string) (err error) {
	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userid
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将loginMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错！err=", err)
	}

	mes, err = tf.Readpkg()
	if err != nil {
		fmt.Println("readPkg(conn)err=", err)
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功！请您重新登录！")
		return
	} else {
		fmt.Println(registerResMes.Error)
		return
	}

	return

}
