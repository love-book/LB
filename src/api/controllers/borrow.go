package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"common"
)

type ExchangeController struct {
	beego.Controller
}

func (this *ExchangeController) StartExchange(http http.Request) {
	uid := http.PostForm.Get("uid")
	bookid := http.PostForm.Get("bookid")
	if len(uid) == 0 || len(bookid) == 0 {
		this.Data["json"] = common.ErrInputData
		this.ServeJSON()
		this.StopRun()
	}


}