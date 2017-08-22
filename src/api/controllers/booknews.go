package controllers

import (
	comm "common/conndatabase"
	"common"
	"models"
	"fmt"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type BooknewsController struct {
	ApiController
}

// @Title 借书消息列表
// @Summary  借书消息列表
// @Description  借书消息列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   length     formData   string  false      "获取分页步长"
// @Param   draw       formData   string  false      "当前页"
// @Param   userid     formData   string  true      "用户编号"
// @Param   bookqid    formData   string  false     "图书唯一编号"
// @Param   newid      formData   string  false       "消息编号"
// @Param   order_state   formData   string  false   "消息状态1:同意2:拒绝,3:完成，0：借书请求'"
// @Failure 500 服务器错误!
// @router /newsList [post]
func (this *BooknewsController) Newslist() {
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var conditions string = " "
	if v := this.GetString("userid");v != ""{
		conditions+= " and userid_from ='"+v+"'"
	}
	if v := this.GetString("bookqid");v !="" {
		conditions+= " and bookqid= '"+v+"'"
	}
	if v := this.GetString("newid");v !="" {
		conditions+= " and newid = '"+v+"'"
	}
	if 	v := this.GetString("order_state");v !="" {
		conditions+= " and order_state = '"+v+"'"
	}
	var start int = 0
	if draw  > 0 {
		start = (draw-1)*length
	}
	books,totalItem := models.BooknewsListData(start,length,conditions)
	var resPonse []interface{}
	book  := map[string]interface{}{}
	JsonBook  := map[string]interface{}{}
	JsonFrom  := map[string]interface{}{}
	JsonTo  := map[string]interface{}{}
	for _,val := range books{
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
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":resPonse}
	this.renderJson(Json)
}



// @Title    发起借书请求
// @Summary  发起借书请求
// @Description 发起借书请求
// @Success 200  {<br/>"userid":"用户编号","bookqid":"图书唯一编号","bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   from   formData   string  true    "书主人用户编号"
// @Param   to   formData   string  true      "借书人用户编号"
// @Param   bookqid   formData   string  true  "书主人书架图书唯一编号"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /libraryrequest [post]
func (this *BooknewsController) Libraryrequest() {
	bookqid  :=   this.GetString("bookqid")
	from    :=   this.GetString("from")
	to      :=   this.GetString("to")
	if from =="" || bookqid=="" || to == ""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	//查询书主人用户书架
	book,err:= models.GetUserBookRack(from,bookqid)
	if err != nil {
		common.ErrSystem.Message = "用户不存在当前图书!"
		this.renderJson(common.ErrSystem)
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
	newBook["state"] = book.State
	//查询书主人以及借书用户信息
	userInfo:= []string{from,to}
	toUser,uerr := models.GetUsersByIds(userInfo)
	if uerr != nil {
		common.ErrSystem.Message = "借书人或书主人信息不存在!"
		this.renderJson(common.ErrSystem)
	}
	UserFrom := map[string]interface{}{}
	UserTo   := map[string]interface{}{}
	for _,val := range toUser{
		if val.Userid == from{
			UserFrom["userid"] = val.Userid
			UserFrom["nickname"] = val.Nickname
			UserFrom["openid"] = val.Openid
			UserFrom["imgurl"] = val.Imgurl
			UserFrom["gender"] = val.Gender
			UserFrom["age"] = val.Age
			UserFrom["signature"] = val.Signature
		}
		if val.Userid == to{
			UserTo["userid"] = val.Userid
			UserTo["nickname"] = val.Nickname
			UserTo["openid"] = val.Openid
			UserTo["imgurl"] = val.Imgurl
			UserTo["gender"] = val.Gender
			UserTo["age"] = val.Age
			UserTo["signature"] = val.Signature
		}
	}
	fty,ferr:= json.Marshal(&UserFrom)
	if ferr != nil{
		common.ErrSystem.Message = "未知错误!"
		this.renderJson(common.ErrSystem)
	}
	tty,terr:= json.Marshal(&UserTo)
	if terr != nil{
		common.ErrSystem.Message = "未知错误!"
		this.renderJson(common.ErrSystem)
	}
	res,rerr:= json.Marshal(&newBook)
	if rerr != nil{
		common.ErrSystem.Message = "未知错误!"
		this.renderJson(common.ErrSystem)
	}
	//借书人消息
	FromInfo := models.Booknews{}
	id := models.GetID()
	FromInfo.Newid =  fmt.Sprintf("%d", id)
	FromInfo.Bookqid = book.Bookqid
	FromInfo.Userid_from = from
	FromInfo.Userid_to = to
	FromInfo.Books  = string(res)
	FromInfo.User_from = string(fty)
	FromInfo.User_to = string(tty)
	FromInfo.Order_type = 1
	FromInfo.Order_state = 0
	t:= time.Now().Unix()
	FromInfo.Create_time = t
	FromInfo.Update_time = t
   //书主人消息
	toInfo := FromInfo
	toInfo.Userid_from =  to
	toInfo.Userid_to  =  from
	toInfo.Order_type =   2
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
			common.Actionsuccess.Message ="借书请求发送成功!"
			common.Actionsuccess.MoreInfo = &FromInfo
			this.renderJson(common.Actionsuccess)
		}
	}
	common.ErrSystem.Message = "未知错误,消息丢失!"
	this.renderJson(common.ErrSystem)
}



// @Title    更改借书请求状态
// @Summary  更改借书请求状态
// @Description 更改借书消息状态
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   newid   formData   string  true    "消息编号"
// @Param   order_state   formData   string  true   "消息状态1:同意借书2:拒绝借书"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /libraryrequestupdate [post]
func (this *BooknewsController) Libraryrequestupdate() {
	newid   :=   this.GetString("newid")
	order_state  :=   this.GetString("order_state")
	if  order_state == "" || newid ==""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	model := models.Booknews{}
	Order := models.Bookorder{}
	model.Newid = newid
	err := comm.Read(&model)
	if err == nil {
		t:= time.Now().Unix()
		model.Update_time = t
		Order.Order_state = model.Order_state
		Order.Update_time = t
		Order.Orderid =  newid
		var flag bool = true
		if  model.Order_state == 0{
			Order.Userid_to = model.Userid_to
			Order.Userid_from = model.Userid_from
			Order.Books   =  model.Books
			Order.Bookqid =  model.Bookqid
			Order.User_from = model.User_from
			Order.User_to  = model.User_to
			Order.Create_time = t
		}else{
			flag = false
		}
		i64, err := strconv.ParseUint(order_state, 10, 8)
		Order_state:=uint8(i64)
		model.Order_state = Order_state
		o := orm.NewOrm()
		err = o.Begin()
		if err == nil{
			_,err =  o.Update(&model,"Order_state","Update_time")
			if flag == true{
				_,err = o.Insert(&Order)
			}else{
				_,err = o.Update(&Order,"Order_state","Update_time")
			}
			err = o.Commit()
			if err != nil{
				err = o.Rollback()
			}
			if err == nil{
				common.Actionsuccess.Message ="当前消息状态已修改!"
				common.Actionsuccess.MoreInfo = &model
				this.renderJson(common.Actionsuccess)
			}
		}
	}
	common.ErrSystem.Message = fmt.Sprint(err)
	this.renderJson(common.ErrSystem)
}