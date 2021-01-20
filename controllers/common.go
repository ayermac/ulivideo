package controllers

import (
    "crypto/md5"
    "encoding/hex"

    beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type CommonController struct {
    beego.Controller
}

type JsonStruct struct {
    Code int            `json:"code"`
    Msg interface{}     `json:"msg"`
    Items interface{}   `json:"items"`
    Count int64         `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct) {
    json = &JsonStruct{Code: code, Msg: msg, Items: items, Count: count}
    return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
    json = &JsonStruct{Code: code, Msg: msg}
    return
}

// md5加盐
func Md5V(password string) string  {
    h := md5.New()

    md5Code, _ := beego.AppConfig.String("md5code")
    h.Write([]byte(password + md5Code))
    return hex.EncodeToString(h.Sum(nil))
}
