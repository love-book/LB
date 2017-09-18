package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
)

type Bookrack struct {
	Bookqid      string  `json:"bookqid" valid:"Required" orm:"pk;size(20);column(bookqid);"`
	Userid       string  `json:"userid" valid:"Required"`
	Bookid       string  `json:"bookid" valid:"Required"`
	Book_state   string  `json:"book_state"`
	Is_borrow    string  `json:"is_borrow"`
	Create_time  int64   `json:"create_time"`
	Update_time  int64   `json:"update_time"`
}


type BookrackList struct {
	Bookqid      string  `json:"bookqid" orm:"pk;size(20);column(bookqid);"`
	Userid       string  `json:"userid"`
	Openid       string  `json:"openid"`
	Bookid       string  `json:"bookid"`
	Book_state   string  `json:"book_state"`
	Is_borrow    string  `json:"is_borrow"`
	Create_time  int64   `json:"create_time"`
	Update_time  int64   `json:"update_time"`
	Bookname     string  `json:"bookname"`
	Author       string	 `json:"auhtor""`
	Imageurl     string	 `json:"imageurl"`
	Imagehead    string	 `json:"imagehead"`
	Imageback    string	 `json:"imageback"`
	Isbn         string	 `json:"isbn"`
	Depreciation uint8	 `json:"depreciation"`
	Price        uint16	 `json:"price"`
	Describe     string	 `json:"describe"`
	State        uint8	 `json:"state" `
	Users
	Radius     	 string	 `json:"radius"`
}


type BookExist struct {
	Userid       string   `json:"userid"`
	Bookid       string   `json:"bookid"`
}


type BookState struct {
	Bookqid       string   `json:"bookqid"`
	Book_state     string   `json:"book_state"`
}

func (b *Bookrack) TableName() string {
	return beego.AppConfig.String("table_bookrack")
}

func init()  {
	orm.RegisterModel(new(Bookrack))
}

func  BooksrackList(start int,length int,conditions string) (books []*BookrackList){
    var  rowsSql  = "select r.*,b.*,u.openid,u.nickname,u.imgurl,u.gender,u.age from  lb_bookrack as r left join lb_books  as b on r.bookid = b.bookid  left join lb_users as u on u.userid = r.userid  where true "+conditions+"  order by r.create_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRows(&books)
	return  books
}

func  MyBooksrackList(start int,length int,conditions string) (books []*BookrackList,totalItem int){
	var  countSql = "select count(*) from  lb_bookrack as r inner join lb_books  as b on r.bookid = b.bookid  inner join lb_users as u on u.userid = r.userid  where true "+conditions
	var  rowsSql  = "select r.*,b.*,u.* from  lb_bookrack as r inner  join lb_books  as b on r.bookid = b.bookid  inner  join lb_users as u on u.userid = r.userid  where true "+conditions+"  order by r.create_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&totalItem) //获取总条数
	o.Raw(rowsSql).QueryRows(&books)
	return  books,totalItem
}

func  MyBooksrackInfo(conditions string) (book *BookrackList){
	var  rowsSql  = "select r.*,b.*,u.* from  lb_bookrack as r inner  join lb_books  as b on r.bookid = b.bookid  inner  join lb_users as u on u.userid = r.userid  where true "+conditions+"  limit 1"
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRow(&book)
	return  book
}

//根据id查询
func GetBookById(bookqid string) (v *Bookrack, err error) {
	o := orm.NewOrm()
	v = &Bookrack{Bookqid: bookqid}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//查询书主人用户书架
func GetUserBookRack(uid string,bookqid string)(b *BookrackList,err error){
	query:= []string{uid,bookqid}
	sql:= "select * from lb_bookrack as r left join lb_books as b on r.bookid = b.bookid  where r.userid=? and r.bookqid=?  limit 1"
	RawSeter:=orm.NewOrm().Raw(sql,query)
	if err = RawSeter.QueryRow(&b);err==nil{
		return b,nil
	}
	return nil,err
}
//根据用户id和书本id查询
func GetBookByUidAndBookId(uid string,bookid string)(b *Bookrack,err error){
	query:= []string{uid,bookid}
	sql:= "select * from lb_bookrack where userid=? and bookid=?"
	RawSeter:=orm.NewOrm().Raw(sql,query)
	if err = RawSeter.QueryRow(&b);err==nil{
		return b,nil
	}
	return nil,err
}

// 修改信息
func UpdateBookRackById(m *Bookrack) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.Update(m); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return
}

