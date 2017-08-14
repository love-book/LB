package models

import(
	//"github.com/astaxie/beego/validation"
	//"log"
	"github.com/astaxie/beego/orm"
	_"common/conndatabase"
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
	State        uint8	`json:"state" valid:"Required;Range(0, 1)"`
}

func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Books))
}

func (this *Books) GetBookinfo (bookid int64) Books  {

	return *this
}
