package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
)

type Bookorder struct {
	Orderid      string `json:"newid" valid:"Required" orm:"pk;size(20);column(orderid);"`
	Userid_from	string	`json:"userid_from" valid:"Required"`
	Userid_to	string	`json:"userid_to" valid:"Required"`
	Bookqid	    string	`json:"bookqid" valid:"Required"`
	Books       string  `json:"books"  valid:"Required"`
	User_from   string	`json:"user_from" valid:"Required"`
	User_to     string	`json:"user_to"`
	Order_state uint8   `json:"order_state" valid:"Required;Range(1,2,3)"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}

func (b *Bookorder) TableName() string {
	return beego.AppConfig.String("table_bookorder")
}

func init()  {
	orm.RegisterModel(new(Bookorder))
}

func  BookOrderListData(start int,length int,conditions string) (order []*Bookorder){
	var  rowsSql  = "select * from  lb_bookorder  where true "+conditions+"  order by pushtime desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRows(&order)
	return  order
}

func  BookOrderListDatback(start int,length int,conditions string) (order []*Bookorder,total int){
	var  countSql = "select count(*) from  lb_bookorder  where true "+conditions
	var  rowsSql  = "select * from  lb_bookorder  where true "+conditions+"  order by pushtime desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&total) //获取总条数
	o.Raw(rowsSql).QueryRows(&order)
	return  order,total
}