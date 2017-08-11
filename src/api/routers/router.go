// @APIVersion 1.0.0
// @Title  API文档
// @Description 文档描述
// @Contact 773683464@qq.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"api/controllers"
	"common"
)

func init() {
	//json形式的接口
	//文档启动命令  bee run -gendoc=true -downdoc=true
	ns := beego.NewNamespace("/v1",
		//beego.NSCond(FilterToken),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/book",
			beego.NSInclude(
				&controllers.BooksController{},
			),
		),
		beego.NSNamespace("/bookrack",
			beego.NSInclude(
				&controllers.BooksrackController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/api/index", &controllers.ApiController{}, "*:Index")
	beego.Router("/user/login", &controllers.ApiController{}, "*:Login")
	beego.Router("/api/test", &controllers.ApiController{}, "*:Test")
	beego.Router("/user/stoprunning/?:type", &controllers.ApiController{}, "*:StopRunning")
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
