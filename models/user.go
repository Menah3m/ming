package models

import (
	"database/sql"
	"fmt"

	"ming/utils"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

const (
	sqlQueryByName  = "select id,username,password from user where username=?"
	sqlQueryAllUser = "select id,employee_id,username,nickname,email,department,post,role,status,gender,address from user"
)

// User用户
type User struct {
	ID         int       `field:"id"`
	EmployeeID string    `field:"employee_id"`
	Username   string    `field:"username"`
	Avatar     string    `field:"avatar"`
	Password   string    `field:"password"`
	Nickname   string    `field:"nickname"`
	Gender     int       `field:"gender"`
	PhoneNum   string    `field:"phone_num"`
	Address    string    `field:"address"`
	Email      string    `field:"email"`
	Department string    `field:"department"`
	Post       string    `field:"post"`
	Role       int       `field:"role"`
	Status     int       `field:"status"`
	CreatedAt  time.Time `field:"created_at"`
	UpdatedAt  time.Time `field:"updated_at"`
}

// GetUserByName 通过用户名获取用户
func GetUserByUsername(username string) *User {
	user := &User{}
	if err := db.QueryRow(sqlQueryByName, username).Scan(&user.ID, &user.Username, &user.Password); err == nil {
		return user
	}
	return nil
}

// ValidPassword 检查密码是否正确
func (u *User) ValidPassword(password string) bool {
	return u.Password == utils.EncodeMD5(password)
}

// GetUserList 获取用户
func GetUserList(q string) []*User {
	//对查询参数做相应的处理：转义、去除多余空格等
	q = utils.Like(q)
	users := make([]*User, 0, 10)
	sqlQuery := sqlQueryAllUser
	var (
		rows *sql.Rows
		err  error
	)
	//定义空接口切片，用来存放不确定个数的查询参数
	queryParam := []interface{}{}
	if q != "" {
		sqlQuery += " WHERE  employee_id like ? ESCAPE '/' OR username like ? ESCAPE '/' OR nickname like ? ESCAPE '/' OR email like ? ESCAPE '/'  OR department like ? ESCAPE '/'  OR post like ? ESCAPE '/'  OR  role like ?  ESCAPE '/' OR status like ? ESCAPE '/'  OR gender like ? ESCAPE '/'  OR address like ? ESCAPE '/' "
		queryParam = append(queryParam, q, q, q, q, q, q, q, q, q, q)
	}
	//查询
	rows, err = db.Query(sqlQuery, queryParam...)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.EmployeeID, &user.Username, &user.Nickname, &user.Email, &user.Department, &user.Post, &user.Role, &user.Status, &user.Gender, &user.Address)
		if err != nil {
			fmt.Println("scan err:", err)
			return nil
		}
		users = append(users, user)
	}
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
