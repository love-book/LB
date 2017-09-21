package controllers

import (
	comm "common/conndatabase"
	"models"
	"time"
	"encoding/json"
	"strconv"
	"math"
)

type BookorderController struct {
	ApiController
}

// @Title        交换记录列表
// @Summary      交换记录列表
// @Description  交换记录列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token   header     string  true  "token"
// @Param	body	body 	 models.OrderlistForm  true   "{ <br/>"length":"获取分页步长",<br/>"length":"获取分页步长", <br/>"order_state":"状态0:待确认;1:已交换2:拒绝交换"<br/> }"
// @Failure 500 服务器错误!
// @router /orderList [post]
func (this *BookorderController) Orderlist() {
	var ob *models.OrderlistForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	length := ob.Length
	draw := ob.Draw
	var conditions string = ""
	if ob.Isbn !="" {
		conditions+= " and JSON_EXTRACT(books,'$.isbn') = '"+ob.Isbn+"'"
	}
	if  ob.OrderState != "" {
		conditions+= " and order_state = '"+ob.OrderState+"'"
	}
	conditions+= " and userid_from = '"+this.Userid+"'"
	books,count := models.BookOrderListDatback((draw-1)*length,length,conditions)
	var resPonse []interface{}
	if len(books) >= 1{
		for _,val := range books{
			book  := map[string]interface{}{}
			JsonBook  := map[string]interface{}{}
			JsonFrom  := map[string]interface{}{}
			JsonTo  := map[string]interface{}{}
			book["orderid"]     =  val.Orderid
			book["userid_from"] =  val.Userid_from
			book["userid_to"]   =  val.Userid_to
			book["bookqid"]     =  val.Bookqid
			book["order_type"]  =  val.Order_type
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
	}
	pageTotal:= math.Ceil(float64(count)/float64(length))
	json := map[string]interface{}{"pageTotal":pageTotal,"draw":draw,"data":&resPonse}
	this.Rsp(true, "获取成功!",&json)
}

// @Title    更改交易记录状态
// @Summary  更改交易记录状态
// @Description 更改交易记录状态
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token     header     string  true  "token"
// @Param	 body	 body 	     models.OrderupdateForm  true   "{ <br/>"orderid":"订单编号", <br/>"order_state":"状态0:待确认;1:已交换2:拒绝交换"<br/> }"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderupdate [post]
func (this *BookorderController) Orderupdate() {
	var ob *models.OrderupdateForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	orderid     :=  ob.Orderid
	order_state :=  ob.OrderState
	if order_state == "" || orderid ==""{
		this.Rsp(false, "参数错误!","")
	}
	//查询订单
	model,err := models.BookOrderInfo(" and orderid="+orderid+" and userid_from="+this.Userid)
	if err!=nil{
		this.Rsp(false, "不存在订单!","")
	}
	i64, err := strconv.ParseUint(order_state, 10, 8)
	model.Order_state =  int64(i64)
	model.Update_time = time.Now().Unix()
	err = comm.Update(&model)
	if err == nil {
		this.Rsp(true, "当前订单状态已修改!",&model)
	}
	this.Rsp(false, "修改失败!","")
}


// @Title    删除交易记录
// @Summary  删除交易记录
// @Description 删除交易记录
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   token    header     string  true  "token"
// @Param	 body	 body 	     models.OrderdeleteForm  true   "{ <br/> "orderid":"订单编号" <br/> }"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /orderdelete [post]
func (this *BookorderController) Orderdelete() {
	var ob *models.OrderdeleteForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	orderid     :=  ob.Orderid
	if orderid == ""{
		this.Rsp(false, "参数错误!","")
	}
	//查询订单
	model,err := models.BookOrderInfo(" and orderid="+orderid+" and userid_from="+this.Userid)
	if err!=nil{
		this.Rsp(false, "不存在订单!","")
	}
	err = comm.Delete(&model)
	if err == nil {
		this.Rsp(true, "当前订单已删除!",&model)
	}
	this.Rsp(false, "删除失败!","")
}