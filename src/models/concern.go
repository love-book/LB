package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	"time"
)

type Concern struct {
	Concernid         string	`orm:"pk;size(20);column(concernid);" json:"concernid" form:"concernid" valid:"Required"`
	UseridTo          string    `orm:"size(20);"  json:"userid_to"  form:"userid_to"`
	UseridFrom        string    `orm:"size(20);"  json:"userid_from" form:"userid_from" valid:"Required"`
	Books             string	`orm:"size(255);" json:"books" form:"books"`
	ConcernType       string	`orm:"size(255);" json:"concern_type" form:"concern_type" valid:"Required";Range(1,2)"`
	CreatedAt         int64		`orm:"size(10);"  json:"created_at" form:"created_at"`
}

func (c *Concern) TableName() string {
	return beego.AppConfig.String("table_concern")
}

func init() {
	orm.RegisterModel(new(Concern))
}


func CheckConcern(c *Concern) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&c)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}



func GetConcernlist(page int64, page_size int64, sort string) (concerns []orm.Params, count int64) {
	o := orm.NewOrm()
	c := new(Concern)
	qs := o.QueryTable(c)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&concerns)
	count, _ = qs.Count()
	return concerns, count
}

func AddConcern(c *Concern) (int64, error) {
	if err := CheckConcern(c); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	concern := new(Concern)
	concern.Concernid  = c.Concernid
	concern.UseridTo   = c.UseridTo
	concern.UseridFrom = c.UseridFrom
	if c.Books!=""{
		concern.Books      = c.Books
	}else{
		concern.Books      = "{}"
	}
	concern.ConcernType= c.ConcernType
	concern.CreatedAt  = time.Now().Unix()
	id, err := o.Insert(concern)
	return id, err
}

func UpdateConcern(c *Concern) (int64, error) {
	if err := CheckConcern(c); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	concern := make(orm.Params)
	if c.CreatedAt != 0 {
		concern["CreatedAt"] = c.CreatedAt
	}
	if len(concern) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Concern
	num, err := o.QueryTable(table).Filter("Concernid", c.Concernid).Update(concern)
	return num, err
}

func DelConcernById(id string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Concern{Concernid: id})
	return status, err
}

func GetConcernById(id string) (v *Concern, err error) {
	o := orm.NewOrm()
	v = &Concern{Concernid: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


func GetConcern(conditions string) (c *Concern,err error) {
	rowSql  := "select * from  "+c.TableName()+"  where true "+conditions+" limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowSql).QueryRow(&c)
	return
}