package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)

type Wxusers struct {
	Id                int
	Sid               int		`orm:"size(10);" form:"Sid"`
	Bin               int    	`orm:"size(64);" form:"Bin" valid:"Required"`
	Uin               int    	`orm:"size(64);" form:"Uin" valid:"Required"`
	NickName          string	`orm:"size(255);" form:"NickName" valid:"Required"`
	HeadImgUrl        string	`orm:"size(255);" form:"HeadImgUrl"`
	RemarkName        string	`orm:"size(255);" form:"RemarkName" `
	Sex               int		`orm:"size(2);"  form:"Sex" `
	Signature         string	`orm:"size(255);" form:"Signature"`
	Statues           int		`orm:"size(2);"  form:"Statues"`
	Province          string	`orm:"size(64);" form:"Province"`
	City              string	`orm:"size(64);" form:"City"`
	Alias             string	`orm:"size(64);" form:"Alias"`
	CreatedAt         int		`orm:"size(10);" form:"CreatedAt"`
	UpdatedAt         int		`orm:"size(10);" form:"UpdatedAt"`
}

func (h *Wxusers) TableName() string {
	return beego.AppConfig.String("table_wx_user")
}

func init() {
	orm.RegisterModel(new(Wxusers))
}


func checkWxuser(h *Wxusers) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&h)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//get Helper list
func GetWxuserlist(page int64, page_size int64, sort string) (helpers []orm.Params, count int64) {
	o := orm.NewOrm()
	helper := new(Wxusers)
	qs := o.QueryTable(helper)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&helpers)
	count, _ = qs.Count()
	return helpers, count
}

func AddWxuser(h *Wxusers) (int64, error) {
	if err := checkWxuser(h); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	helper := new(Wxusers)
	helper.Sid = 0
	helper.Bin  = h.Bin
	helper.Uin  = h.Uin
	helper.NickName   = h.NickName
	helper.RemarkName = h.RemarkName
	helper.CreatedAt  = h.CreatedAt
	helper.UpdatedAt  = 0
	id, err := o.Insert(helper)
	return id, err
}

func UpdateWxuser(h *Wxusers) (int64, error) {
	if err := checkWxuser(h); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	helper := make(orm.Params)
	if h.CreatedAt != 0 {
		helper["CreatedAt"] = h.CreatedAt
	}
	if h.UpdatedAt != 0 {
		helper["UpdatedAt"] = h.UpdatedAt
	}
	if len(helper) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Wxusers
	num, err := o.QueryTable(table).Filter("Uin", h.Uin).Update(helper)
	return num, err
}

func DelWxuserByName(uin int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Wxusers{Uin: uin})
	return status, err
}

func WxuserList() (helpers []orm.Params) {
	o := orm.NewOrm()
	helper := new(Wxusers)
	qs := o.QueryTable(helper)
	qs.Values(&helpers, "name", "data")
	return helpers
}

func GetWxuser(conditions string) (helper *Wxusers,err error) {
	rowSql  := "select * from  "+helper.TableName()+"  where true "+conditions+" limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowSql).QueryRow(&helper)
	return
}