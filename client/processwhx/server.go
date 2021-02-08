package processwhx

import (
	"char_room/common/message"
	"char_room/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面..
func showMenu() {
	fmt.Println("---------------恭喜xxx登录成功---------------")
	fmt.Println("---------------1.显示在线用户列表---------------")
	fmt.Println("---------------2.群聊用户发送消息---------------")
	fmt.Println("---------------3.私聊用户发送消息---------------")
	fmt.Println("---------------4.信息列表---------------")
	fmt.Println("---------------5.退出系统---------------")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string
	var oppositeid int
	var oppositecontent string

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你想对大家说点什么：")
		fmt.Scanf("%s\n", &content)
		smsMes := SmsProcess{}
		smsMes.SendGroupMes(content)
	case 3:
		fmt.Println("请输入你想私聊的用户：")
		fmt.Scanf("%v\n", &oppositeid)
		fmt.Println("请输入你想对该用户说些什么：")
		fmt.Scanf("%v\n", &oppositecontent)
		oppositeMes := SmsProcess{}
		oppositeMes.SendOppositeMes(oppositecontent, oppositeid)
	case 4:
		fmt.Println("信息列表")
	case 5:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确..")
	}
}

//和服务器端保持通讯
func serverProcessMes(Conn net.Conn) {
	//创建一个tranfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: Conn,
		Buf:  [8096]byte{},
	}

	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.Readpkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		//如果读取到消息，又是下一步处理逻辑
		//fmt.Println("mes=%v",mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线
			//1.取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json unmarshal err=", err)
				return
			}
			//2.把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		//处理
		case message.SmsMesType: //
			outputGroupMes(&mes)
		case message.OppositeType:
			outputOppositeMes(&mes)
		default:
			fmt.Println("服务器端返回了未知的消息类型")

		}
	}

}
