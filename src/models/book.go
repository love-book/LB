package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
	"fmt"
	"errors"
)

type Books struct {
	Bookid       string `json:"bookid" valid:"Required" orm:"pk;size(20);column(bookid);"`
	Bookname     string `json:"bookname" valid:"Required"`
	Author       string	`json:"auhtor" valid:"Required"`
	Imageurl     string	`json:"imageurl"`
	Imagehead    string	`json:"imagehead"`
	Imageback    string	`json:"imageback"`
	Isbn         string	`json:"isbn" valid:"Required"`
	Depreciation uint8	`json:"depreciation"`
	Price        uint16	`json:"price" valid:"Numeric"`
	Describe     string	`json:"describe"`
	Userid		 string	`json:"userid"`
	State        uint8	`json:"state" valid:"Required;Range(0, 1)"`
}

func (b *Books) TableName() string {
	return beego.AppConfig.String("table_books")
}

func init() {
	orm.RegisterModel(new(Books))
}


//获取图书列表
func  GetBookList(start int,length int,conditions string) (books []*Books){
	var  rowsSql  = "select *  from  lb_books where true "+conditions+"  order by bookid desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRows(&books)
	return  books
}

//获取图书列表后端
func  GetBookListBack(start int,length int,conditions string) (books []*Books,totalItem int){
	var  countSql = "select count(*) from  lb_books  where true "+conditions
	var  rowsSql  = "select *  from  lb_books where true "+conditions+"  order by bookid desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&totalItem) //获取总条数
	o.Raw(rowsSql).QueryRows(&books)
	return  books,totalItem
}

//查询数据库是否存在该条形码
func  GetIbsn(isbn string)(b *Books,err error){
	args := []string{isbn}
	sql := "select * from lb_books where isbn=? limit 1"
	RawSeter := orm.NewOrm().Raw(sql,args)
	if err = RawSeter.QueryRow(&b);err == nil{
        return b,nil
	}
	return nil, err
}
//修改图书信息
func UpdateBookById(m *Books) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.Update(m); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return
}


//获取图书
func  GetBookInfo(conditions string) (book *Books,err error){
	var  rowsSql  = "select *  from  lb_books where true "+conditions+"  order by bookid desc  limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowsSql).QueryRow(&book)
	return  book,err
}

//获取收藏列表
func  GetConcernList(start int,length int,conditions string) (concerns []*Concern,totalItem int){
	var  countSql = "select count(*) from  lb_concern  where true "+conditions
	var  rowsSql  = "select *  from  lb_concern where true "+conditions+"  order by created_at desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&totalItem) //获取总条数
	o.Raw(rowsSql).QueryRows(&concerns)
	return  concerns,totalItem
}


func GetBook(id string) (b *Books, err error) {
	o := orm.NewOrm()
	v := Books{Bookid: id}
	if err = o.Read(&v); err == nil {
		return b, nil
	}
	return nil, errors.New("Book not exists")
}