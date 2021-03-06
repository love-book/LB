package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
	"strconv"
)

type Bookorder struct {
	Orderid     string `json:"newid" valid:"Required" orm:"pk;size(20);column(orderid);"`
	Userid_from	string	`json:"userid_from" valid:"Required"`
	Userid_to	string	`json:"userid_to"  valid:"Required"`
	Book_from   string  `json:"book_from"  valid:"Required"`
	Book_to     string  `json:"book_to"    valid:"Required"`
	User_from   string	`json:"user_from"  valid:"Required"`
	User_to     string	`json:"user_to"`
	Order_type  int64   `json:"order_state" valid:"Required;Range(1,2)"`
	Order_state int64   `json:"order_state" valid:"Required;Range(1,2,3)"`
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
	var  rowsSql  = "select * from  lb_bookorder  where true "+conditions+"  order by update_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(rowsSql).QueryRows(&order)
	return  order
}

func  BookOrderListDatback(start int,length int,conditions string) (order []*Bookorder,total int){
	var  countSql = "select count(*) from  lb_bookorder  where true "+conditions
	var  rowsSql  = "select * from  lb_bookorder  where true "+conditions+"  order by update_time desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&total) //获取总条数
	o.Raw(rowsSql).QueryRows(&order)
	return  order,total
}

func  BookOrderInfo(conditions string) (order *Bookorder,err error){
	var  rowsSql  = "select * from  lb_bookorder  where true "+conditions+"  order by update_time desc  limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowsSql).QueryRow(&order)
	return
}