package models

import(
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
)

type Booknews struct {
	Newid       string  `json:"newid" valid:"Required" orm:"pk;size(20);column(newid);"`
	Userid_from	string	`json:"userid_from" valid:"Required"`
	Userid_to	string	`json:"userid_to" valid:"Required"`
	Bookqid	    string	`json:"bookqid" valid:"Required"`
	Books       string  `json:"books"   valid:"Required"`
	User_from   string	`json:"user_from" valid:"Required"`
	User_to     string	`json:"user_to"`
	Order_type  uint8   `json:"order_state" valid:"Required;Range(1,2)"`
	Order_state uint8   `json:"order_state" valid:"Required;Range(0,1,2,3)"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}

type BooknewsList struct {
	Newid       string  `json:"newid" orm:"pk;size(20);column(newid);"`
	Userid_from	string	`json:"userid_from"`
	Userid_to	string	`json:"userid_to"`
	Bookqid	    string	`json:"bookqid"`
	Books       string  `json:"books"`
	User_from   string	`json:"user_from"`
	User_to     string	`json:"user_to"`
	Order_state uint8   `json:"order_state"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}


func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Booknews))
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

