package models

import (
	"github.com/beego/beego/v2/client/orm"
	"ming/utils"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

// User用户
type User struct {
	ID         int    `orm:"column(id)"`
	EmployeeID string `orm:"column(employee_id);size(32)"`
	Username   string `orm:"size(64)"`
	Avatar     string `orm:"size(1024)"`
	Password   string `orm:"size(1024)"`
	Nickname   string `orm:"size(64)"`
	Gender     int
	PhoneNum   string `orm:"column(phone_num);size(32)"`
	Address    string `orm:"size(128)"`
	Email      string `orm:"size(64)"`
	Department string `orm:"size(64)"`
	Post       string `orm:"size(64)"`
	Role       int
	Status     int
	CreatedAt  time.Time `orm:"auto_now_add"`
	UpdatedAt  time.Time `orm:"auto_now"`
}

// 注册model
func init() {
	orm.RegisterModel(new(User))
}

// GetUserByName 通过用户名获取用户
func GetUserByUsername(username string) *User {
	user := &User{
		Username: username,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Username"); err == nil {
		return user
	}
	return nil

}

// ValidPassword 检查密码是否正确
func (u *User) ValidPassword(password string) bool {
	//u.Password 是数据库中保存的pwd值
	return utils.VerifyPassword(password, u.Password)
}

// GetUserList 查询特定用户
func GetUserList(q string) []*User {
	var users []*User
	queryset := orm.NewOrm().QueryTable(&User{})
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("username__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("department__icontains", q)
		cond = cond.Or("gender__icontains", q)
		cond = cond.Or("post__icontains", q)
		cond = cond.Or("role__icontains", q)
		cond = cond.Or("status__icontains", q)
		queryset = queryset.SetCond(cond)
	}

	queryset.All(&users)
	return users
}

// GenderText 转换性别显示文字
func (u *User) GenderText() string {
	switch u.Gender {
	case 0:
		return "女"
	case 1:
		return "男"
	}
	return "未知"
}

// StatusText 转换状态的显示文字
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "已禁用"
	case 1:
		return "使用中"
	}
	return "未知"
}

// RoleText 转换角色的显示文字
func (u *User) RoleText() string {
	switch u.Role {
	case 0:
		return "系统管理员"
	case 1:
		return "部门管理员"
	case 2:
		return "普通用户"

	}
	return "未知"
}
