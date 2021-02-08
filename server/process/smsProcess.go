package process

import (
	"char_room/common/message"
	"char_room/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的onlineUserS map[int]*UserProcess
	//将消息转发取出

	//取出mes的内容SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//这里过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个Tranfer实例
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}

func (this *SmsProcess) SendOppositeMes(mes *message.Message) {
	var oppositeMes message.OppoiteMes
	err := json.Unmarshal([]byte(mes.Data), &oppositeMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里过滤掉自己
		if id == oppositeMes.UserId {
			continue
		} else if id == oppositeMes.OppositeUserId {
			this.SendMesToEachOnlineUser(data, up.Conn)
		}
	}
}
