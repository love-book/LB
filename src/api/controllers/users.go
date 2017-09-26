package controllers

import (
	comm "common/conndatabase"
	"models"
	"time"
	"encoding/json"
	"strings"
	"math"
	"strconv"
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
// @Param   token   header     string  true     "token"
// @Param	body	body  models.UserinfoForm	true "{ <br/>"userid":"用户编号", <br/>"nickname":"昵称",<br/>"telphone":"手机号"}"
// @Failure 500 服务器错误!
// @router /userinfo [post]
func (this *UsersController) Userinfo() {
	var ob  models.UserinfoForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	var conditions string = ""
	if ob.Userid != ""{
		conditions+= " and userid ='"+ob.Userid+"'"
	}else{
		conditions+= " and userid ='"+this.Userid+"'"
	}
	if ob.Nickname != ""{
		conditions+= " and nickname ="+ob.Nickname
	}
	if ob.Telphone !="" {
		conditions+= " and telphone = "+ob.Telphone
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
// @Param	body	body   models.UseraddForm  true   "{<br/>"nickname":"昵称",<br/>"openid":"openid",<br/>"wnickname":"微信昵称",<br/>"wimgurl":"微信头像",<br/>"imgurl":"头像",<br/>"gender":"性别1:男2:女3:保密",<br/>"age":"年龄",<br/>"telphone":"手机号",<br/>"qq":"qq",<br/>"Constellation":"星座",<br/>"weino":"微博",<br/>"signature":"个性签名"<br/>}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /updateuser [post]
func (this *UsersController) Updateuser(){
	var ob *models.UseraddForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	user := models.Users{}
	user.Userid = this.Userid
	if err := comm.Read(&user);err == nil {
		if ob.Nickname!=""{
			user.Nickname = ob.Nickname
		}
		if ob.Openid!=""{
			user.Openid = ob.Openid
		}
		if ob.Wnickname!=""{
			user.Wnickname = ob.Wnickname
		}
		if ob.Wimgurl!= ""{
			user.Wimgurl =ob.Wimgurl
		}
		if  ob.Password!=""{
			user.Password = models.Md5([]byte(ob.Password))
		}
		if ob.Imgurl!= "" {
			user.Imgurl = ob.Imgurl
		}
		if ob.Gender!= 0{
			user.Gender = ob.Gender
		}
		if ob.Age!= 0{
			user.Age = ob.Age
		}
		if ob.Telphone!= ""{
			user.Telphone = ob.Telphone
		}
		if ob.Qq != ""{
			user.Qq = ob.Qq
		}
		if ob.Weino!= ""{
			user.Weino = ob.Weino
		}
		if ob.Signature!= ""{
			user.Signature = ob.Signature
		}
		if ob.Address!= ""{
			user.Address = ob.Address
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
// @Param   token   header     string  true  "token"
// @Param	body	body  models.GetUsersByLocaltionForm	true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页",<br/> "gender":"性别1:男2:女0:保密",<br/> "age":"年龄范围20-30",<br/>"radius":"方圆多少米范围内200-300"<br/>}"
// @Success 200  {<br/> "lat": "经度",<br/>"lang": "纬度",<br/> "openid": "openid"}
// @Failure 403 :openid is empty
// @router /getusersbylocaltion [post]
func (this *UsersController) GetUsersByLocaltion() {
	var ob  *models.GetUsersByLocaltionForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw   := ob.Draw
	Openid := this.Openid
	if Openid == "" {
		this.Rsp(false,"提交参数错误!","")
	}
	var conditions string = " "
	if ob.Gender!=""{
		conditions+= " and gender ="+ob.Gender
	}
	if ob.Age !="" {
		ageRange:=strings.Split(ob.Age,"-")
		conditions+= " and u.age >="+ageRange[0]
		conditions+= " and u.age <="+ageRange[1]
	}
	var  openstr  string
	geokey := this.Province
	if ob.Radius != "" {
		radiusRange:=strings.Split(ob.Radius,"-")
		radius,_:=strconv.ParseInt(radiusRange[1], 10, 64)
		re,err := models.GetUsersByLocaltion(this.Openid,geokey,radius,draw*length)
		if err ==nil{
			for k,v := range re{
				if k >= ((draw-1)*length){
					openstr+= "'"+v["member"]+"',"
				}
			}
			openstr=strings.Trim(openstr,",")
		}
		if openstr != ""{
			conditions+= " and openid in("+openstr+")"
		}
	}
	users,count := models.GetUserList((draw-1)*length,length,conditions)
	if len(users)<1 {
		users = []*models.UsersList{}
	}else{
		for _,kv:= range users{
			radius,err:=models.GetUsersRadiusByMembers(this.Province,this.Openid,kv.Openid)
			if err == nil{
				rad:=strings.Split(radius,".")
				kv.Radius = rad[0]+" m"
			}else{
				kv.Radius = "9 m"
			}
		}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&users}
	this.Rsp(true, "获取成功!",&json)
}


// 用户上传地理位置
// @Title  用户上传地理位置
// @Summary   用户上传地理位置
// @Description 用户上传地理位置
// @Param   token       header     string  true  "token"
// @Param	body	body  models.AddLocaltionByIDForm	true  "{ <br/>"lat":"经度", <br/>"lang":"纬度"<br/>}"
// @Success 200  {<br/> "lat": "经度",<br/>"lang": "纬度",<br/> "openid": "openid"}
// @Failure 403 :openid is empty
// @router /addlocaltionbyid [post]
func (this *UsersController) AddLocaltionByID() {
	var ob *models.AddLocaltionByIDForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	Openid := this.Openid
	Lang := ob.Lang
	Lat := ob.Lat
	if Openid == "" || Lang== 0 || Lat== 0{
		this.Rsp(false,"提交参数错误!","")
	}
	err:= models.AddLocationByID(Openid,Lang,Lat)
	if err == nil{
		this.Rsp(true,"成功!",&Openid)
	}
	this.Rsp(false,"失败!","")
}
