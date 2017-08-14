package models

import (
	"github.com/astaxie/beego/orm"
)

type Bookrack struct {
	Userid       string  `json:"userid" valid:"Required" orm:"pk;size(20);column(userid);"`
	Books        string  `json:"books" valid:"Required"`
	Update_time int64  `json:"update_time"`
}


func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Bookrack))
}