package main

/*
//写一个函数，完成登录
func login(userId int,userPwd string)(err error) {

	////下一个就要开始定协议..
	//fmt.Printf("userId=%d userPwd=%v\n",userId,userPwd)
	//return nil

	//1.连接到服务器端
	conn,err:=net.Dial("tcp","localhost:8889")
	if err!=nil {
		fmt.Println("net.dial err=",err)
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
	loginMes.UserPwd= userPwd

	//4.将loginMes序列化
	data,err:=json.Marshal(loginMes)
	if err!=nil {
        fmt.Println("json.Marshal err=",err)
        return
	}

	//5.把data赋给mes.data字段
	mes.Data=string(data)

	//6.将mes进行序列化
	data,err=json.Marshal(mes)
	if err!=nil {
		fmt.Println("json.Marshal err=",err)
		return
	}

	//7.data就是要发送的消息
	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
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

	//发送消息本身
	_,err=conn.Write(data)
	if err!=nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度=%d 内容=%s",len(data),string(data))


	//time.Sleep(time.Second*20)
	//fmt.Println("休眠了20秒..")
	//这里需要处理服务器端返回的消息
	mes,err= utils.readpkg(conn)
	if err!=nil {
		fmt.Println("readPkg(conn)err=",err)
	}
	//mes就是
	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code==200 {
		fmt.Println("登陆成功！")
	}else if loginResMes.Code==500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

*/
