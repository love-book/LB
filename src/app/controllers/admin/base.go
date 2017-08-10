package admin

import (
	"github.com/astaxie/beego"
	"strings"
	"html/template"
)

type baseController struct {
	beego.Controller
}

//json输出
func (this *baseController) renderJson(Json interface{}) {
	//Json := map[string]interface{}{"code":0,"msg": "成功!","data":"1111"}
	this.Data["json"] = Json
	this.ServeJSON()
	this.StopRun()
}

//渲染模版(包含布局)
func (this *baseController) render(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname =   tpl[0] + ".html"
	} else {
		controller, action := this.GetControllerAndAction()
		controller = strings.ToLower(controller[0 : len(controller)-10])
		action  =  strings.ToLower(action)
		tplname = controller + "/" + action + ".html"
	}
	this.Layout = "admin/main.html"
	this.LayoutSections["Header"]  = "admin/header.html"
	this.LayoutSections["Left"]    = "admin/left.html"
	this.LayoutSections["Content"] = "admin/content.html"
	this.LayoutSections["Footer"]  = "admin/footer.html"
	this.LayoutSections["Right"]   = "admin/right.html"
	this.Data["asset"]   = "/static/"
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["xsrfdata"]   = template.HTML(this.XSRFFormHTML())
	this.Data["asset"]   = "/static/admin"
	this.TplName = tplname
}

//渲染模版(不包含布局)
func (this *baseController) renderPartial(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname =  tpl[0] + ".html"
	} else {
		controller, action := this.GetControllerAndAction()
		controller = strings.ToLower(controller[0 : len(controller)-10])
		action  =  strings.ToLower(action)
		tplname = controller + "/" + action + ".html"
	}
	this.TplName = tplname
}

//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}