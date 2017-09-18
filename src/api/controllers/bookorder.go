package controllers

import (
	comm "common/conndatabase"
	"models"
	"fmt"
	"time"
	"encoding/json"
	"strconv"
)

type BookorderController struct {
	ApiController
}

// @Title 订单列表
// @Summary 订单列表
// @Description  订单列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"\
// @Param   bookqid    formData   string  false   "用户书架图书唯一编号"
// @Param   isbn      formData   string  false  "图书条形码"
// @Param   orderid    formData   string  false   "订单编号"
// @Param   userid_from    formData   string  false   "书主人编号"
// @Param   userid_to    formData   string  false   "借书人编号"
// @Param   order_state    formData   string  false   "消息状态1:同意2:拒绝,3:完成，0：借书请求'"
// @Failure 500 服务器错误!
// @router /orderList [post]
func (this *BookorderController) Orderlist() {
	length, _ := this.GetInt("length",10) //获取分页步长
	draw, _ := this.GetInt("draw",1) //获取请求次数
	var conditions string = " "
	if v := this.GetString("bookid");v !="" {
	   conditions+= " and JSON_EXTRACT(books,'$.bookid') = '"+v+"'"
	}
	if v := this.GetString("bookstate");v !="" {
		conditions+= " and JSON_EXTRACT(books,'$.bookstate') = '"+v+"'"
	}
	if v := this.GetString("bookname");v != ""{
		conditions+= " and JSON_EXTRACT(books,'$.bookname') = '"+v+"'"
	}
	if v := this.GetString("author");v !="" {
		conditions+= " and JSON_EXTRACT(books,'$.author') = '"+v+"'"
	}
	if v := this.GetString("isbn");v !="" {
		conditions+= " and JSON_EXTRACT(books,'$.isbn') = '"+v+"'"
	}
	if v := this.GetString("orderid");v !="" {
		conditions+= " and orderid = '"+v+"'"
	}
	if v := this.GetString("userid_from");v !="" {
		conditions+= " and userid_from = '"+v+"'"
	}
	if v := this.GetString("userid_to");v !="" {
		conditions+= " and userid_to = '"+v+"'"
	}
	if v := this.GetString("order_state");v !="" {
		conditions+= " and order_state = '"+v+"'"
	}
	if v := this.GetString("bookqid");v !="" {
		conditions+= " and bookqid = '"+v+"'"
	}
	books :=models.BookOrderListData((draw-1)*length,length,conditions)
	var resPonse []interface{}
	book  := map[string]interface{}{}
	JsonBook  := map[string]interface{}{}
	JsonFrom  := map[string]interface{}{}
	JsonTo  := map[string]interface{}{}
	for _,val := range books{
		book["orderid"]     =  val.Orderid
		book["userid_from"] =  val.Userid_from
		book["userid_to"]   =  val.Userid_to
		book["bookqid"]     =  val.Bookqid
		book["order_state"] =  val.Order_state
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
	this.Rsp(true, "获取成功!",&resPonse)
}



// @Title  创建订单
// @Summary  创建订单
// @Description 创建订单
// @Success 200  {<br/>"userid":"用户编号", "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param   from   formData   string  true    "书主人用户编号"
// @Param   bookqid   formData   string  true  "书主人书架图书唯一编号"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderadd [post]
func (this *BookorderController) Orderadd() {
	bookqid  :=   this.GetString("bookqid")
	from    :=   this.GetString("from")
	to      :=   this.Userid
	if from =="" || bookqid=="" || to == ""{
		this.Rsp(false, "参数错误!","")
	}
	//查询书主人用户书架
	book,err:= models.GetUserBookRack(from,bookqid)
	if err != nil {
		this.Rsp(false, "用户不存在当前图书!","")
	}
	//书信息
	newBook := map[string]interface{}{}
	newBook["bookid"] = book.Bookid
	newBook["bookname"] = book.Bookname
	newBook["auhtor"] = book.Author
	newBook["imageurl"] = book.Imageurl
	newBook["imagehead"] = book.Imagehead
	newBook["imageback"] = book.Imageback
	newBook["isbn"] = book.Isbn
	newBook["depreciation"] = book.Depreciation
	newBook["price"] = book.Price
	newBook["describe"] = book.Describe
	newBook["state"] = book.State
	//查询书主人以及借书用户信息
	userInfo:= []string{from,to}
	toUser,uerr := models.GetUsersByIds(userInfo)
	if uerr != nil {
		this.Rsp(false, "借书人或书主人信息不存在!","")
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
	NewsInfo := models.Bookorder{}
	id := models.GetID()
	NewsInfo.Orderid =  fmt.Sprintf("%d", id)
	NewsInfo.Bookqid = book.Bookqid
	NewsInfo.Userid_from = from
	NewsInfo.Userid_to = to
	NewsInfo.Books  = string(res)
	NewsInfo.User_from = string(fty)
	NewsInfo.User_to = string(tty)
	NewsInfo.Order_state = 0
	t:= time.Now().Unix()
	NewsInfo.Create_time = t
	NewsInfo.Update_time = t
	toRes :=  comm.Insert(&NewsInfo)
    if toRes == nil {
		this.Rsp(true, "订单添加成功!",&NewsInfo)
	}
	this.Rsp(false, "未知错误!","")
}



// @Title    更改订单
// @Summary  更改订单
// @Description 更改订单
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token       header     string  true  "token"
// @Param   orderid   formData   string  true    "订单编号"
// @Param   order_state   formData   string  true   "消息状态状态1:同意2:拒绝,3:完成，0：借书请求"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderupdate [post]
func (this *BookorderController) Orderupdate() {
	orderid   :=   this.GetString("orderid")
	order_state  :=   this.GetString("order_state")
	if order_state == "" || orderid ==""{
		this.Rsp(false, "参数错误!","")
	}
	model := models.Bookorder{}
	model.Orderid = orderid
	if err := comm.Read(&model);err == nil {
		i64, err := strconv.ParseUint(order_state, 10, 8)
		model.Order_state = uint8(i64)
		t:= time.Now().Unix()
		model.Update_time = t
		err = comm.Update(&model)
		if err == nil {
			this.Rsp(true, "当前消息状态已修改!",&model)
		}
	}
	this.Rsp(false, "修改失败!","")
}