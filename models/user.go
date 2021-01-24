package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
	redisClient "ulivideoapi/services/redis"
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

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
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

// 手机号登录
func IsMobileLogin(mobile string, password string) (int, string)  {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}

	return user.Id, user.Name
}

func GetUserInfo(userId int) (UserInfo, error) {
	o := orm.NewOrm()
	var userInfo UserInfo
	err := o.Raw("SELECT id, name, add_time, avatar FROM user where id = ? limit 1", userId).QueryRow(&userInfo)

	return userInfo, err
}

// 从Redis获取用户信息
func RedisGetUserInfo(userId int) (UserInfo, error) {
	var userInfo UserInfo
	conn := redisClient.PoolConnect()
	defer conn.Close()

	rKey := "user:info:" + strconv.Itoa(userId)

	//判断redis中是否存在
	exists, err := redis.Bool(conn.Do("exists", rKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", rKey))
		err = redis.ScanStruct(res, &userInfo)
	} else {
		o := orm.NewOrm()
		err := o.Raw("SELECT id, name, add_time, avatar FROM user where id = ? limit 1", userId).QueryRow(&userInfo)
		if err == nil {
			_, err := conn.Do("hmset", redis.Args{rKey}.AddFlat(userInfo)...)
			if err == nil {
				conn.Do("expire", rKey, 86400)
			}
		}

		if err != nil {
			panic(err)
		}
	}
	return userInfo, err
}
