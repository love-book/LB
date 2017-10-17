package controllers

import (
	"github.com/astaxie/beego"
	comm "common/conndatabase"
	"github.com/garyburd/redigo/redis"
	"models"
	"time"
	"fmt"
	"common"
	"encoding/json"
	"common/wechat"
	"strconv"
	"os"
	"strings"
	"os/exec"
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

// @Title  用户登录获取token
// @Summary  用户登录获取token
// @Description 用户登录获取token
// @Param	body	body  models.LoginForm	true  {<br/>"telphone":"手机号",<br/>"password":"密码"<br/>}
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
		token,err := common.SetToken(u.Userid+";"+u.Openid+";"+u.Province+";"+u.Imgurl+";"+u.Nickname)
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

// @Title          openid获取token
// @Summary        openid获取token
// @Description    openid获取token
// @Param	body   body  models.AccesstokenForm  true  {<br/>"openid":"openid"<br/>}
// @Success 200   {<br/> token <br/>}
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
		token,err := common.SetToken(u.Userid+";"+u.Openid+";"+u.Province+";"+u.Imgurl+";"+u.Nickname)
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
// @Title         获取微信JSDK签名信息
// @Summary       获取微信JSDK签名信息
// @Description   获取微信JSDK签名信息
// @Param	body  body  models.WxConfigForm  true  {<br/>"openid":"openid"<br/>}
// @Success 200   {<br/>"appid":"微信公众平台的AppID"<br/>,"secret":"微信公众平台的AppSecret"<br/>}
// @Failure 403 :openid is empty
// @router /getwxconfig [post]
func (this *AppController) Getwxconfig()  {
	var ob  *models.WxConfigForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if  ob.Url ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	token  :=  beego.AppConfig.String("token") // 微信公众平台的Token
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	mp := wechat.New(token, appid, secret)
	jsapi_ticket,_:=mp.TicketToken.Fresh()
	t:= strconv.FormatInt(time.Now().Unix(), 10)
	conf:= map[string]string{
		"noncestr":"ARTWm3WZYTPz0wzccnW",
		"timestamp":t,
		"jsapi_ticket":jsapi_ticket,
		"url":ob.Url,
	}
	wxConf := mp.GetSignature(conf)
	this.Rsp(true,"成功!",&wxConf)
}


// @Title         根据微信授权code获取用户token
// @Summary       根据微信授权code获取用户token
// @Description   根据微信授权code获取用户token
// @Param	code  formData  string  true  code
// @Success 200   {<br/>"data":"根据微信授权code获取用户token"<br/>}
// @Failure 403 :openid is empty
// @router /getusertokenbycode [post]
func (this *AppController) Getusertokenbycode()  {
	/*var ob  *models.WxcodeForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if  ob.Code ==""{
		this.Rsp(false,"提交参数错误!","")
	}*/
	code:=this.GetString("code")
	if  code ==""{
		this.Rsp(false,"提交参数错误!","")
	}
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	mp := wechat.NewUseracesstoken(appid, secret,code)
	//获取用户openid
	fetchRespose,err:= mp.Fetch()
	if err !=nil{
		this.Rsp(false,"失败!",err.Error())
	}
	//判断数据库是否有该用户信息,没有则需要注册一个用户
	openid := []string{fetchRespose.Openid}
	u,err :=models.GetUsersByOpenId(openid)
	if err != nil {
		 //获取微信用户信息
		 if fetchRespose.Scope == "snsapi_userinfo"{
			 mp.Openid =fetchRespose.Openid
			 mp.AccessToken = fetchRespose.AccessToken
			 info,err:=mp.GetUserinfo()
            if err!=nil{
				this.Rsp(false,"失败!",err.Error())
			}
			t:= time.Now().Unix()
			id    :=  models.GetID()
			uid   :=  fmt.Sprintf("%d", id)
			u = &models.Users{
				Userid  :uid,
				Openid  :info.Openid,
				Wnickname:info.Nickname,
				Wimgurl :info.Headimgurl,
				Nickname:info.Nickname,
				Imgurl  :info.Headimgurl,
				Gender  :info.Sex,
				Age   	:1,
				Signature :"暂未设置",
				Constellation :"暂未设置",
				Province  	:info.Province,
				City  	   :info.City,
				Logintime	:t,
				Created_at :t,
				Updated_at  :t,
			}
			if _,err:=models.AddUsers(u);err!=nil{
				 this.Rsp(false,"失败!",err.Error())
			}
		 }
	}
	token,err := common.SetToken(u.Userid+";"+u.Openid+";"+u.Province+";"+u.Imgurl+";"+u.Nickname)
	this.Rsp(true,"成功!",&token)
}




// @Title    手机号获取验证码
// @Summary  手机号获取验证码
// @Description 手机号获取验证码
// @Param	body	body  models.SmsForm	true  {<br/>"telphone":"手机号"<br/>}
// @Success 200     {}
// @Failure 403     :openid is empty
// @router /phonecode [post]
func (this *AppController) Phonecode(){
	var ob  *models.SmsForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	telphone := ob.Telphone
	if telphone =="" {
		this.Rsp(false,"提交参数错误!","")
	}
	code := GetPhoneCode(6)
	pwd,_:= os.Getwd()
	fmt.Println(pwd)
	//先进入 src\common\aliyun\core 目录执行 python setup.py install
	//然后进入src\common\aliyun\dysmsapi 目录执行 python sms.py 18515422930 678543 测试
	py:= pwd+"/../common/aliyun/dysmsapi/sms.py"
	v:="python "+py+" "+telphone+" "+code
	fmt.Println(v)
	args := strings.Split(v, " ")
	cmd := exec.Command(args[0], args[1:]...)
	buf, err := cmd.Output()
	if err == nil{
		Msg := struct {
			Message   string
			RequestId string
			Code      string
		}{}
		err:=json.Unmarshal(buf,&Msg)
		fmt.Println(Msg)
		if err!=nil{
			this.Rsp(false,"失败,请稍后重试!","")
		}
        if Msg.Message!="OK"{
			this.Rsp(false,Msg.Message,"")
		}
		rc := comm.Pool.Get()
		defer rc.Close()
		//缓存验证码十分钟
		rc.Do("SETEX",telphone,60*10,code)
		this.Rsp(true,"成功!","")
	}
	fmt.Printf("%s\n%s",buf,err)
	this.Rsp(false,"失败!","")
}



// @Title    手机号验证验证码
// @Summary  手机号验证验证码
// @Description 手机号验证验证码
// @Param	body	body  models.SmsCheckForm  true  {<br/>"telphone":"手机号"<br/>"code":"验证码"<br/>}
// @Success 200     {}
// @Failure 403     :openid is empty
// @router /checkphonecode [post]
func (this *AppController) Checkponecode(){
	var ob  *models.SmsCheckForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	telphone := ob.Telphone
	phonecode := ob.Code
	if telphone =="" || phonecode==""{
		this.Rsp(false,"提交参数错误!","")
	}
	rc := comm.Pool.Get()
	defer rc.Close()
	code,err:=redis.String(rc.Do("GET",telphone))
	fmt.Println(err)
	if code == phonecode{
		this.Rsp(true,"成功!",code)
	}
	this.Rsp(false,"失败!","")
}