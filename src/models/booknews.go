package models

import(
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
)

type Booknews struct {
	Newid        string `json:"newid" valid:"Required" orm:"pk;size(20);column(newid);"`
	Userid		 string	`json:"userid" valid:"Required"`
	From_to      string	`json:"from_to" valid:"Required"`
	Books        string `json:"books"  valid:"Required"`
	Users        string	`json:"users" valid:"Required"`
	Order_state  uint8  `json:"order_state" valid:"Required;Range(0, 1)"`
	Pushtime     int64	`json:"pushtime"`
}

func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Booknews))
}

