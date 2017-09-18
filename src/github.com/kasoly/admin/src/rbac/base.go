package rbac

import (
	"github.com/astaxie/beego"
	//. "github.com/kasoly/admin/src"
	//m "github.com/kasoly/admin/src/models"
)

type BaseController struct {
	beego.Controller
	Templatetype string //ui template type
}

func (this *BaseController) Rsp(status bool, msg string,data interface{}) {
	this.Data["json"] = &map[string]interface{}{"status": status, "msg": msg,"data":data}
	this.ServeJSON()
}

func (this *BaseController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "easyui"
	}
	return templatetype
}