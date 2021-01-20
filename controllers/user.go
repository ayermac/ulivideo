package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"regexp"
	"ulivideoapi/models"
)

// Operations about Users
type UserController struct {
	CommonController
}


// @router /register/save [post]
func (this *UserController) SaveRegister() {
	var (
		mobile string
		password string
		//err error
	)

	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.ReturnError(4001, "手机号不能为空")
	}
	isOrNo, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isOrNo {
		this.ReturnError(4002, "手机号格式不正确")
	}

	if password == "" {
		this.ReturnError(4003, "密码不能为空")
	}

	status := models.IsUserMobile(mobile)
	if status {
		this.ReturnError(4005, "手机号已注册")
	} else {
		err := models.UserSave(mobile, Md5V(password))
		if err == nil {
			this.ReturnSuccess("注册成功", nil, 0)
		} else {
			this.ReturnError(5000, err)
		}
	}
}
