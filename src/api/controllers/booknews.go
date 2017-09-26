package controllers

import (
	"models"
	"fmt"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"math"
)

type BooknewsController struct {
	ApiController
}

// @Title    借书消息列表
// @Summary  借书消息列表
// @Description  借书消息列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token      header     string  true     "token"
// @Param	body	body 	 models.NewslistForm  true   "{ <br/>"length":"获取分页步长", <br/>"draw":"当前页",<br/> "order_state":"消息状态1:同意2:拒绝,3:完成，0：借书请求'<br/>}"
// @Failure 500 服务器错误!
// @router /newsList [post]
func (this *BooknewsController) Newslist() {
	var ob  *models.NewslistForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	var conditions string = " "
	conditions+= " and userid_from ='"+this.Userid+"'"
	if ob.OrderState!="" {
		conditions+= " and order_state = '"+ob.OrderState+"'"
	}
	var resPonse []interface{}
	books,count := models.BooknewsListDataBack((draw-1) * length,length,conditions)
	if len(books) >= 1 {
		for _,val := range books{
			book  := map[string]interface{}{}
			JsonBook  := map[string]interface{}{}
			JsonFrom  := map[string]interface{}{}
			JsonTo  := map[string]interface{}{}
			book["newid"] =  val.Newid
			book["userid_from"] =  val.Userid_from
			book["userid_to"]   =  val.Userid_to
			book["bookqid"]     =  val.Bookqid
			book["order_state"] =  val.Order_state
			book["order_type"] =  val.Order_type
			book["create_time"] =  val.Create_time
			book["update_time"] =  val.Update_time
			Books := []byte(val.Books)
			err := json.Unmarshal(Books, &JsonBook)
			if err == nil{
				book["books"] = &JsonBook
			}else{
				book["books"] = ""
			}
			UserFrom := []byte(val.User_from)
			err = json.Unmarshal(UserFrom, &JsonFrom)
			if err == nil{
				book["user_from"] =  JsonFrom
			}else{
				book["user_from"] =  ""
			}
			UserTo := []byte(val.User_to)
			err = json.Unmarshal(UserTo, &JsonTo)
			if err == nil{
				book["user_to"]   =  JsonTo
			}else{
				book["user_to"]   =  ""
			}
			resPonse = append(resPonse,&book)
		}
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&resPonse}
	this.Rsp(true, "获取成功!",&json)
}



// @Title    发起借书请求
// @Summary  发起借书请求
// @Description 发起借书请求
// @Success 200  {<br/>"userid":"用户编号","bookqid":"图书唯一编号","bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token   header   string  true  "token"
// @Param	body	body     models.LibraryrequestForm 	true   "{ <br/>"bookqid":"图书唯一编号","order_type":"类型:1:别人借我的书;2:我借别人的书", <br/>"order_state":"状态0:借书请求;1:接受2:拒绝"<br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /libraryrequest [post]
func (this *BooknewsController) Libraryrequest() {
	var ob  *models.LibraryrequestForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	bookqid :=   ob.Bookqid
	to      :=   this.Userid
	if  bookqid=="" || to == "" || ob.Telphone==""|| ob.Qq=="" || ob.Wechat== ""{
		this.Rsp(false, "参数错误!","")
	}
	//查询书主人用户书架
	book,err:= models.GetUserBookRack(bookqid)
	if err != nil {
		this.Rsp(false, "不存在当前图书!","")
	}
	//书信息
	newBook := map[string]interface{}{}
	newBook["bookid"] = book.Bookid
	newBook["bookname"] = book.Bookname
	newBook["auhtor"] = book.Author
	newBook["imageurl"] = book.Imageurl
	newBook["imagehead"] = book.Imagehead
	newBook["imageback"] = book.Imageback
	newBook["isbn"]  = book.Isbn
	newBook["depreciation"] = book.Depreciation
	newBook["price"] = book.Price
	newBook["describe"] = book.Describe
	newBook["state"]   =  book.State
	//查询书主人以及借书用户信息
	userInfo:= []string{book.Userid,to}
	toUser,uerr := models.GetUsersByIds(userInfo)
	if uerr != nil {
		this.Rsp(false, "借书人或书主人信息不存在!","")
	}
	UserFrom := map[string]interface{}{}
	UserTo   := map[string]interface{}{}
	for _,val := range toUser{
		if val.Userid == book.Userid {
			UserFrom["userid"] = val.Userid
			UserFrom["nickname"] = val.Nickname
			UserFrom["openid"] = val.Openid
			UserFrom["imgurl"] = val.Imgurl
			UserFrom["telphone"]= val.Telphone
			UserFrom["qq"]      = val.Qq
			UserFrom["wechat"]  = val.Wechat
			UserFrom["gender"] = val.Gender
			UserFrom["age"]  = val.Age
			UserFrom["signature"] = val.Signature
		}
		if val.Userid == to{
			UserTo["userid"] = val.Userid
			UserTo["nickname"] = val.Nickname
			UserTo["openid"] = val.Openid
			UserTo["imgurl"] = val.Imgurl
			UserTo["telphone"] = ob.Telphone
			UserTo["qq"]       = ob.Qq
			UserTo["wechat"]   = ob.Wechat
			UserTo["gender"]   = val.Gender
			UserTo["age"]      = val.Age
			UserTo["signature"] = val.Signature
		}
	}
	fty,ferr:= json.Marshal(&UserFrom)
	if ferr != nil{
		this.Rsp(false, "未知错误!","")
	}
	tty,terr:= json.Marshal(&UserTo)
	if terr != nil{
		this.Rsp(false, "未知错误!","")
	}
	res,rerr:= json.Marshal(&newBook)
	if rerr != nil{
		this.Rsp(false, "未知错误!","")
	}
	//书主人消息
	FromInfo := models.Booknews{}
	id := models.GetID()
	FromInfo.Newid =  fmt.Sprintf("%d", id)
	FromInfo.Bookqid = book.Bookqid
	FromInfo.Userid_from = book.Userid
	FromInfo.Userid_to = to
	FromInfo.Books  = string(res)
	FromInfo.User_from = string(fty)
	FromInfo.User_to = string(tty)
	FromInfo.Order_type = 1
	FromInfo.Order_state = 0
	t:= time.Now().Unix()
	FromInfo.Create_time = t
	FromInfo.Update_time = t
    //借书人消息
	toInfo := FromInfo
	toInfo.Userid_from =  to
	toInfo.Userid_to   =  book.Userid
	toInfo.User_from   =  FromInfo.User_to
	toInfo.User_to     =  FromInfo.User_from
	toInfo.Order_type  =  2
	o := orm.NewOrm()
	err = o.Begin()
	if err == nil{
		_,err = o.Insert(&FromInfo)
		_,err = o.Insert(&toInfo)
		 err = o.Commit()
		if err != nil{
			err = o.Rollback()
		}
		if err == nil{
			this.Rsp(true, "借书请求发送成功!",FromInfo.Newid)
		}
	}
	this.Rsp(false, "未知错误,消息丢失!","")
}



// @Title    同意借书请求
// @Summary  同意借书请求
// @Description 同意借书请求
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token   header   string  true  "token"
// @Param	body	body     models.AgreeLibraryrequestForm 	true   "{ <br/>"newid":"消息编号"<br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /agreelibraryrequest [post]
func (this *BooknewsController) Agreelibraryrequest() {
	var ob  *models.AgreeLibraryrequestForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if   ob.Newid ==""|| ob.Qq=="" || ob.Wechat=="" || ob.Telphone==""{
		this.Rsp(false, "参数错误!","")
	}
	model,err := models.BooknewsInfo(" and newid="+ob.Newid+" and userid_from="+this.Userid)
	if err!=nil{
		this.Rsp(false, "无法更改当前消息!","")
	}
	var User_from  map[string]interface{}
	err = json.Unmarshal([]byte(model.User_from),&User_from)
	User_from["qq"]       = ob.Qq
	User_from["wechat"]   = ob.Wechat
	User_from["telphone"] = ob.Telphone
	from,err := json.Marshal(User_from)
	model.User_from = string(from)
	OrderTo   := models.Bookorder{}
	OrderFrom := models.Bookorder{}
	if err == nil {
		t:= time.Now().Unix()
		model.Update_time = t
		OrderTo.Orderid     =  ob.Newid
		OrderTo.Order_state =  model.Order_state
		OrderTo.Update_time =  t
		OrderTo.Userid_to   =  model.Userid_to
		OrderTo.Userid_from =  model.Userid_from
		OrderTo.Books       =  model.Books
		OrderTo.Bookqid     =  model.Bookqid
		OrderTo.User_from   =  model.User_from
		OrderTo.User_to     =  model.User_to
		OrderTo.Order_type  =  model.Order_type
		OrderTo.Create_time =  t
		OrderTo.Update_time =  t
		OrderFrom =  OrderTo
		OrderFrom.Userid_to   =  OrderTo.Userid_from
		OrderFrom.Userid_from =  OrderTo.Userid_to
		OrderFrom.User_from   =  OrderTo.User_to
		OrderFrom.User_to     =  OrderTo.User_from
		OrderFrom.Order_type  =  2
		model.Order_state = 1
		o := orm.NewOrm()
		err = o.Begin()
		if err == nil{
			Booknews := new(models.Booknews)
			_, err = o.QueryTable(Booknews).Filter(
				"newid", model.Newid).Filter("userid_from", model.Userid_from).Update(orm.Params{
				"order_state": model.Order_state,"update_time":model.Update_time,
				"user_to":model.User_to,"user_from":model.User_from,
			})
			_, err = o.QueryTable(Booknews).Filter(
				"newid", model.Newid).Filter("userid_to", model.Userid_from).Update(orm.Params{
				"order_state": model.Order_state,"update_time":model.Update_time,
				"user_to":model.User_from,"user_from":model.User_to,
			})
			_,err = o.Insert(&OrderTo)
			_,err = o.Insert(&OrderFrom)
			err = o.Commit()
			if err != nil{
				err = o.Rollback()
			}
			if err == nil{
				this.Rsp(true, "当前消息状态已修改!",&model.Newid)
			}
		}
	}
	this.Rsp(false, "未知错误!","")
}


// @Title    拒绝借书请求
// @Summary   拒绝借书请求
// @Description  拒绝借书请求
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token   header   string  true  "token"
// @Param	body	body     models.RefuseLibraryrequestForm 	true   "{ <br/>"newid":"消息编号"<br/>}"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /refuselibraryrequest [post]
func (this *BooknewsController) Refuselibraryrequest() {
	var ob  *models.RefuseLibraryrequestForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	newid   :=   ob.Newid
	if newid ==""{
		this.Rsp(false, "参数错误!","")
	}
	model,err := models.BooknewsInfo(" and newid="+newid+" and userid_from="+this.Userid)
	if err!=nil{
		this.Rsp(false, "未知错误!","")
	}
	model.Update_time = time.Now().Unix()
	model.Order_state = 2
	o := orm.NewOrm()
	_,err =  o.Update(&model,"Order_state","Update_time")
	if err == nil{
		this.Rsp(true, "已拒绝借书请求!",&model.Newid)
	}
	this.Rsp(false, "未知错误!","")
}