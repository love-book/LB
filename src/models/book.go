package models

import(
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
	"fmt"
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

func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Books))
}

func (this *Books) GetBookinfo (bookid int64) Books  {

	return *this
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