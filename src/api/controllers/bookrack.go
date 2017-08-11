package controllers

import (
	comm "common/conndatabase"
	"common"
	"models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type BooksrackController struct {
	ApiController
}

// @Title 用户获取图书列表
// @Summary  用户获取图书列表
// @Description  用户获取图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"
// @Param   isbn      formData   string  false  "图书条形码"
// @Param   userid    formData   string  false   "用户编号"
// @Failure 500 服务器错误!
// @router /booklist [post]
func (this *BooksrackController) BookrackList() {
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var conditions string = " "
	id := this.GetString("userid")
	if id != ""{
		conditions+= " and userid ='"+id+"'"
	}
	bookname := this.GetString("bookname")
	if bookname != ""{
		conditions+= " and bookname ="+bookname
	}
	author := this.GetString("author")
	if author !="" {
		conditions+= " and author = "+author
	}
	isbn := this.GetString("isbn")
	if isbn !="" {
		conditions+= " and isbn = "+isbn
	}
	state := this.GetString("state")
	if state !="" {
		conditions+= " and state = "+state
	}
	var books []models.Books
	conditions += "  order by bookid desc"
	var  TableName = "lb_books"
	totalItem, res :=models.GetPagesInfo(TableName,0,length,conditions)
	res.QueryRows(&books)
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":books}
	this.renderJson(Json)
}



// @Title 添加到书架
// @Summary  添加到书架
// @Description 添加到书架
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   userid   formData   string  true    "用户编号"
// @Param   bookid   formData   string  true    "图书编号"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookadd [post]
func (this *BooksrackController) Bookadd() {
	userid   :=   this.GetString("userid")
	bookid   :=   this.GetString("bookid")
	if userid =="" || bookid==""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	//查询图书信息
	model := models.Books{}
	model.Bookid = bookid
	if err := comm.Read(&model);err != nil {
		common.ErrSystem.Message = "没有当前图书"
		this.renderJson(common.ErrSystem)
	}
	//加入用户书架
	book := models.Bookrack{}
	book.Userid = userid
	err := comm.Read(&book)
	if err != nil {
		//没有书架
	}else{

	}
}

// @Title 修改书籍
// @Description 修改书籍
// @Summary  修改书籍
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   bookname   formData   string  true    "书名"
// @Param   author     formData   string  false   "作者"
// @Param   imageurl  formData   string  false    "图书封面图"
// @Param   imagehead    formData   string  false     "图书正面图"
// @Param   imageback       formData   string  false    "图书背面图"
// @Param   isbn      formData   string  false    "条形码"
// @Param   describe      formData   string  false    "图书简介"
// @Param   price    formData   string  false    "标价"
// @Param   state    formData   string  false    "状态1:上架;2:下架;3:待补充"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookupdate [post]
func (this *BooksrackController) Bookupdate(){
	model := models.Books{}
	model.Bookid = this.GetString("bookid")
	valid := validation.Validation{}
	valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			common.ErrSystem.Message = fmt.Sprint(err.Message)
			this.renderJson(common.ErrSystem)
		}
	}
	if err := comm.Read(&model);err == nil {
		Bookname := this.GetString("bookname")
		if Bookname != ""{
			model.Bookname = Bookname
		}
		Author := this.GetString("author")
		if Author != ""{
			model.Author = Author
		}
		imagehead := this.GetString("imagehead")
		if imagehead != ""{
			model.Imagehead = imagehead
		}
		imageback := this.GetString("imageback")
		if imageback != ""{
			model.Imageback = imageback
		}
		Imgurl := this.GetString("imgurl")
		if Imgurl != ""{
			model.Imageurl = Imgurl
		}
		isbn := this.GetString("isbn")
		if isbn != ""{
			model.Isbn = isbn
		}
		var Price ,_ = this.GetUint16("price")
		if Price != 0{
			model.Price = Price
		}
		Describe := this.GetString("describe")
		if Describe != ""{
			model.Describe = Describe
		}
		State ,_:= this.GetUint8("state")
		if State != 0{
			model.State = State
		}
		if update:= comm.Update(&model);update ==nil{
			common.Actionsuccess.MoreInfo =  &model
			this.renderJson(common.Actionsuccess)
		}else{
			common.ErrSystem.Message = "修改失败!"
		}
	}else{
		common.ErrSystem.Message = "没有当前记录"
	}
	this.renderJson(common.ErrSystem)
}