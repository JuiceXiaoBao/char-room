package processwhx

import (
	"char_room/common/message"
	"char_room/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct{}

////发送私聊的消息
func (this *SmsProcess) SendOppositeMes(content string, oppositeid int) {
	//1.创建一个结构体
	var mes message.Message
	mes.Type = message.OppositeType

	//2.创建一个OppositeMes实例
	var oppositeMes message.OppoiteMes
	oppositeMes.OppositeUserId = oppositeid
	oppositeMes.UserId = CurUser.UserId
	oppositeMes.UserStatus = CurUser.UserStatus
	oppositeMes.Content = content

	//3.序列化smsMes
	data, err := json.Marshal(oppositeMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err.Error())
		return
	}

	//5.将mes发送给服务器..
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
		Buf:  [8096]byte{},
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
	}
	return

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content               //内容.
	smsMes.UserId = CurUser.UserId         //
	smsMes.UserStatus = CurUser.UserStatus //

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err.Error())
		return
	}

	//5.将mes发送给服务器..
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
		Buf:  [8096]byte{},
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
	}
	return
}
