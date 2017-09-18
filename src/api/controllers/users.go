package controllers

import (
	comm "common/conndatabase"
	"models"
	"time"
)

// User API
type UsersController struct {
	ApiController
}

// @Title 上传文件
// @Description 上传文件
// @Summary  上传文件
// @Success 200  {<br/>  "url": "url" <br/> }
// @Param   fname  formData   file  true   "文件"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @Consumes  multipart/form-data
// @router /uploadfile [post]
func (this *UsersController) Uploadfile() {
	f, h, err := this.GetFile("fname")
	if err != nil {
		this.Rsp(false, err.Error(),"")
	}
	defer f.Close()
	FilesUrl := comm.FileUrl()
	url := FilesUrl + h.Filename
	file:=this.SaveToFile("fname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	if  file != nil {
		this.Rsp(false, file.Error(),"")
	}
	this.Rsp(true, "上传成功!", url)
}
// @Title 获取用户信息
// @Summary  获取用户信息
// @Description 获取用户信息
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token     header     string  true     "token"
// @Param   userid    formData   string  false    "用户编号"
// @Param   nickname  formData   string  false    "昵称"
// @Param   telphone  formData   string  false    "手机号"
// @Failure 500 服务器错误!
// @router /userinfo [post]
func (this *UsersController) Userinfo() {
	var conditions string = ""
	if v := this.GetString("userid");v != ""{
		conditions+= " and userid ='"+v+"'"
	}else{
		conditions+= " and userid ='"+this.Userid+"'"
	}
	if v := this.GetString("nickname");v != ""{
		conditions+= " and nickname ="+v
	}
	if v := this.GetString("telphone");v !="" {
		conditions+= " and telphone = "+v
	}
	if v:= this.GetString("openid");v !="" {
		conditions+= " and openid = "+v
	}
	if len(conditions) == 0{
		this.Rsp(false,"不存在条件","")
	}
	var user  models.Users
	sql := "select * from lb_users where true "+conditions
	res := comm.RawSeter(sql)
	err := res.QueryRow(&user)
	if err == nil {
		this.Rsp(true,"成功!",&user)
	}
	this.Rsp(false,"不存在当前用户","")
}
// @Title 修改用户
// @Description 修改用户
// @Summary  修改用户
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param   userid      formData   string  true  "用户编号"
// @Param   nickname    formData   string  false  "昵称"
// @Param   openid      formData   string  false  "openid"
// @Param   wnickname   formData   string  false  "微信昵称"
// @Param   wimgurl     formData   string  false  "微信头像"
// @Param   mgurl       formData   string  false  "头像"
// @Param   gender      formData   int  false     "性别1:男2:女3:保密"
// @Param   age         formData   int  false     "年龄"
// @Param   telphone    formData   string  false  "手机号"
// @Param   password    formData   string  false  "密码"
// @Param   qq          formData   string  false  "QQ"
// @Param   weino       formData   string  false  "微博"
// @Param   signature   formData   string  false  "个性签名"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /updateuser [post]
func (this *UsersController) Updateuser(){
	user := models.Users{}
	user.Userid = this.GetString("userid")
	err := user.UserValidation()
	if err != nil {
		this.Rsp(false,err.Error(),"")
	}
	if err := comm.Read(&user);err == nil {
		if v:=this.GetString("nickname");v!=""{
			user.Nickname = v
		}
		if v:=this.GetString("openid");v!=""{
			user.Openid = v
		}
		if v:=this.GetString("wnickname");v!=""{
			user.Wnickname = v
		}
		if v:=this.GetString("wimgurl");v!=""{
			user.Wimgurl = v
		}
		if v:=this.GetString("password");v!=""{
			user.Password = models.Md5([]byte(v))
		}
		if v := this.GetString("imgurl");v!= ""{
			user.Imgurl = v
		}
		if v,_ := this.GetInt64("gender");v!= 0{
			user.Gender = v
		}
		if v ,_ := this.GetInt32("age");v!= 0{
			user.Age = v
		}
		if v := this.GetString("telphone");v!= ""{
			user.Telphone = v
		}
		if v := this.GetString("qq");v != ""{
			user.Qq = v
		}
		if v := this.GetString("weino");v!= ""{
			user.Weino = v
		}
		if v := this.GetString("signature");v!= ""{
			user.Signature = v
		}
		if v := this.GetString("address");v!= ""{
			user.Address = v
		}
		user.Updated_at = time.Now().Unix()
		if update:= comm.Update(&user);update ==nil{
			this.Rsp(true,"成功",&user)
		}else{
			this.Rsp(false,"失败","")
		}
	}
	this.Rsp(false,"没有当前记录","")
}


// 根据openid获取地理位置
// Title  根据openid获取地理位置
// Summary   根据openid获取地理位置
// Description 根据openid获取地理位置
// @Param   token       header     string  true  "token"
// Success 200  {<br/> "lat": "经度",<br/>"lang": "纬度",<br/> "openid": "openid"}
// Failure 403 :openid is empty
// router /getlocaltionbyid [post]
func (this *UsersController) GetLocaltionByID() {
	Openid := this.Openid
	if Openid == "" {
		this.Rsp(false,"openid不能为空!","")
	}
	l,_ := models.GetLocationByID(Openid)
	l["openid"] = Openid
	this.Rsp(true,"成功",& l)
}
// @Title  获取附近的人
// @Summary   获取附近的人
// @Description 获取附近的人
// @Param   token       header     string  true  "token"
// @Param	radius	formData  string	true  "方圆多少米范围内"
// @Success 200  {<br/> "lat": "经度",<br/>"lang": "纬度",<br/> "openid": "openid"}
// @Failure 403 :openid is empty
// @router /getusersbylocaltion [post]
func (this *UsersController) GetUsersByLocaltion() {
	Openid := this.Openid
	Radius,_:= this.GetInt64("radius",1000)
	if Openid == "" {
		this.Rsp(false,"提交参数错误!","")
	}
	var re  []map[string]string
	openid := []string{Openid}
	u,err:=models.GetUsersByOpenId(openid)
	if err == nil{
		geokey :=u.Province+"-"+u.City
		re,_= models.GetUsersByLocaltion(Openid,geokey,Radius)
	}
	this.Rsp(true,"成功!",&re)
}


// 用户上传地理位置
// @Title  用户上传地理位置
// @Summary   用户上传地理位置
// @Description 用户上传地理位置
// @Param   token       header     string  true  "token"
// @Param	lang	formData  string	true  "纬度"
// @Param	lat  	formData  string	true  "经度"
// @Success 200  {<br/> "lat": "经度",<br/>"lang": "纬度",<br/> "openid": "openid"}
// @Failure 403 :openid is empty
// @router /addlocaltionbyid [post]
func (this *UsersController) AddLocaltionByID() {
	Openid := this.Openid
	Lang,_ := this.GetFloat("lang")
	Lat,_ := this.GetFloat("lat")
	if Openid == "" || Lang== 0 || Lat== 0{
		this.Rsp(false,"提交参数错误!","")
	}
	err:= models.AddLocationByID(Openid,Lang,Lat)
	if err == nil{
		this.Rsp(true,"成功!",&Openid)
	}
	this.Rsp(false,"失败!","")
}
