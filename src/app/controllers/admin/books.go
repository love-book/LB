package admin

import (
	comm "app/common/conndatabase"
	"github.com/astaxie/beego/validation"
	"app/common"
	"app/models"
	"fmt"
)
//  后台书籍 API
type BooksController struct {
	baseController
}


func (this *BooksController) Prepare() {
	//关闭xsrf校验
	this.EnableXSRF = false
	this.EnableRender =false
}

// @Title 获取图书列表
// @Summary  获取图书列表
// @Description 获取图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   start     formData   int  false      "获取起始位置"
// @Param   length    formData   int  false      "获取分页步长"
// @Param   draw      formData   int  false      "获取请求次数"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"
// @Param   isbn    formData   string  false  "图书条形码"
// @Failure 500 服务器错误!
// @router /booklist [post]
func (this *BooksController) BookList() {
	start, _ := this.GetInt("start") //获取起始位置
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt32("draw") //获取请求次数
	var conditions string = " "
	id := this.GetString("bookid")
	if id != ""{
		conditions+= " and bookid ='"+id+"'"
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
	/*
	starttime := this.GetString("starttime")
	if starttime !="" {
		tm1, _ := time.Parse("01/02/2006", starttime)
		endtime := this.GetString("endtime")
		if endtime == ""{
			endtime = fmt.Sprintf("%d",time.Now().Unix())
		}else{
			tm2, _ := time.Parse("01/02/2006", endtime)
			endtime = fmt.Sprintf("%d",tm2)
		}
		starttime = fmt.Sprintf("%d",tm1)
		conditions+= " and created_at  bettwen "+starttime+" and "+endtime
	}*/
	var books []models.Books
	conditions += "  order by bookid desc"
	var  TableName = "lb_books"
	totalItem, res :=models.GetPagesInfo(TableName,start,length,conditions)
	res.QueryRows(&books)
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":books}
	this.renderJson(Json)
}


// @Title 添加书籍
// @Summary  添加书籍
// @Description 添加书籍
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   bookname   formData   string  true    "书名"
// @Param   author     formData   string  false   "作者"
// @Param   imgurl  formData   string  false    "图书封面图"
// @Param   imghead    formData   string  false     "图书正面图"
// @Param   imgback       formData   string  false    "图书背面图"
// @Param   isbn      formData   string  false    "条形码"
// @Param   describe      formData   string  false    "图书简介"
// @Param   price    formData   string  false    "标价"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookadd [post]
func (this *BooksController) Bookadd() {
	model := models.Books{}
	Uid := models.GetID()
	model.Bookid   =  fmt.Sprintf("%d", Uid)
	model.Bookname = this.GetString("bookname")
	model.Author  = this.GetString("author")
	model.Imagehead = this.GetString("imageback")
	model.Imageback = this.GetString("imageback")
	model.Imageurl  = this.GetString("imgurl")
	model.Isbn = this.GetString("isbn")
	model.Describe = this.GetString("describe")
	model.Price,_ = this.GetUint16("price")
	model.Isbn  =  this.GetString("isbn")
	model.Depreciation = 0
	model.State = 1
	valid := validation.Validation{}
	valid.Required(model.Bookid,  "bookid").Message("书籍编号不能为空！")
	valid.Required(model.Bookname, "bookname").Message("书名不能为空！")
	valid.Required(model.Author,"author").Message("作者不能为空！")
	valid.Required(model.Author,"isbn").Message("条形码不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			common.ErrSystem.Message = fmt.Sprint(err.Message)
			this.renderJson(common.ErrSystem)
		}
	}
	err := comm.Insert(&model)
	if err != nil {
		common.ErrSystem.Message = fmt.Sprint(err)
		this.renderJson(common.ErrSystem)
	}
	common.Actionsuccess.MoreInfo =  &model
	this.renderJson(common.Actionsuccess)
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
// @Param   state    formData   string  false    "状态1:上架;2:下架"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookupdate [post]
func (this *BooksController) Bookupdate(){
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
		isbn := this.GetString("Barcode")
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