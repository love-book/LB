package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type Menu struct {
	Id     int32    `orm:"size(11);form:"Id"`
	Name   string   `orm:"size(128);"form:"Data" valid:"Required"`
	Parent int32    `orm:"size(11);null" form:"Parent"`
	Route  string   `orm:"size(255);null" form:"Route"`
	Order  int32    `orm:"size(11);;null" form:"Order"`
	Data   string   `orm:"size(255);null" form:"Data"`
}

func (r *Menu) TableName() string {
	return beego.AppConfig.String("rbac_menu_table")
}

func init() {
	orm.RegisterModel(new(Menu))
}


func checkMenu(m *Menu) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&m)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

