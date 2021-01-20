package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id int
	Name string
	Password string
	Status int
	AddTime int64
	Mobile string
	Avatar string
}
func init()  {
	orm.RegisterModel(new(User))
}

// 判断手机号是否注册
func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "Mobile")

	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}

	return true
}

// 保存用户
func UserSave(mobile string, password string) error {
	o := orm.NewOrm()
	var user User
	user.Name = ""
	user.Password = password
	user.Mobile = mobile
	user.Status = 1
	user.AddTime = time.Now().Unix()
	_, err := o.Insert(&user)
	return err
}