// @APIVersion 1.0.0
// @Title  API文档
// @Description 文档描述
// @Contact 773683464@qq.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"api/controllers"
	"github.com/astaxie/beego/plugins/cors"
	"common"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type","token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type","token"},
		AllowCredentials: true,
	}))
	//json形式的接口
	//文档启动命令  bee run -gendoc=true -downdoc=true
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSCond(FilterToken),
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/book",
			beego.NSCond(FilterToken),
			beego.NSInclude(
				&controllers.BooksController{},
			),
		),
		beego.NSNamespace("/bookrack",
			beego.NSCond(FilterToken),
			beego.NSInclude(
				&controllers.BooksrackController{},
			),
		),
		beego.NSNamespace("/booknews",
			beego.NSCond(FilterToken),
			beego.NSInclude(
				&controllers.BooknewsController{},
			),
		),
		beego.NSNamespace("/bookorder",
			beego.NSCond(FilterToken),
			beego.NSInclude(
				&controllers.BookorderController{},
			),
		),
		//不做验证
		beego.NSNamespace("/app",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/api/index", &controllers.ApiController{}, "*:Index")
	beego.Router("/api/test", &controllers.ApiController{}, "*:Test")
	beego.Router("/api/botmsg", &controllers.ApiController{}, "*:BotMsg")
	beego.Router("/api/bindwxuser", &controllers.ApiController{}, "*:BindWxUser")
    beego.Router("/user/stoprunning/?:type", &controllers.ApiController{}, "*:StopRunning")
}

var FilterToken = func(ctx *context.Context) bool {
	auth := ctx.Request.Header.Get("token")
	if 0 == len(auth){
		return false
	}
	isaccess := common.VerifyToken(auth)
	if true != isaccess{
		return false
	}
	return true
}


