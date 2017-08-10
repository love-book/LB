package admin

import (
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}


//json输出
func (this * baseController) renderJson(Json interface{}) {
	//Json := map[string]interface{}{"code":0,"msg": "成功!","data":"1111"}
	this.Data["json"] = Json
	this.ServeJSON()
	this.StopRun()
}


