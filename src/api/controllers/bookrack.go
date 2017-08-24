package controllers

import (
	comm "common/conndatabase"
	"common"
	"models"
	"fmt"
	"time"
	"strconv"
	"github.com/astaxie/beego/validation"
)

type BooksrackController struct {
	ApiController
}

// @Title    用户获取书城图书列表
// @Summary  用户获取书架图书列表
// @Description  用户获取书架图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"
// @Param   bookqid    formData   string  false   "图书唯一码"
// @Param   isbn      formData   string  false  "图书条形码"
// @Param   userid    formData   string  false   "用户编号"
// @Param   book_state    formData   string  false   "图书状态"
// @Param   gender    formData   string  false   "性别"
// @Param   age    formData   string  false   "年龄"
// @Failure 500 服务器错误!
// @router /bookracklist [post]
func (this *BooksrackController) Bookracklist() {
	length, _ := this.GetInt("length",10) //获取分页步长
	draw, _ := this.GetInt("draw",1) //获取请求次数
	var conditions string = " "
	if 	v := this.GetString("userid");v != ""{
		conditions+= " and r.userid ='"+v+"'"
	}
	if v := this.GetString("bookqid");v !="" {
		conditions+= " and r.bookqid ='"+v+"'"
	}
	if v := this.GetString("bookid");v !="" {
		conditions+= " and r.bookid ='"+v+"'"
	}
	if v := this.GetString("book_state");v !="" {
		conditions+= " and r.book_state ='"+v+"'"
	}
	if v := this.GetString("bookname");v != ""{
		conditions+= " and b.bookname like '%"+v+"%'"
	}
	if v := this.GetString("author");v !="" {
		conditions+= " and b.author like '%"+v+"%'"
	}
	if 	v := this.GetString("isbn");v !="" {
		conditions+= " and b.isbn ="+v
	}
	if 	v := this.GetString("gender");v !="" {
		conditions+= " and u.gender ="+v
	}
	if v := this.GetString("age");v !="" {
		conditions+= " and u.age ="+v
	}
	books := models.BooksrackList((draw-1)*length,length,conditions)
	Json := map[string]interface{}{"draw":draw,"data":books}
	this.renderJson(Json)
}

// @Title 添加到书架
// @Summary  添加到书架
// @Description 添加到书架
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   userid   formData   string  true    "用户编号"
// @Param   bookid   formData   string  true    "图书编号"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackadd [post]
func (this *BooksrackController) Bookrackadd() {
	userid   :=   this.GetString("userid")
	bookid   :=   this.GetString("bookid")
	if userid =="" || bookid==""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	//查询图书信息
	model := models.Books{}
	model.Bookid = bookid
	if Bookserr := comm.Read(&model);Bookserr != nil {
		common.ErrSystem.Message = "没有当前图书"
		this.renderJson(common.ErrSystem)
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
		book.Update_time = 0
		err := comm.Insert(&book)
		if err != nil {
			common.ErrSystem.Message = fmt.Sprint(err)
			this.renderJson(common.ErrSystem)
		}
	}
	common.Actionsuccess.Message ="当前图书已加入书架"
	common.Actionsuccess.MoreInfo = &book
	this.renderJson(common.Actionsuccess)
}

// @Title 扫条码添加到书架
// @Summary  扫条码添加到书架
// @Description 扫条码添加到书架
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   userid   formData   string  true  "用户编号"
// @Param   isbn     formData   string  true  "图书条码"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackaddbysn [post]
func (this *BooksrackController) Bookrackaddbysn() {
	userid  :=   this.GetString("userid")
	isbn   := this.GetString("isbn")
	if userid =="" || isbn==""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	//查询数据库是否存在该条形码
	model,err:= models.GetIbsn(isbn)
	if err != nil {
		//state 1:扫码正常添加，3:扫码没有查询到结果待补充
		var state uint8 = 1
		res, berr:= common.GetBarcodeInfo(isbn)
		//查询失败
		if berr != nil{
			common.ErrSystem.Message = "网络繁忙!"
			this.renderJson(common.ErrSystem)
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
				common.ErrSystem.Message = fmt.Sprint(err.Message)
				this.renderJson(common.ErrSystem)
			}
		}
		err = comm.Insert(model)
		if err != nil{
			common.ErrSystem.Message = "没有当前图书信息!"
			this.renderJson(common.ErrSystem)
		}
	}
	book := models.Bookrack{}
	book.Userid  = userid
	book.Bookid = model.Bookid
	book.Book_state = "1"
	book.Is_borrow  = "1"
	_,err = models.GetBookByUidAndBookId(userid,book.Bookid)
	if err != nil {
		//加入用户书架
		id := models.GetID()
		book.Bookqid = fmt.Sprintf("%d",id)
		t:=time.Now().Unix()
		book.Create_time = t
		book.Update_time = t
		err := comm.Insert(&book)
		if err != nil {
			common.ErrSystem.Message = fmt.Sprint(err)
			this.renderJson(common.ErrSystem)
		}
	}
	common.Actionsuccess.Message ="当前图书已加入书架"
	common.Actionsuccess.MoreInfo = &book
	this.renderJson(common.Actionsuccess)
}


// @Title    用户更改书架图书状态
// @Summary  用户更改书架图书状态
// @Description 用户更改书架图书状态
// @Success 200  { <br/>"bookqid": "图书唯一编号",<br/> "userid": "用户id",<br/> "bookid": "图书编号", <br/>"book_state": "状态1:上架;2:下架;3:待补充",<br/> "is_borrow": "状态1:可借阅;2:已借出;3:不可借",<br/> "create_time": "上架时间",<br/>"update_time":"信息修改时间"<br/> }
// @Param   bookqid   formData   string  true    "图书唯一编号"
// @Param   bookstate   formData   string  true   "书架图书状态1:上架;2:下架;3:待补充"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackupdate [post]
func (this *BooksrackController) Bookrackupdate() {
	bookqid   :=   this.GetString("bookqid")
	bookstate    :=   this.GetString("bookstate")
	if bookqid =="" || bookstate == ""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	book,err:= models.GetBookById(bookqid)
	if err == nil {
		book.Bookqid = bookqid
		book.Book_state = bookstate
		if book.Is_borrow == "2"{
			common.ErrSystem.Message = "图书已借出,不能操作状态"
			this.renderJson(common.ErrSystem)
		}
		if  bookstate !="1"{
			 book.Is_borrow =  "2"
		}
		book.Update_time = time.Now().Unix()
		if err :=models.UpdateBookRackById(book);err == nil {
			common.Actionsuccess.Message ="当前图书状态修改成功!"
			common.Actionsuccess.MoreInfo = &book
			this.renderJson(common.Actionsuccess)
		}
	}
	common.ErrSystem.Message = fmt.Sprint(err)
	this.renderJson(common.ErrSystem)
}