package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type User struct {
	Id    		int32   `orm:"size(11);"  form:"Id"`
	Username   	string  `orm:"size(30);"  form:"Username" valid:"Required"`
	Password   	string  `orm:"size(64);"  form:"Password" valid:"Required"`
	Avatar  	string  `orm:"size(255);null" form:"Avatar"`
	Email  		string  `orm:"size(255);null"  form:"Email"`
	CreatedAt  	int32 	`orm:"size(11);" form:"CreatedAt" valid:"Required"`
	UpdatedAt  	int32 	`orm:"size(11);" form:"UpdatedAt" valid:"Required"`
}

func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}

func init() {
	orm.RegisterModel(new(User))
}


func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}
