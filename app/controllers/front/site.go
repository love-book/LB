package front

import (
	"fmt"
	"strings"
	"app/models"
)

type SiteController struct {
	frontBaseController
}


//首页
func (this *SiteController) Index() {
	this.layout = "layouts/main-top.html"
	this.render()
}


//登陆
func (this *SiteController) Signup(){
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(this.GetString("username"))
		nickname := strings.TrimSpace(this.GetString("nickname"))
		errmsg := make(map[int]string)
		if username=="" || nickname=="" {
			errmsg[0] = "用户名或密码不能为空!"
		}else{

		}
		user := models.User{}

		if len(errmsg)>0 {
			 errmsg[0] = "用户名或密码错误!"
		}
		fmt.Println(user)
	}
	this.Error("用户名或密码错误!","/site/error")
	/*user := models.User{}
	user.Username="kasoly";
	user.Nickname="kasoly";
	user.Auth_key = "kasoly";
	user.Password_hash  = "123456";
	user.Email    = "773683464@qq.com";
	user.Phone    = 185154229;
	user.Status   = 10
	user.Created_at = 0;
	user.Updated_at = 0;
    fmt.Println(user)
	user.Insert()
	//this.showmsg("抱歉，只有超级管理员才能进行该操作！")
	*/
	this.renderPartial()
}


//登陆
func (this *SiteController) Login(){
	this.renderPartial()
}

//登出
func (this *SiteController) Logout() {

}