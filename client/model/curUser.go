package model

import (
	"char_room/common/message"
	"net"
)

//因为在客户端很多地方都会使用到curUser，我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
