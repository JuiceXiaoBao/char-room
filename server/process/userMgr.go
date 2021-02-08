package process

import "fmt"

//因为UserMgr实例在服务器端有且只有一个
//因为在很多的地方都会使用到，因此定义为全局变量
var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{onlineUsers: make(map[int]*UserProcess, 1024)}
}

//完成对onlineUsers添加
func (this *UserMgr) AddonlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelateonlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map中取出一个值，待检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明当前查找的用户不在线
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
