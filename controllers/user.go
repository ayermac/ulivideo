package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	"regexp"
	"strconv"
	"strings"
	"ulivideoapi/models"
)

// Operations about Users
type UserController struct {
	CommonController
}


// 注册用户
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

// 用户登录
// @router /login/do [*]
func (this *UserController) LoginDo()  {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

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

	uid, name := models.IsMobileLogin(mobile, Md5V(password))
	if uid != 0 {
		this.ReturnSuccess("登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
	} else {
		this.ReturnError(4004, "手机号或密码不正确")
	}
}



//批量发送通知消息
// @router /send/message [*]
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")
	
	if uids == "" {
		this.ReturnError(4001, "请填写接收人~")
	}
	if content == "" {
		this.ReturnError(4002, "请填写发送内容")
	}

	messageId, err := models.SendMessageDo(content)
	if err == nil {
		uidConfig := strings.Split(uids, ",")
		for _, v := range uidConfig {
			userId, _ := strconv.Atoi(v)
			_ = models.SendMessageUserToMq(userId, messageId)
		}

		this.ReturnSuccess("发送成功", "", 1)
	}
	this.ReturnError(500, "发送失败，请联系客服")
}