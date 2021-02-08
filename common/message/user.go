package message

//定义一个用户的结构体

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的json字符串的key和结构体的字段对应的tag名字一致！！
	UserId         int    `json:"user_id"`
	UserPwd        string `json:"user_pwd"`
	UserName       string `json:"user_name"`
	UserStatus     int    `json:"user_status"`      //用户状态
	Sex            string `json:"sex"`              //性别
	OppositeUserId int    `json:"opposite_user_id"` //对方用户id
}
