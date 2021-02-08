package processwhx

import (
	"char_room/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	//显示即可
	//1.反序列化mes.data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println("---------艹---------")
}

func outputOppositeMes(mes *message.Message) {
	var oppositeMes message.OppoiteMes
	err := json.Unmarshal([]byte(mes.Data), &oppositeMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对用户\t%d说：\t%s", oppositeMes.UserId, oppositeMes.OppositeUserId, oppositeMes.Content)
	fmt.Println(info)
	fmt.Println("-------kiao--------")
}
