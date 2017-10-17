package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)

type Opinions struct {
	Id                string    `orm:"pk;size(20);column(id);"`
	Userid            string	`orm:"size(10);" form:"userid" valid:"Required"`
	Opinions          string	`form:"opinions"`
	Images            string	`form:"images"`
	CreatedAt         int64	    `orm:"size(10);" form:"CreatedAt" valid:"Required"`
}

func (h *Opinions) TableName() string {
	return beego.AppConfig.String("table_opinions")
}

func init() {
	orm.RegisterModel(new(Opinions))
}


func checkOpinions(h *Opinions) (err error) {
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


func GetOpinionslist(page int64, page_size int64, sort string) (helpers []orm.Params, count int64) {
	o := orm.NewOrm()
	helper := new(Opinions)
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

func AddOpinions(h *Opinions) (int64, error) {
	if err := checkOpinions(h); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	helper := new(Opinions)
	helper.Id = h.Id
	helper.Userid  = h.Userid
	helper.Opinions  = h.Opinions
	helper.Images   = h.Images
	helper.CreatedAt  = h.CreatedAt
	id, err := o.Insert(helper)
	return id, err
}
