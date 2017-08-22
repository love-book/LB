package controllers

import (
	comm "common/conndatabase"
	"models"
	"time"
	"common"
	"fmt"
	"encoding/json"
	"strings"
	"errors"
)

// User API
type UsersController struct {
	ApiController
}


// @Title 获取用户信息
// @Summary  获取用户信息
// @Description 获取用户信息
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   userid    formData   string  false    "用户编号"
// @Param   nickname  formData   string  false    "昵称"
// @Param   telphone  formData   string  false    "手机号"
// @Failure 500 服务器错误!
// @router /userinfo [post]
func (this *UsersController) Userinfo() {
	var user  models.Users
	var conditions string = ""
	Uid := this.GetString("userid")
	if Uid != ""{
		conditions+= " and userid ='"+Uid+"'"
	}
	Nickname := this.GetString("nickname")
	if Nickname != ""{
		conditions+= " and nickname ="+Nickname
	}
	Telphone := this.GetString("telphone")
	if Telphone !="" {
		conditions+= " and telphone = "+Telphone
	}
	if len(conditions) == 0{
		common.ErrSystem.Message = "不存在条件"
		this.renderJson(common.ErrSystem)
	}
	sql := "select * from lb_users where true "+conditions
	res :=comm.RawSeter(sql)
	err := res.QueryRow(&user)
	if err == nil {
		common.Actionsuccess.MoreInfo = &user
		this.renderJson(common.Actionsuccess)
	}
	common.ErrSystem.Message = "不存在当前用户"
	this.renderJson(common.ErrSystem)
}

// @Title 添加用户
// @Summary  添加用户
// @Description 添加用户
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   nickname  formData   string  true      "昵称"
// @Param   openid    formData   string  false     "openid"
// @Param   wnickname  formData   string  false    "微信昵称"
// @Param   wimgurl    formData   string  false     "微信头像"
// @Param   mgurl       formData   string  false    "头像"
// @Param   gender      formData   int  false    "性别1:男2:女3:保密"
// @Param   age         formData   int  false    "年龄"
// @Param   telphone    formData   int  false    "手机号"
// @Param   qq          formData   string  false    "QQ"
// @Param   weino       formData   string  false    "微博"
// @Param   signature   formData   string  false    "个性签名"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /useradd [post]
func (this *UsersController) Useradd() {
	user := models.Users{}
	id := models.GetID()
	user.Userid   =  fmt.Sprintf("%d", id)
	user.Nickname = this.GetString("nickname")
	user.Openid  = this.GetString("openid")
	user.Wnickname = this.GetString("wnickname")
	user.Wimgurl = this.GetString("wimgurl")
	user.Imgurl = this.GetString("imgurl")
	var Gender,_ = this.GetInt8("gender")
	if Gender != 0{
		user.Gender = Gender
	}
	var Age ,_ = this.GetInt32("age")
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
	user.Updated_at = 0
	err := user.InsertValidation()
	if err != nil {
		common.ErrSystem.Message = fmt.Sprint(err)
		this.renderJson(common.ErrSystem)
	}
	err = comm.Insert(&user)
	if err != nil {
		common.ErrSystem.Message = fmt.Sprint(err)
		this.renderJson(common.ErrSystem)
	}
	common.Actionsuccess.MoreInfo = &user
	this.renderJson(common.Actionsuccess)
}


// @Title 修改用户
// @Description 修改用户
// @Summary  修改用户
// @Success 200  {<br/> "userid": "用户编号",<br/> "openid": "openid",<br/> "wnickname": "微信昵称",<br/> "wimgurl": "微信头像", <br/>"nickname": "昵称",<br/> "imgurl": "头像",<br/> "gender":"性别1:男2:女3:保密",<br/> "age":"年龄",<br/> "telphone":"手机号", <br/>"QQ": "QQ",<br/> "weino": "微博号", <br/>"signature": "个性签名", <br/>"address": "个人地址",<br/> "created_at": "注册时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   userid      formData   int  true  "用户编号"
// @Param   nickname    formData   string  false    "昵称"
// @Param   openid      formData   string  false    "openid"
// @Param   wnickname   formData   string  false    "微信昵称"
// @Param   wimgurl     formData   string  false    "微信头像"
// @Param   mgurl       formData   string  false   "头像"
// @Param   gender      formData   int  false     "性别1:男2:女3:保密"
// @Param   age         formData   int  false     "年龄"
// @Param   telphone    formData   string  false     "手机号"
// @Param   qq          formData   string  false  "QQ"
// @Param   weino       formData   string  false  "微博"
// @Param   signature   formData   string  false  "个性签名"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /userupdate [post]
func (this *UsersController) Userupdate(){
	user := models.Users{}
	user.Userid = this.GetString("userid")
	err := user.UserValidation()
	if err != nil {
		common.ErrSystem.Message =  fmt.Sprint(err)
		this.renderJson(common.ErrSystem)
	}
	if err := comm.Read(&user);err == nil {
		Nickname := this.GetString("nickname")
		if Nickname != ""{
			user.Nickname = Nickname
		}
		Openid := this.GetString("openid")
		if Openid != ""{
			user.Openid = Openid
		}
		Wnickname := this.GetString("wnickname")
		if Wnickname != ""{
			user.Wnickname = Wnickname
		}
		Wimgurl := this.GetString("wimgurl")
		if Wimgurl != ""{
			user.Wimgurl = Wimgurl
		}
		Imgurl := this.GetString("imgurl")
		if Imgurl != ""{
			user.Imgurl = Imgurl
		}
		var Gender,_ = this.GetInt8("gender")
		if Gender != 0{
			user.Gender = Gender
		}
		var Age ,_ = this.GetInt32("age")
		if Age != 0{
			user.Age = Age
		}
		var telphone  = this.GetString("telphone")
		if telphone != ""{
			user.Telphone = telphone
		}
		Qq := this.GetString("qq")
		if Qq != ""{
			user.Qq = Qq
		}
		Weino := this.GetString("weino")
		if Weino != ""{
			user.Weino = Weino
		}
		Signature := this.GetString("signature")
		if Signature != ""{
			user.Signature = Signature
		}
		Address := this.GetString("address")
		if Address != ""{
			user.Address = Address
		}
		user.Updated_at = time.Now().Unix()
		if update:= comm.Update(&user);update ==nil{
			common.Actionsuccess.MoreInfo = &user
			this.renderJson(common.Actionsuccess)
		}
	}else{
		common.ErrSystem.Message = "没有当前记录"
	}
	this.renderJson(common.ErrSystem)
}


// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *UsersController) Post() {
	var v models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddUsers(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	v, err := models.GetUsersById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UsersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	v := models.Users{Userid: idStr}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUsersById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	if err := models.DeleteUsers(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
