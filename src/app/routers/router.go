// @APIVersion 1.0.0
// @Title  API文档
// @Description 文档描述
// @Contact 773683464@qq.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"app/controllers/admin"
	"app/controllers/front"
	"common"
)

func init() {
	//json形式的接口
	//文档启动命令  bee run -gendoc=true -downdoc=true
	ns := beego.NewNamespace("/v1",
		//beego.NSCond(FilterToken),
		   beego.NSNamespace("/user",
		    	beego.NSInclude(
			    	&admin.UserController{},
		    	),
	     	),
		   beego.NSNamespace("/book",
			    beego.NSInclude(
				   &admin.BooksController{},
			   ),
		   ),
	)
	beego.AddNamespace(ns)
	//后台路由
	beego.Router("/", &front.SiteController{}, "*:Index")
	beego.Router("/site/index", &front.SiteController{}, "*:Index")
	beego.Router("/site/signup", &front.SiteController{}, "*:Signup")
	beego.Router("/site/login", &front.SiteController{}, "*:Login")
	beego.Router("/site/logout", &front.SiteController{}, "*:Logout")
}

var FilterToken = func(ctx *context.Context) bool {
	auth := ctx.Request.Header.Get("Authorization")
	if 0 == len(auth){
		return false

	}
	isaccess := common.VerifyToken(auth)
	if true != isaccess{
		return false
	}
	return true
}
