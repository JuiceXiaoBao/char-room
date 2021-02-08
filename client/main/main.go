package main

import (
	"char_room/client/processwhx"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for true {
		fmt.Println("------------------欢迎登陆多人聊天系统------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天系统")
			fmt.Println("请输入用户的id号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%v\n", &userPwd)
			//完成登录
			//1.创建一个UserProcess的实例
			up := &processwhx.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字（nickname）：")
			fmt.Scanf("%s\n", &userName)
			//2.调用UserProcess，完成注册的请求
			up := &processwhx.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop=false
		case 3:
			fmt.Println("退出系统")
			//loop=false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入！")

		}
	}

	////根据用户的输入，显示新的提示信息
	//if key== 1 {
	//	//说明用户要登陆
	//	fmt.Println("请输入用户的id号")
	//	fmt.Scanf("%d\n",&userId)
	//	fmt.Println("请输入用户的密码")
	//	fmt.Scanf("%v\n",&userPwd)
	//	//先把登陆的函数，写到另外一个文件,写到login.go
	//	//这里需要重新调用
	//	login(userId,userPwd)
	//
	//}else if key==2 {
	//	fmt.Println("进行用户注册的逻辑...")
	//}

}
