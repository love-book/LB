package controllers

import (
	comm "common/conndatabase"
	"github.com/astaxie/beego/validation"
	"common"
	"models"
	"fmt"
	"strconv"
	"github.com/astaxie/beego/orm"
	"time"
	"encoding/json"
	"math"
)

type BooksController struct {
	ApiController
}

// Title 获取图书列表
// Summary  获取图书列表
// Description 获取图书列表
// Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// Param   length    formData   int  false      "分页步长"
// Param   draw      formData   int  false      "请求页数"
// Param   bookname  formData   string  false   "书名"
// Param   author    formData   string  false   "作者"
// Param   bookid    formData   string  false   "图书编号"
// Param   isbn     formData   string  false  "图书条形码"
// Param   state    formData   string  false    "状态1:上架;2:下架;3:待补充"
// Failure 500 服务器错误!
// router /booklist [post]
func (this *BooksController) BookList() {
	length, _ := this.GetInt("length",10) //获取分页步长
	draw, _ := this.GetInt("draw",1) //获取请求次数
	var conditions string = " "
	if v := this.GetString("bookid");v != ""{
		conditions+= " and bookid ='"+v+"'"
	}
	if v := this.GetString("bookname");v != ""{
		conditions+= " and bookname ="+v
	}
	if v := this.GetString("author");v !="" {
		conditions+= " and author = "+v
	}
	if v := this.GetString("isbn");v !="" {
		conditions+= " and isbn = "+v
	}
	if v := this.GetString("state");v !="" {
		conditions+= " and state = "+v
	}
	/*
	if starttime := this.GetString("starttime");starttime !="" {
		tm1, _ := time.Parse("01/02/2006", starttime)
		if endtime := this.GetString("endtime");endtime == ""{
			endtime = fmt.Sprintf("%d",time.Now().Unix())
		}else{
			tm2, _ := time.Parse("01/02/2006", endtime)
			endtime = fmt.Sprintf("%d",tm2)
		}
		starttime = fmt.Sprintf("%d",tm1)
		conditions+= " and created_at  bettwen "+starttime+" and "+endtime
	}*/
	books :=models.GetBookList((draw-1) * length,length,conditions)
	this.Rsp(true, "获取成功!",&books)
}


// Title 修改书籍
// Description 修改书籍
// Summary  修改书籍
// Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// Param   bookid    formData   string  true    "图书编号"
// Param   bookname   formData   string  false    "书名"
// Param   author     formData   string  false   "作者"
// Param   imageurl  formData   string  false    "图书封面图"
// Param   imagehead    formData   string  false     "图书正面图"
// Param   imageback       formData   string  false    "图书背面图"
// Param   isbn      formData   string  false    "条形码"
// Param   describe      formData   string  false    "图书简介"
// Param   price    formData   string  false    "标价"
// Param   state    formData   string  false    "状态1:上架;2:下架;3:待补充"
// Failure 100 错误提示信息!
// Failure 500 服务器错误!
// router /bookupdate [post]
func (this *BooksController) Bookupdate(){
	model := models.Books{}
	model.Bookid = this.GetString("bookid")
	valid := validation.Validation{}
	valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.Rsp(false, err.Error(),"")
		}
	}
	if err:= comm.Read(&model);err==nil{
		if v := this.GetString("bookname");v != ""{
			model.Bookname = v
		}
		if v := this.GetString("author");v != ""{
			model.Author = v
		}
		if v := this.GetString("imagehead");v != ""{
			model.Imagehead = v
		}
		if v := this.GetString("imageback");v != ""{
			model.Imageback = v
		}
		if v := this.GetString("imgurl");v != ""{
			model.Imageurl = v
		}
		if v := this.GetString("isbn");v != ""{
			model.Isbn = v
		}
		if v ,_ := this.GetUint16("price");v != 0{
			model.Price = v
		}
		if v := this.GetString("describe");v != ""{
			model.Describe = v
		}
		if v ,_:= this.GetUint8("state");v != 0{
			model.State = v
		}
		if res:=models.UpdateBookById(&model);res == nil{
			this.Rsp(true, "修改成功!",&model)
		}
	}
	this.Rsp(false, "修改失败!","")
}



// @Title 添加书籍
// @Summary  添加书籍
// @Description 添加书籍
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"折旧",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param	body	body  models.BookaddForm  true  {参数含义参考返回值}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookadd [post]
func (this *BooksController) Bookadd() {
	var ob  *models.BookaddForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	model := models.Books{}
	id := models.GetID()
	model.Bookid    =  fmt.Sprintf("%d", id)
	model.Bookname  = ob.Bookname
	model.Author    = ob.Author
	model.Imagehead = ob.Imagehead
	model.Imageback = ob.Imageback
	model.Imageurl  = ob.Imageurl
	model.Describe  = ob.Describe
	model.Price  = ob.Price
	model.Isbn   =  ob.Isbn
	model.Userid = this.Userid
	model.Depreciation = 0
	model.State = 1
	valid := validation.Validation{}
	valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
	valid.Required(model.Bookname, "bookname").Message("书名不能为空！")
	valid.Required(model.Author,"author").Message("作者不能为空！")
	valid.Required(model.Author,"isbn").Message("条形码不能为空！")
	valid.Required(model.Userid,"userid").Message("用户编号不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.Rsp(false,err.Error(),"")
		}
	}
	//查询数据库是否存在该条形码
	book,res:= models.GetIbsn(model.Isbn)
	if res == nil{
		this.Rsp(true,"当前图书已添加",&book)
	}
	var bookRack  = models.Bookrack{}
	Bid := models.GetID()
	bookRack.Bookqid   =  fmt.Sprintf("%d", Bid)
	bookRack.Bookid = model.Bookid
	bookRack.Userid = model.Userid
	bookRack.Book_state = "1"
	bookRack.Is_borrow = "1"
	bookRack.Create_time = time.Now().Unix()
	//事物提交
	o := orm.NewOrm()
	err := o.Begin()
	if err == nil{
		_,toRes :=  o.Insert(&model)
		_,fromRes := o.Insert(&bookRack)
		if toRes==nil && fromRes==nil{
			err = o.Commit()
		}else{
			err = o.Rollback()
		}
	}
	if err != nil {
		this.Rsp(false,"未知错误","")
	}
	this.Rsp(true,"图书添加成功",&model)
}



// Title 条形码添加书籍
// Summary  条形码添加书籍
// Description 条形码添加书籍
// Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// Param   token       header     string  true  "token"
// Param  userid    formData   string  true    "用户id"
// Param   isbn      formData   string  true    "条形码"
// Failure 100 错误提示信息!
// Failure 500 服务器错误!
// router /bookaddbycode [post]
func (this *BooksController) Bookaddbycode() {
	Isbn := this.GetString("isbn")
	if Isbn == ""{
		this.Rsp(false,  "条形码不能为空！","")
	}

	//查询数据库是否存在该条形码
	book,re:= models.GetIbsn(Isbn)
	if re == nil{
		this.Rsp(true,  "当前图书已添加",&book)
	}
	//state 1:扫码正常添加，3:扫码没有查询到结果待补充
	var state uint8 = 1
	res, berr:= common.GetBarcodeInfo(Isbn)
	//查询失败
	if berr != nil{
		this.Rsp(false,  "网络繁忙!","")
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
	Uid := models.GetID()
	model := models.Books{}
	model.Bookid   = fmt.Sprintf("%d", &Uid)
	model.Bookname = body.GoodsName
	model.Author   = body.ManuName
	model.Imagehead = ""
	model.Imageback = ""
	model.Imageurl = body.Img
	model.Describe = body.GoodsName
	price, _:= strconv.ParseUint(body.Price, 10, 16)
	model.Price  = uint16(price)
	model.Isbn   = Isbn
	model.Depreciation = 0
	model.State = state
	model.Userid = this.GetString("userid")
	valid := validation.Validation{}
	valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
	valid.Required(model.Bookname, "bookname").Message("书名不能为空！")
    valid.Required(model.Isbn,"isbn").Message("条形码不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.Rsp(false,  err.Error(),"")
		}
	}
	var bookRack  = models.Bookrack{}
	Bid := models.GetID()
	bookRack.Bookqid   =  fmt.Sprintf("%d", Bid)
	bookRack.Bookid = model.Bookid
	bookRack.Userid = model.Userid
	bookRack.Book_state = "1"
	bookRack.Is_borrow = "1"
	bookRack.Create_time = time.Now().Unix()
	//事物提交
	o := orm.NewOrm()
	err := o.Begin()
	if err == nil{
		_,toRes :=  o.Insert(&model)
		_,fromRes := o.Insert(&bookRack)
		if toRes==nil && fromRes==nil{
			err = o.Commit()
			this.Rsp(true,"当前图书添加成功!",&model)
		}else{
			err = o.Rollback()
		}
	}
	if err != nil {
		this.Rsp(false,  err.Error(),"")
	}
	this.Rsp(true,"当前图书添加成功!",&model)
}





// @Title   收藏人/书籍
// @Summary  收藏人/书籍
// @Description 收藏人/书籍
// @Success 200  {<br/> "userid_to": "收藏人编号",<br/> "userid_from": "书主人/图书编号",<br/> "concern_type": "1:收藏书籍;2:收藏人",<br/> "created_at": "收藏时间" <br/>}
// @Param   token   header   string   true   "token"
// @Param	body	body     models.AddconcernForm  true  {参数含义参考返回值}
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /addconcern [post]
func (this *BooksController) Addconcern(){
	var ob  *models.AddconcernForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if ob.UseridFrom==""||ob.ConcernType==""{
		this.Rsp(false, "参数错误!","")
	}
	// 查询是否已已经收藏
	c := models.Concern{}
	c.UseridTo = this.Userid
	c.ConcernType = ob.ConcernType
	c.UseridFrom  = ob.UseridFrom
	var conditions string =""
	conditions+=" and userid_to='"+c.UseridTo+"'"
	conditions+=" and userid_from='"+c.UseridFrom+"'"
	conditions+=" and concern_type='"+c.ConcernType+"'"
	concern,err :=models.GetConcern(conditions)
	if err==nil{
		this.Rsp(true, "你已经收藏过了!",&concern)
	}
	//查询书籍或者用户信息
	if c.ConcernType=="1"{
		b,err:=models.GetBookInfo(" and bookid='"+c.UseridFrom+"'")
		if err!=nil{
			this.Rsp(false, "当前图书不存在!","")
		}
		bk,_:=json.Marshal(&b)
		c.Books = string(bk)
	}else if(c.ConcernType=="2"){
		b,err:=models.GetUsersById(c.UseridFrom)
		if err!=nil{
			this.Rsp(false, "当前用户不存在!","")
		}
		bk,_:=json.Marshal(&b)
		c.Books = string(bk)
	}
	cid := models.GetID()
	c.Concernid  = fmt.Sprintf("%d", cid)
	_, err = models.AddConcern(&c)
	if err == nil {
		this.Rsp(true, "成功!",&c)
	} else {
		this.Rsp(false, "失败!","")
	}
}



// @Title    批量删除收藏
// @Summary  批量删除收藏
// @Description 批量删除收藏
// @Success 200  { <br/>"concernid": "收藏编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充4:删除",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   token   header     string  true  "token"
// @Param	body	body     models.DelbookconcernForm  true   "[{"concernid": "收藏1"},{"concernid": "收藏2"}]"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /delbookconcern [post]
func (this *BooksController) Delbookconcern() {
	var ob  []*models.DelbookconcernForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if len(ob)<=0{
		this.Rsp(false, "参数错误!","")
	}
	fmt.Println(ob)
	return
	userid   :=   this.Userid
	for _,v:= range ob{
		c,err:= models.GetConcernById(v.Concernid)
		if err == nil {
			if c.UseridTo == userid{
				models.DelConcernById(v.Concernid)
			}
		}else{
			this.Rsp(false,"收藏不存在","")
		}
	}
	this.Rsp(true,"操作成功!","")
}


// @Title    收藏的图书列表
// @Summary  收藏的图书列表
// @Description 收藏的图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token   header     string  true  "token"
// @Param	body	body 	 models.ConcernBookListForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页"<br/> }"
// @Failure 500 服务器错误!
// @router /concernbooklist [post]
func (this *BooksController) ConcernBookList() {
	var ob  *models.ConcernBookListForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	var conditions string = " "
	conditions+= " and userid_to ='"+this.Userid+"'"
	conditions+=" and concern_type ='1' "
	lists,count := models.GetConcernList((draw-1)*length,length,conditions)
	pageTotal:= math.Ceil(float64(count)/float64(length))
	if len(lists) < 1 {
		lists = []*models.Concern{}
		json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&lists}
		this.Rsp(true, "获取成功!",&json)
	}else{
		resPonse := []interface{}{}
		for _,v:= range lists{
			concerns := map[string]interface{}{}
			b := map[string]interface{}{}
			concerns["concernid"] = v.Concernid
			concerns["userid_to"] = v.UseridTo
			concerns["userid_from"] = v.UseridFrom
			concerns["concern_type"] = v.ConcernType
			concerns["created_at"] = v.CreatedAt
			err:=json.Unmarshal([]byte(v.Books),&b)
			if err == nil{
				concerns["books"] = &b
			}else{
				concerns["books"] = ""
			}
			resPonse = append(resPonse,&concerns)
		}
		json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&resPonse}
		this.Rsp(true, "获取成功!",&json)
	}
}


// @Title    收藏的用户列表
// @Summary  收藏的用户列表
// @Description 收藏的图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param	body	body 	 models.ConcernBookListForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页"<br/> }"
// @Failure 500 服务器错误!
// @router /concernuserlist [post]
func (this *BooksController) ConcernUserList() {
	var ob  *models.ConcernBookListForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	var conditions string = " "
	conditions+= " and userid_to ='"+this.Userid+"'"
	conditions+=" and concern_type ='2' "
	lists,count := models.GetConcernList((draw-1)*length,length,conditions)
	pageTotal:= math.Ceil(float64(count)/float64(length))
	if len(lists) < 1 {
		lists = []*models.Concern{}
		json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&lists}
		this.Rsp(true, "获取成功!",&json)
	}else{
		var resPonse []interface{}
		for _,v:= range lists{
			concerns := map[string]interface{}{}
			b := map[string]interface{}{}
			concerns["concernid"] = v.Concernid
			concerns["userid_to"] = v.UseridTo
			concerns["userid_from"] = v.UseridFrom
			concerns["concern_type"] = v.ConcernType
			concerns["created_at"] = v.CreatedAt
			err:=json.Unmarshal([]byte(v.Books),&b)
			if err == nil{
				concerns["books"] = &b
				//根据用户编号获取其书架最新的三本书
				b:=models.BooksrackList(0,3," and u.userid ='"+v.UseridFrom+"'")
				concerns["booksList"] = &b
			}else{
				concerns["books"] = ""
				concerns["booksList"] =""
			}
			resPonse = append(resPonse,&concerns)
		}
		json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&resPonse}
		this.Rsp(true, "获取成功!",&json)
	}
}
