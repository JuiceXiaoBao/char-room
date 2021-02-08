package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	OppositeType            = "Opposite"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	Useroffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息...后面需要再增加

type LoginMes struct {
	UserId   int    `json:"user_id"`   //用户id
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名
}

type LoginResMes struct {
	Code    int         `json:"code"` //返回状态码 500表示该用户未注册 200表示登录成功
	UserIds []int       //增加字段，保存用户id的切片
	Error   interface{} `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是user结构体
}

type RegisterResMes struct {
	Code  int         `json:"code"`  //返回状态码 400表示该用户已经占用 200表示注册成功
	Error interface{} `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"user_id"` //用户id
	Status int `json:"status"`  //用户状态
}

//增加一个SmsMes  //发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体 继承
}

//私聊结构体
type OppoiteMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体
}
