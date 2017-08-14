package controllers

import (
	comm "common/conndatabase"
	"common"
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
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"
// @Param   isbn      formData   string  false  "图书条形码"
// @Param   orderid    formData   string  false   "订单编号"
// @Param   userid_from    formData   string  false   "书主人编号"
// @Param   userid_to    formData   string  false   "借书人编号"
// @Param   order_state    formData   string  false   "消息状态1:同意2:拒绝,3:完成，0：借书请求'"
// @Failure 500 服务器错误!
// @router /orderList [post]
func (this *BookorderController) Orderlist() {
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var conditions string = " "
	bookid := this.GetString("bookid")
	if bookid !="" {
	   conditions+= " and JSON_EXTRACT(books,'$.bookid') = '"+bookid+"'"
	}
	bookstate := this.GetString("bookstate")
	if bookstate !="" {
		conditions+= " and JSON_EXTRACT(books,'$.bookstate') = '"+bookstate+"'"
	}
	bookname := this.GetString("bookname")
	if bookname != ""{
		conditions+= " and JSON_EXTRACT(books,'$.bookname') = '"+bookname+"'"
	}
	author := this.GetString("author")
	if author !="" {
		conditions+= " and JSON_EXTRACT(books,'$.author') = '"+author+"'"
	}
	isbn := this.GetString("isbn")
	if isbn !="" {
		conditions+= " and JSON_EXTRACT(books,'$.isbn') = '"+isbn+"'"
	}
	orderid := this.GetString("orderid")
	if orderid !="" {
		conditions+= " and orderid = '"+orderid+"'"
	}
	userid_from := this.GetString("userid_from")
	if userid_from !="" {
		conditions+= " and userid_from = '"+userid_from+"'"
	}
	userid_to := this.GetString("userid_to")
	if userid_to !="" {
		conditions+= " and userid_to = '"+userid_to+"'"
	}
	order_state := this.GetString("order_state")
	if order_state !="" {
		conditions+= " and order_state = '"+order_state+"'"
	}
	var start int = 0
	if draw  > 0 {
		start = (draw-1)*length
	}
	var books []models.Bookorder
	conditions += "  order by pushtime desc"
	var  TableName = "lb_bookorder"
	totalItem, res :=models.GetPagesInfo(TableName,start,length,conditions,"*")
	res.QueryRows(&books)
	resPonse := map[int]interface{}{}
	book  := map[string]interface{}{}
	JsonBook  := map[string]interface{}{}
	JsonRes  := map[string]interface{}{}
	for key,val := range books{
		book["orderid"] =  val.Orderid
		book["userid_from"] =  val.Userid_from
		book["userid_to"] =  val.Userid_to
		book["order_state"] =  val.Order_state
		book["pushtime"] =  val.Pushtime
		Books := []byte(val.Books)
		err := json.Unmarshal(Books, &JsonBook)
		if err == nil{
			book["books"] = &JsonBook
		}else{
			book["books"] = ""
		}
		Users := []byte(val.Users)
		err = json.Unmarshal(Users, &JsonRes)
		if err == nil{
			book["from"] =  JsonRes["from"]
			book["to"]   =  JsonRes["to"]
		}else{
			book["from"] =  ""
			book["to"]   =  ""
		}
		resPonse[key] = &book
	}
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":resPonse}
	this.renderJson(Json)
}



// @Title  创建订单
// @Summary  创建订单
// @Description 创建订单
// @Success 200  {<br/>"userid":"用户编号", "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   from   formData   string  true    "书主人用户编号"
// @Param   to   formData   string  true      "借书人用户编号"
// @Param   bookid   formData   string  true  "书主人书架图书编号"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderadd [post]
func (this *BookorderController) Orderadd() {
	bookid  :=   this.GetString("bookid")
	from    :=   this.GetString("from")
	to      :=   this.GetString("to")
	if from =="" || bookid=="" || to == ""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	//查询书主人用户书架
	book := models.BookrackList{}
	book.Userid = from
	book.Bookid = bookid
	query:= []string{from,bookid}
	sql:= "select * from lb_users as u left join lb_bookrack as b  on u.userid= b.userid  where b.userid=? and b.bookid=?  limit 1"
	RawSeter := comm.RawSeter(sql,query)
	err := RawSeter.QueryRow(&book)
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
	newBook["isbn"] = book.Isbn
	newBook["depreciation"] = book.Depreciation
	newBook["price"] = book.Price
	newBook["describe"] = book.Describe
	newBook["state"] = book.State
	//查询书主人以及借书用户信息
	toUser := []models.Users{}
	userInfo:= []string{from,to}
	UserSql:= "select * from lb_users  where userid in(? ,?)"
	UserRawSeter := comm.RawSeter(UserSql,userInfo)
	_,uerr := UserRawSeter.QueryRows(&toUser)
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
	Users := map[string]interface{}{}
	Users["from"] = UserFrom
	Users["to"]   = UserTo
	bty,jerr:= json.Marshal(&Users)
	if jerr != nil{
		common.ErrSystem.Message = "未知错误!"
		this.renderJson(common.ErrSystem)
	}
	res,rerr:= json.Marshal(&newBook)
	if rerr != nil{
		common.ErrSystem.Message = "未知错误!"
		this.renderJson(common.ErrSystem)
	}

	NewsInfo := models.Bookorder{}
	id := models.GetID()
	NewsInfo.Orderid =  fmt.Sprintf("%d", id)
	NewsInfo.Userid_to = to
	NewsInfo.Userid_from = from
	NewsInfo.Books  = string(res)
	NewsInfo.Users  = string(bty)
	NewsInfo.Order_state = 0
	t:= time.Now().Unix()
	NewsInfo.Pushtime = t
	toRes :=  comm.Insert(&NewsInfo)
    if toRes == nil {
		common.Actionsuccess.Message ="订单添加成功!"
		common.Actionsuccess.MoreInfo = &NewsInfo
		this.renderJson(common.Actionsuccess)
	}
	common.ErrSystem.Message = "未知错误!"
	this.renderJson(common.ErrSystem)
}



// @Title    更改订单
// @Summary  更改订单
// @Description 更改订单
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   orderid   formData   string  true    "订单编号"
// @Param   order_state   formData   string  true   "消息状态状态1:同意2:拒绝,3:完成，0：借书请求"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderupdate [post]
func (this *BookorderController) Orderupdate() {
	orderid   :=   this.GetString("orderid")
	order_state  :=   this.GetString("order_state")
	if   order_state == "" || orderid ==""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	model := models.Bookorder{}
	model.Orderid = orderid
	err := comm.Read(&model)
	if err == nil {
		i64, err := strconv.ParseUint(order_state, 10, 8)
		model.Order_state = uint8(i64)
		err = comm.Update(&model)
		if err == nil {
			common.Actionsuccess.Message ="当前订单状态已修改"
			common.Actionsuccess.MoreInfo = &model
			this.renderJson(common.Actionsuccess)
		}
	}
	common.ErrSystem.Message = fmt.Sprint(err)
	this.renderJson(common.ErrSystem)
}