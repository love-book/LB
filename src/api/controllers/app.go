package controllers

import (
	"github.com/astaxie/beego"
	comm "common/conndatabase"
	"models"
	"time"
	"fmt"
	"common"
	"encoding/json"
	//mq "common/rabbitmq"
)

// App API
type AppController struct {
	ApiController
}

// @Title 用户注册
// @Summary   用户注册
// @Description  用户注册
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/>"constellation":"星座",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param	body	body   models.UseraddForm  true   "{<br/>"nickname":"昵称",<br/>"openid":"openid",<br/>"wnickname":"微信昵称",<br/>"wimgurl":"微信头像",<br/>"imgurl":"头像",<br/>"gender":"性别1:男2:女3:保密",<br/>"age":"年龄",<br/>"telphone":"手机号",<br/>"qq":"qq",<br/>"weino":"微博",<br/>"signature":"个性签名"<br/>}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /useradd [post]
func (this *AppController) Useradd() {
	var ob *models.UseraddForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	user := models.Users{}
	id := models.GetID()
	user.Userid   =  fmt.Sprintf("%d", id)
	user.Nickname = ob.Nickname
	user.Openid   = ob.Openid
	user.Wnickname= ob.Wnickname
	user.Wimgurl  = ob.Wimgurl
	user.Imgurl =  ob.Imgurl
	user.Gender =  ob.Gender
	user.Age    =  ob.Age
	user.Telphone   =  ob.Telphone
	user.Qq         =  ob.Qq
	user.Weino      =  ob.Weino
	user.Signature  =  ob.Signature
	user.Address    =  ob.Address
	user.Created_at =  time.Now().Unix()
	user.Updated_at =  time.Now().Unix()
	err := user.InsertValidation()
	if err != nil {
		this.Rsp(false,err.Error(),"")
	}
	err = comm.Insert(&user)
	if err != nil {
		this.Rsp(false,"失败!","")
	}
	this.Rsp(true,"成功!",&user)
}
// 用户登录
// @Title  用户登录获取token
// @Summary  用户登录获取token
// @Description 用户登录获取token
// @Param	body	body  models.LoginForm	true  "{<br/>"telphone":"手机号",<br/>"password":"密码"<br/>}
// @Success 200  {<br/>oX8vKwqsI52Me2J2kCTkkYNohHfQ<br/>oX8vKwueTHOC3wrUkm2eJBnm-m6A<br/>}
// @Failure 403 :openid is empty
// @router /login [post]
func (this *AppController) Login()  {
	var ob  *models.LoginForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	telphone := ob.Telphone
	password := ob.Password
	if telphone =="" || password ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	pass:= models.Md5([]byte(password))
	u,err:=models.GetUsersBypass([]string{telphone,pass})
	if err==nil{
		token,err := common.SetToken(u.Userid+";"+u.Openid+";"+u.Province+";"+u.City)
		if 0 == len(token) {
			fmt.Println(token,err.Error())
			this.Rsp(false,"登录失败!","")
		}else{
			this.Rsp(true,"登录成功!",&token)
		}
	}else{
		this.Rsp(false,"不存在当前用户!","")
	}
}
// openid获取token
// @Title    openid获取token
// @Summary  openid获取token
// @Description openid获取token
// @Param	body	body  models.AccesstokenForm  true  "{<br/>"openid":"openid"<br/>}
// @Success 200  {<br/> token}
// @Failure 403 :openid is empty
// @router /accesstoken [post]
func (this *AppController) Accesstoken()  {
	var ob  *models.AccesstokenForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if  ob.Openid ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	u,err:=models.GetUsersByOpenId([]string{ob.Openid})
	if err==nil{
		token,err := common.SetToken(u.Userid+";"+u.Openid+";"+u.Province+";"+u.City)
		if 0 == len(token) {
			fmt.Println(token,err.Error())
			this.Rsp(false,"登录失败!","")
		}else{
			this.Rsp(true,"登录成功!",&token)
		}
	}else{
		this.Rsp(false,"不存在当前用户!","")
	}
}
// openid       获取微信公众号信息
// @Title       获取微信公众号信息
// @Summary     获取微信公众号信息
// @Description 获取微信公众号信息
// @Success 200 "{<br/>"appid":"微信公众平台的AppID"<br/>,"secret":"微信公众平台的AppSecret"<br/>}
// @Failure 403 :openid is empty
// @router /getwxaccount [post]
func (this *AppController) Getwxaccount()  {
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	account := &map[string]string{"appid":appid,"secret":secret}
	this.Rsp(true,"成功!",&account)
}