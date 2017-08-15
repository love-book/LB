package controllers

import (
	comm "common/conndatabase"
	"common"
	"models"
	"fmt"
	"time"
)

type BooksrackController struct {
	ApiController
}

// @Title 用户获取书架图书列表
// @Summary  用户获取书架图书列表
// @Description  用户获取书架图书列表
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookname  formData   string  false   "书名"
// @Param   author    formData   string  false   "作者"
// @Param   bookid    formData   string  false   "图书编号"
// @Param   isbn      formData   string  false  "图书条形码"
// @Param   userid    formData   string  false   "用户编号"
// @Param   bookstate    formData   string  false   "图书状态"
// @Param   gender    formData   string  false   "性别"
// @Param   age    formData   string  false   "年龄"
// @Failure 500 服务器错误!
// @router /bookracklist [post]
func (this *BooksrackController) Bookracklist() {
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var conditions string = " "
	id := this.GetString("userid")
	if id != ""{
		conditions+= " and r.userid ='"+id+"'"
	}
	bookid := this.GetString("bookid")
	if bookid !="" {
		conditions+= " and r.bookid ='"+bookid+"'"
	}
	bookstate := this.GetString("bookstate")
	if bookstate !="" {
		conditions+= " and r.bookstate ='"+bookstate+"'"
	}
	bookname := this.GetString("bookname")
	if bookname != ""{
		conditions+= " and b.bookname like '%"+bookname+"%'"
	}
	author := this.GetString("author")
	if author !="" {
		conditions+= " and b.author like '%"+author+"%'"
	}
	isbn := this.GetString("isbn")
	if isbn !="" {
		conditions+= " and b.isbn ="+isbn
	}
	gender := this.GetString("gender")
	if gender !="" {
		conditions+= " and u.gender ="+gender
	}
	age := this.GetString("age")
	if age !="" {
		conditions+= " and u.age ="+age
	}
	var start int = 0
	if draw  > 0 {
		start = (draw-1)*length
	}
	conditions += "  order by r.create_time desc"
	books,totalItem := models.BooksrackList(start,length,conditions,"r.*,b.*,u.nickname,u.imgurl,u.gender,u.age")
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":books}
	this.renderJson(Json)
}



// @Title 用户获取书架图书详情
// @Summary  用户获取书架图书详情
// @Description  用户获取书架图书详情
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "bookstate": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   length    formData   string  false      "获取分页步长"
// @Param   draw      formData   string  false      "当前页"
// @Param   bookid    formData   string  true   "图书编号"
// @Param   userid    formData   string  true   "用户编号"
// @Param   bookstate    formData   string  false   "图书状态"
// @Failure 500 服务器错误!
// @router /bookrackinfo [post]
func (this *BooksrackController) Bookrackinfo() {
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var conditions string = " "
	id := this.GetString("userid")
	if id != ""{
		conditions+= " and r.userid ='"+id+"'"
	}
	bookid := this.GetString("bookid")
	if bookid !="" {
		conditions+= " and r.bookid ='"+bookid+"'"
	}
	bookstate := this.GetString("bookstate")
	if bookstate !="" {
		conditions+= " and r.bookstate ='"+bookstate+"'"
	}
	var start int = 0
	if draw  > 0 {
		start = (draw-1)*length
	}
	conditions += "  order by r.create_time desc"
	books,totalItem:=models.BooksrackInfo(start,length,conditions,"r.*,b.*,u.*")
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
	book.Bookstate = "1"
	query:= []string{userid,bookid}
	sql:= "select userid,bookid from lb_bookrack where userid=? and bookid=?"
	RawSeter := comm.RawSeter(sql,query)
	err := RawSeter.QueryRow(&book)
	if err != nil {
		book.Create_time = time.Now().Unix()
		book.Update_time = 0
		err = comm.Insert(&book)
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
// @Success 200  {<br/> "bookid": "图书编号",<br/> "bookname": "书名",<br/> "author": "作者",<br/> "imgurl": "图书封面图", <br/>"imgheadurl": "图书正面图",<br/> "imgbackurl": "图书背面图",<br/> "barcode":"条形码",<br/> "depreciation":"",<br/> "price":"标价", <br/>"describe": "图书简介",<br/> "state": "状态",<br/> "created_at": "上架时间",<br/>"updated_at":"信息修改时间"<br/> }
// @Param   userid   formData   string  true    "用户编号"
// @Param   bookid   formData   string  false    "图书编号"
// @Param   bookstate   formData   string  true   "图书状态"
// @Failure 100 错误提示信息!
// @Failure 500 服务器错误!
// @router /bookrackupdate [post]
func (this *BooksrackController) Bookrackupdate() {
	userid   :=   this.GetString("userid")
	bookid   :=   this.GetString("bookid")
	bookstate    :=   this.GetString("bookstate")
	if userid =="" || bookstate == ""{
		common.ErrSystem.Message = "参数错误!"
		this.renderJson(common.ErrSystem)
	}
	book := models.Bookrack{}
	book.Userid = userid

	book.Update_time = time.Now().Unix()
	book.Bookstate = bookstate
	var query []string
	var where string
	if bookid !="" {
		where = " userid=? and bookid=? "
		book.Bookid = bookid
		query = []string{userid,bookid}
	}else{
		where = " userid=? "
		query = []string{userid}
	}
	sql:= "select userid,bookid from lb_bookrack where "+where
	RawSeter := comm.RawSeter(sql,query)
	err := RawSeter.QueryRow(&book)
	if err != nil {
		book.Create_time = time.Now().Unix()
		book.Update_time = 0
		err = comm.Update(&book)
		if err != nil {
			common.ErrSystem.Message = fmt.Sprint(err)
			this.renderJson(common.ErrSystem)
		}
	}
	common.Actionsuccess.Message ="当前图书状态修改成功!"
	common.Actionsuccess.MoreInfo = &book
	this.renderJson(common.Actionsuccess)
}