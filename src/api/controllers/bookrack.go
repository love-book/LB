package controllers

import (
	comm "common/conndatabase"
	"common"
	"models"
	"fmt"
	"time"
	"strconv"
	"github.com/astaxie/beego/validation"
	"strings"
	"math"
	"encoding/json"
)

type BooksrackController struct {
	ApiController
}

// @Title    书城
// @Summary  书城
// @Description  书城
// @Success 200 {<br/>"bookqid":"图书唯一编号",<br/>"userid": "用户编号",<br/>"openid": "oX8vKwueTHOC3wrUkm2eJBnm-m6A",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"wnickname": "微信昵称",<br/>"wimgurl": "微信头像",<br/>"nickname": "用户昵称",<br/>"imgurl": "用户头像",<br/>"gender": "性别1:男2:女0:保密",<br/>"age": "年龄",<br/>"telphone": "手机号",<br/>"password": "密码",<br/>"qq": "QQ号",<br/>"weibo": "微博号",<br/>"signature": "个性签名",<br/>"constellation": "星座",<br/>"province": "北京市",<br/>"city": "北京市",<br/>"address": "地址",<br/>"long":"经度",<br/>"lat":"纬度",<br/>"logintime": "用户最后登录时间",<br/>"created_at": "用户注册时间",<br/>"updated_at": "用户修改资料时间",<br/>"radius": "距离" <br/>}
// @Param   token     header   string  true  "token"
// @Param	body	  body 	models.BookracklistForm   true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页",<br/> "gender":"性别1:男2:女0:保密",<br/> "age":"年龄范围20-30",<br/>"radius":"方圆多少米范围内"<br/>}"
// @Failure 500 服务器错误!
// @router /bookracklist [post]
func (this *BooksrackController) Bookracklist() {
	var ob  *models.BookracklistForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	var conditions string = " "
	conditions+= " and r.book_state ='1'"
	if ob.Gender !="" {
		conditions+= " and u.gender ="+ob.Gender
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
			conditions+= " and u.openid in("+openstr+")"
		}
	}
	books,count := models.MyBooksrackList((draw-1)*length,length,conditions)
	if len(books) < 1 {
		books = []*models.BookrackList{}
	}else{
		for _,kv:= range books{
			radius,err:=models.GetUsersRadiusByMembers(geokey,this.Openid,kv.Openid)
			if err == nil{
				rad:=strings.Split(radius,".")
				kv.Radius = rad[0]+" m"
			}else{
				kv.Radius = "9 m"
			}
		}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&books}
	this.Rsp(true, "获取成功!",&json)
}

// @Title 添加到书架
// @Summary  添加到书架
// @Description 添加到书架
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   token   header   string  true  "token"
// @Param	body	body   models.BookrackaddForm	true "{ <br/>"bookid":"图书编号" <br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackadd [post]
func (this *BooksrackController) Bookrackadd() {
	var ob  *models.BookrackaddForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	userid := this.Userid
	bookid := ob.Bookid
	if bookid=="" || userid==""{
		this.Rsp(false, "参数错误!","")
	}
	//查询图书信息
	if _,err:= models.GetBook(bookid);err != nil {
		this.Rsp(false, "没有当前图书!","")
	}
	//加入用户书架
	book := models.Bookrack{}
	book.Userid = userid
	book.Bookid = bookid
	book.Book_state = "1"
	book.Is_borrow  = "1"
	_,err:=models.GetBookByUidAndBookId(userid,bookid)
	if err != nil {
		id := models.GetID()
		book.Bookqid = fmt.Sprintf("%d",id)
		book.Create_time = time.Now().Unix()
		book.Update_time = time.Now().Unix()
		id,err = models.AddBookrack(&book)
		if err != nil {
			this.Rsp(false, "图书加入书架失败!","")
		}
	}
	this.Rsp(true, "当前图书已加入书架!",&book)
}

// @Title 扫条码添加到书架
// @Summary  扫条码添加到书架
// @Description 扫条码添加到书架
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   token    header     string  true  "token"
// @Param	body	body  models.BookrackaddbysnForm	true "{ <br/>"isbn":"图书条码" <br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackaddbysn [post]
func (this *BooksrackController) Bookrackaddbysn() {
	var ob  models.BookrackaddbysnForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	userid  :=   this.Userid
	isbn := ob.Isbn
	if userid =="" ||isbn ==""{
		this.Rsp(false, "参数错误!","")
	}
	var Bookid string
	model,err:= models.GetIbsn(isbn)
	if err != nil {
		//state 1:扫码正常添加，3:扫码没有查询到结果待补充
		var state uint8 = 1
		res, berr:= common.GetBarcodeInfo(isbn)
		//查询失败
		if berr != nil{
			this.Rsp(false, "网络繁忙!","")
		}
		//查询不到信息
		if  res.Charge == false{
			state = 3
			res.Result.Showapi_res_body.GoodsName = "未知"
			res.Result.Showapi_res_body.ManuName = ""
			res.Result.Showapi_res_body.Img = ""
			res.Result.Showapi_res_body.Price = "0"
		}
		body :=  &res.Result.Showapi_res_body
		model := models.Books{}
		Uid := models.GetID()
		model.Bookid   = fmt.Sprintf("%d", Uid)
		Bookid = model.Bookid
		model.Bookname = body.GoodsName
		model.Author   = body.ManuName
		model.Imagehead = ""
		model.Imageback = ""
		model.Imageurl  = body.Img
		model.Describe =body.GoodsName
		price, _:= strconv.ParseUint(body.Price, 10, 16)
		model.Price  = uint16(price)
		model.Isbn   = isbn
		model.Depreciation = 0
		model.State = state
		model.Userid = userid
		valid := validation.Validation{}
		valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
		valid.Required(model.Bookname, "bookname").Message("书名不能为空！")
		valid.Required(model.Isbn,"isbn").Message("条形码不能为空！")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				this.Rsp(false, err.Error(),"")
			}
		}
		err = comm.Insert(&model)
		if err != nil{
			this.Rsp(false,"添加失败!","")
		}
	}else{
		Bookid = model.Bookid
	}
	book := models.Bookrack{}
	b,err := models.GetBookByUidAndBookId(userid,Bookid)
	if err != nil {
		//加入用户书架
		book.Userid  = userid
		book.Bookid  = Bookid
		book.Book_state = "1"
		book.Is_borrow  = "1"
		id := models.GetID()
		book.Bookqid = fmt.Sprintf("%d",id)
		t:=time.Now().Unix()
		book.Create_time = t
		book.Update_time = t
		err := comm.Insert(&book)
		if err == nil {
			this.Rsp(true,"当前图书已加入书架!",&book)
		}
	}else{
		this.Rsp(true,"当前图书已加入书架!",&b)
	}
	this.Rsp(false,"添加失败!","")
}


// @Title   删除书架图书
// @Summary   删除书架图书
// @Description  删除书架图书
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充4:删除",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   token   header     string  true  "token"
// @Param	body	body  models.BookStateForm 	true "{<br/>"bookqid":string,<br/>}"
// @Failure 100 错误提示信息!y
// @Failure 500 服务器错误!
// @router /bookrackdelete [post]
func (this *BooksrackController) Bookrackdelete() {
	var ob  *models.BookStateForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if len(ob.Bookqid)<=0 {
		this.Rsp(false,"参数错误!","")
	}
	for _,v:= range ob.Bookqid{
		book,err:= models.GetBookById(v)
		if err == nil {
			if book.Is_borrow == "2"{
				this.Rsp(false,"图书已借出,不能操作状态",v)
			}
			models.DelBookRackById(v)
		}else{
			this.Rsp(false,"当前图书不存在",v)
		}
	}
	this.Rsp(true,"操作成功!","")
}


// @Title 我的书架
// @Summary  我的书架
// @Description 我的书架
// @Success 200 {<br/>"bookqid": "图书唯一编号",<br/>"userid": "用户编号",<br/>"openid": "oX8vKwueTHOC3wrUkm2eJBnm-m6A",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"wnickname": "微信昵称",<br/>"wimgurl": "微信头像",<br/>"nickname": "用户昵称",<br/>"imgurl": "用户头像",<br/>"gender": "性别1:男2:女0:保密",<br/>"age": "年龄",<br/>"telphone": "手机号",<br/>"password": "密码",<br/>"qq": "QQ号",<br/>"weibo": "微博号",<br/>"signature": "个性签名",<br/>"constellation": "星座",<br/>"province": "北京市",<br/>"city": "北京市",<br/>"address": "地址",<br/>"long":"经度",<br/>"lat":"纬度",<br/>"logintime": "用户最后登录时间",<br/>"created_at": "用户注册时间",<br/>"updated_at": "用户修改资料时间",<br/>"radius": "距离"<br/>}
// @Param   token       header     string  true  "token"
// @Param	body	body 	 models.ConcernBookListForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页"<br/> }"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /mybookrack [post]
func (this *BooksrackController) Mybookrack() {
	var ob  *models.ConcernBookListForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	userid   :=  this.Userid
	if ob.Userid !=""{
		userid = ob.Userid
	}
	if userid =="" {
		this.Rsp(false, "参数错误!","")
	}
	var conditions string = " "
	conditions+= " and u.userid ='"+userid+"'"
	books,count := models.MyBooksrackList((draw-1)*length,length,conditions)
	if len(books) < 1 {
		books = []*models.BookrackList{}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&books}
	this.Rsp(true, "获取成功!",&json)
}



// @Title 拥有此书的人
// @Summary  拥有此书的人
// @Description 拥有此书的人
// @Success 200 {<br/>"bookqid": "图书唯一编号",<br/>"userid": "用户编号",<br/>"openid": "oX8vKwueTHOC3wrUkm2eJBnm-m6A",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"wnickname": "微信昵称",<br/>"wimgurl": "微信头像",<br/>"nickname": "用户昵称",<br/>"imgurl": "用户头像",<br/>"gender": "性别1:男2:女0:保密",<br/>"age": "年龄",<br/>"telphone": "手机号",<br/>"password": "密码",<br/>"qq": "QQ号",<br/>"weibo": "微博号",<br/>"signature": "个性签名",<br/>"constellation": "星座",<br/>"province": "北京市",<br/>"city": "北京市",<br/>"address": "地址",<br/>"long":"经度",<br/>"lat":"纬度",<br/>"logintime": "用户最后登录时间",<br/>"created_at": "用户注册时间",<br/>"updated_at": "用户修改资料时间",<br/>"radius": "距离"<br/>}
// @Param   token       header     string  true  "token"
// @Param	body	  body 	models.GetbookusersForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页",<br/> "bookid":"书编号"<br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /getbookusers [post]
func (this *BooksrackController) Getbookusers() {
	var ob  *models.GetbookusersForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	bookid := ob.Bookid
	if bookid==""{
		this.Rsp(false, "参数错误!","")
	}
	var conditions string = " "
	conditions+= " and r.bookid ='"+bookid+"'"
	books,count := models.MyBooksrackList((draw-1)*length,length,conditions)
	if len(books) < 1 {
		books = []*models.BookrackList{}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&books}
	this.Rsp(true, "获取成功!",&json)
}

// @Title    查看一本书信息
// @Summary  查看一本书信息
// @Description 查看一本书信息
// @Success 200 {<br/>"bookqid": "图书唯一编号",<br/>"userid": "用户编号",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"concernid":"收藏编号",<br/>"userid_to":"收藏人id", <br/>"userid_from":"图书编号",<br/>"concern_type":"收藏类型1:图书2:人",<br/> "created_at":"收藏时间"<br/>}
// @Param   token       header     string  true  "token"
// @Param	body	body  models.GetbookinfoForm	true  "{<br/>"bookid":"图书编号-string"<br/>}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /getbookinfo [post]
func (this *BooksrackController) Getbookinfo() {
	var ob  *models.GetbookinfoForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	bookid:= ob.Bookid
	if bookid==""  {
		this.Rsp(false, "参数错误!","")
	}
	var conditions string = " "
	conditions+= " and r.bookid ='"+bookid+"'"
	book := models.MyBookconcernInfo(conditions)
	this.Rsp(true, "获取成功!",&book)
}


// @Title    查看用户信息
// @Summary  查看用户信息
// @Description 查看用户信息
// @Success 200 {<br/>"bookqid": "图书唯一编号",<br/>"userid": "用户编号",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"concernid":"收藏编号",<br/>"userid_to":"收藏人id", <br/>"userid_from":"图书编号",<br/>"concern_type":"收藏类型1:图书2:人",<br/> "created_at":"收藏时间"<br/>}
// @Param   token       header     string  true  "token"
// @Param	body	body  models.GetuserinfoForm	true  "{<br/>"userid":"用户编号"<br/>}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /getuserinfo [post]
func (this *BooksrackController) Getuserinfo() {
	var ob  *models.GetuserinfoForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	userid:= ob.Userid
	if userid==""  {
		this.Rsp(false, "参数错误!","")
	}
	var conditions string = " "
	conditions+= " and u.userid ='"+userid+"'"
	user := models.MyUserconcernInfo(conditions)
	this.Rsp(true, "获取成功!",&user)
}

// @Title    上架记录
// @Summary  上架记录
// @Description 上架记录
// @Success 200 {<br/>"bookqid": "图书唯一编号",<br/>"userid": "用户编号",<br/>"openid": "oX8vKwueTHOC3wrUkm2eJBnm-m6A",<br/>"bookid": "图书编号",<br/>"book_state": "状态1:上架;2:下架;3:待补充,4:删除",<br/>"is_borrow": "状态1:可借阅;2:不可借;",<br/>"create_time": "上架时间",<br/>"update_time": "修改时间",<br/>"bookname": "书名",<br/>"auhtor": "作者",<br/>"imageurl": "图书封面图",<br/>"imagehead": "图书正面图",<br/>"imageback": "图书背面图",<br/>"isbn": "图书条形码",<br/>"depreciation": "图书折旧",<br/>"price": "图书标价",<br/>"describe": "图书描述",<br/>"state": '状态 0非锁定状态 1：锁定状态',<br/>"wnickname": "微信昵称",<br/>"wimgurl": "微信头像",<br/>"nickname": "用户昵称",<br/>"imgurl": "用户头像",<br/>"gender": "性别1:男2:女0:保密",<br/>"age": "年龄",<br/>"telphone": "手机号",<br/>"password": "密码",<br/>"qq": "QQ号",<br/>"weibo": "微博号",<br/>"signature": "个性签名",<br/>"constellation": "星座",<br/>"province": "北京市",<br/>"city": "北京市",<br/>"address": "地址",<br/>"long":"经度",<br/>"lat":"纬度",<br/>"logintime": "用户最后登录时间",<br/>"created_at": "用户注册时间",<br/>"updated_at": "用户修改资料时间",<br/>"radius": "距离" <br/>}
// @Param   token       header     string  true  "token"
// @Param	body	  body 	 models.GetbookusersForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页"<br/> }"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /getmybooklist [post]
func (this *BooksrackController) Getmybooklist() {
	var ob  *models.GetbookusersForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	userid   :=   this.Userid
	if userid ==""{
		this.Rsp(false, "参数错误!","")
	}
	var conditions string = " "
	conditions+= " and u.userid ='"+userid+"'"
	conditions+= " and r.book_state ='1' "
	books,count := models.MyBooksrackList((draw-1)*length,length,conditions)
	if len(books) < 1 {
		books = []*models.BookrackList{}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&books}
	this.Rsp(true, "获取成功!",&json)
}


// @Title 扫条码获取图书
// @Summary  扫条码获取图书
// @Description 扫条码获取图书
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param	body	body  models.BookrackaddbysnForm	true  "{<br/>"isbn":"图书条码"<br/>}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /getbookbysn [post]
func (this *BooksrackController) Getbookbysn() {
	var ob  *models.BookrackaddbysnForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	isbn  := ob.Isbn
	userid := this.Userid
	if userid ==""|| isbn ==""{
		this.Rsp(false, "参数错误!","")
	}
	//查询数据库是否存在该条形码
	model,err:= models.GetIbsn(isbn)
	if err != nil {
		//state 1:扫码正常添加，3:扫码没有查询到结果待补充
		var state uint8 = 1
		res, berr:= common.GetBarcodeInfo(isbn)
		//查询失败
		if berr != nil{
			this.Rsp(false, "网络繁忙!","")
		}
		//查询不到信息
		if  res.Charge == false{
			state = 3
			res.Result.Showapi_res_body.GoodsName = "未知"
			res.Result.Showapi_res_body.ManuName = ""
			res.Result.Showapi_res_body.Img = ""
			res.Result.Showapi_res_body.Price = "0"
		}
		body :=  &res.Result.Showapi_res_body
		model := models.Books{}
		Uid := models.GetID()
		model.Bookid   = fmt.Sprintf("%d", Uid)
		model.Bookname = body.GoodsName
		model.Author   = body.ManuName
		model.Imagehead = ""
		model.Imageback = ""
		model.Imageurl  = body.Img
		model.Describe =body.GoodsName
		price, _:= strconv.ParseUint(body.Price, 10, 16)
		model.Price  = uint16(price)
		model.Isbn   = isbn
		model.Depreciation = 0
		model.State = state
		model.Userid = userid
		valid := validation.Validation{}
		valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
		valid.Required(model.Bookname, "bookname").Message("书名不能为空！")
		valid.Required(model.Isbn,"isbn").Message("条形码不能为空！")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				this.Rsp(false, err.Error(),"")
			}
		}
		err = comm.Insert(&model)
		if err != nil{
			this.Rsp(false,"添加失败!","")
		}
		if state==3{
			this.Rsp(false,"获取失败!","")
		}
	}
	if model.State == 3{
		this.Rsp(false,"获取失败!","")
	}
	this.Rsp(true,"获取成功!",&model)
}