package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type Rule struct {
	Name  string `orm:"pk;size(64);column(name);" form:"Name" valid:"Required"`
	Data  string  `form:"Data"`
	CreatedAt  int32 `orm:"size(11);" form:"CreatedAt" valid:"Required"`
	UpdatedAt  int32 `orm:"size(11);" form:"UpdatedAt" valid:"Required"`
	Item  []*Item  `orm:"reverse(many)"`
}

func (r *Rule) TableName() string {
	return beego.AppConfig.String("rbac_rule_table")
}

func init() {
	orm.RegisterModel(new(Rule))
}


func checkRule(r *Rule) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&r)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//get Rule list
func GetRulelist(page int64, page_size int64, sort string) (rules []orm.Params, count int64) {
	o := orm.NewOrm()
	rule := new(Rule)
	qs := o.QueryTable(rule)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&rules)
	count, _ = qs.Count()
	return rules, count
}

func AddRule(r *Rule) (int64, error) {
	if err := checkRule(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	rule := new(Rule)
	rule.Name = r.Name
	rule.Data = r.Data
	rule.CreatedAt = r.CreatedAt
	rule.UpdatedAt = r.UpdatedAt
	id, err := o.Insert(rule)
	return id, err
}

func UpdateRule(r *Rule) (int64, error) {
	if err := checkRule(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	rule := make(orm.Params)
	if len(r.Name) > 0 {
		rule["Name"] = r.Name
	}
	if r.Data != "" {
		rule["Data"] = r.Data
	}
	if r.CreatedAt != 0 {
		rule["CreatedAt"] = r.CreatedAt
	}
	if r.UpdatedAt != 0 {
		rule["UpdatedAt"] = r.UpdatedAt
	}
	if len(rule) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Rule
	num, err := o.QueryTable(table).Filter("Name", r.Name).Update(rule)
	return num, err
}

func DelRuleByName(Name string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Rule{Name: Name})
	return status, err
}

func RuleList() (rules []orm.Params) {
	o := orm.NewOrm()
	rule := new(Rule)
	qs := o.QueryTable(rule)
	qs.Values(&rules, "name", "data")
	return rules
}
