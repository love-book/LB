package front

import (
	"strings"
	"github.com/astaxie/beego"
	"html/template"
)

type frontBaseController struct {
	beego.Controller
	userid          int64
	username        string
	controller      string
	action          string
	layout          string
	layoutSections  map[string]string
	defaultController string
	defaultAction    string
}

func (this *frontBaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.defaultController = "site"
	this.defaultAction = "index"
	this.controller = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.action = strings.ToLower(actionName)
	this.layout = "layouts/main.html"
	this.layoutSections = make(map[string]string)
	this.layoutSections["Header"]  = "layouts/header.html"
	this.layoutSections["Left"]    = "layouts/left.html"
	this.layoutSections["Content"] = "layouts/content.html"
	this.layoutSections["Footer"]  = "layouts/footer.html"
	this.layoutSections["Right"]   = "layouts/right.html"
	//this.auth()
	//this.checkPermission()
}

//登录状态验证
func (this *frontBaseController) auth() {

}


//渲染模版(包含布局)
func (this *frontBaseController) render(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname =   tpl[0] + ".html"
	} else {
		tplname =  this.controller + "/" + this.action + ".html"
	}
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	this.Data["asset"]   = "/static/"
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Layout = this.layout
	this.LayoutSections = this.layoutSections
	this.TplName = tplname
}


//渲染模版(不包含布局)
func (this *frontBaseController) renderPartial(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname =  tpl[0] + ".html"
	} else {
		tplname =  this.controller + "/" + this.action + ".html"
	}
	this.Data["xsrfdata"]   = template.HTML(this.XSRFFormHTML())
	this.Data["asset"]   = "/static/"
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.TplName = tplname
}

//显示错误提示
func (this *frontBaseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["asset"]   = "/static/"
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout  = this.layout
	this.TplName = "site/error.html"
	this.Render()
	this.StopRun()
}

//显示错误提示
func (this *frontBaseController) Error(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["asset"]   = "/static/"
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.layout
	this.TplName = "/site/error.html"
	this.Render()
	this.StopRun()
}

//是否post提交
func (this *frontBaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *frontBaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//权限验证
func (this *frontBaseController) checkPermission() {
	if this.userid != 1 && this.controller == "user" {
		this.showmsg("抱歉，只有超级管理员才能进行该操作！")
	}
}


