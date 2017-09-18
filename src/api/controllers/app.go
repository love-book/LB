package controllers

import (
	comm "common/conndatabase"
	"models"
	"time"
	"fmt"
	"common"
)

// App API
type AppController struct {
	ApiController
}

// @Title 添加用户
// @Summary  添加用户
// @Description 添加用户
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   nickname    formData   string  true     "昵称"
// @Param   openid      formData   string  false    "openid"
// @Param   wnickname   formData   string  false   	"微信昵称"
// @Param   wimgurl     formData   string  false    "微信头像"
// @Param   mgurl       formData   string  false    "头像"
// @Param   gender      formData   int  false    	"性别1:男2:女3:保密"
// @Param   age         formData   int  false    	"年龄"
// @Param   telphone    formData   int  false    	"手机号"
// @Param   qq          formData   string  false    "QQ"
// @Param   weino       formData   string  false    "微博"
// @Param   signature   formData   string  false    "个性签名"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /useradd [post]
func (this *AppController) Useradd() {
	user := models.Users{}
	id := models.GetID()
	user.Userid   =  fmt.Sprintf("%d", id)
	user.Nickname = this.GetString("nickname")
	user.Openid  = this.GetString("openid")
	user.Wnickname = this.GetString("wnickname")
	user.Wimgurl = this.GetString("wimgurl")
	user.Imgurl = this.GetString("imgurl")
	var Gender,_ = this.GetInt64("gender")
	if Gender != 0{
		user.Gender = Gender
	}
	var Age ,_ = this.GetInt32("age",1)
	if Age != 0{
		user.Age = Age
	}
	var Telphone = this.GetString("telphone")
	if Telphone != ""{
		user.Telphone = Telphone
	}
	user.Qq  = this.GetString("qq")
	user.Weino = this.GetString("weino")
	user.Signature = this.GetString("signature")
	user.Address = this.GetString("address")
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
// @Param	telphone	formData  string	true  "手机号"
// @Param	password	formData  string	true  "密码"
// @Success 200  {<br/> "telphone": "手机号",<br/>"password": "密码",<br/>}
// @Failure 403 :openid is empty
// @router /login [post]
func (this *AppController) Login()  {
	telphone := this.GetString("telphone")
	password := this.GetString("password")
	if telphone =="" || password ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	pass:= models.Md5([]byte(password))
	u,err:=models.GetUsersBypass([]string{telphone,pass})
	if err==nil{
		token,err := common.SetToken(u.Userid+"-"+u.Openid)
		if 0 == len(token) {
			fmt.Println(token,err.Error())
			this.Rsp(false,"登录失败!","")
		}else{
			this.Rsp(true,"登录成功!",token)
		}
	}else{
		this.Rsp(false,"不存在当前用户!","")
	}
}
// openid获取token
// @Title    openid获取token
// @Summary  openid获取token
// @Description openid获取token
// @Param	openid	formData  string	true  "openid"
// @Success 200  {<br/> token}
// @Failure 403 :openid is empty
// @router /accesstoken [post]
func (this *AppController) Accesstoken()  {
	openid := this.GetString("openid")
	if  openid ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	u,err:=models.GetUsersByOpenId([]string{openid})
	if err==nil{
		token,err := common.SetToken(u.Userid+"-"+u.Openid)
		if 0 == len(token) {
			fmt.Println(token,err.Error())
			this.Rsp(false,"登录失败!","")
		}else{
			this.Rsp(true,"登录成功!",token)
		}
	}else{
		this.Rsp(false,"不存在当前用户!","")
	}
}