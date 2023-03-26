package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"serveLine/models"
)

type GetToken struct {
	beego.Controller
}

func (k *GetToken) Get() {
	models.GetToken()
	k.Data["json"] = "请求了"
	err := k.ServeJSON()
	if err != nil {
		return
	}
}
