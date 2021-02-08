package utils

import (
	"char_room/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用的缓冲
}

func (this *Transfer) Readpkg() (mes message.Message, err error) {
	//buf:=make([]byte,8096)
	fmt.Println("读取客户端发送的数据...")

	n, err := this.Conn.Read(this.Buf[0:4])
	if n != 4 || err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[0:4]读到的长度转成一个uint32类型
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkglen读取消息内容
	//从conn里读取pkglen个字节到缓存buf里面去
	n, err = this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.read fail err=", err)
		return
	}

	//把pkg反序列化成 -> message.Message
	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("Json.Unmarsha err=", err)
		return
	}

	return

}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	//var buf []byte
	//buf=make([]byte,4)
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	//发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}
