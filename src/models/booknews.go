package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
)

type Booknews struct {
	Newid       string  `json:"newid" valid:"Required" orm:"pk;size(20);column(newid);"`
	Userid_from	string	`json:"userid_from" valid:"Required"`
	Userid_to	string	`json:"userid_to" valid:"Required"`
	Book_from   string  `json:"book_from"`
	Book_to     string  `json:"book_to"`
	User_from   string	`json:"user_from" valid:"Required"`
	User_to     string	`json:"user_to"`
	Order_type  int64   `json:"order_state" valid:"Required;Range(1,2)"`
	Order_state int64   `json:"order_state" valid:"Required;Range(0,1,2,3)"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}

type BooknewsList struct {
	Newid       string  `json:"newid" orm:"pk;size(20);column(newid);"`
	Userid_from	string	`json:"userid_from"`
	Userid_to	string	`json:"userid_to"`
	Book_from   string  `json:"book_from"`
	Book_to     string  `json:"book_to"`
	User_from   string	`json:"user_from"`
	User_to     string	`json:"user_to"`
	Order_state int64   `json:"order_state"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}

func (b *Booknews) TableName() string {
	return beego.AppConfig.String("table_booknews")
}

func init()  {
	orm.RegisterModel(new(Booknews))
}

func  BooknewsListData(start int,length int,conditions string) (books []*Booknews){
    var  rowsSql  = "select * from  lb_booknews  where true "+conditions+"  order by update_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRows(&books)
	return  books
}

func  BooknewsListDataBack(start int,length int,conditions string) (books []*Booknews,total int){
	var  countSql = "select count(*) from  lb_booknews  where true "+conditions
	var  rowsSql  = "select * from  lb_booknews  where true "+conditions+"  order by update_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&total) //获取总条数
	o.Raw(rowsSql).QueryRows(&books)
	return  books,total
}

func  BooknewsInfo(conditions string) (b *Booknews,err error){
	var  rowsSql  = "select * from  lb_booknews  where true "+conditions+"  limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowsSql).QueryRow(&b)
	return
}
