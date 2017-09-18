package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type Item struct {
	Name  string  `orm:"pk;size(64);column(name);" form:"Name" valid:"Required"`
	Type  int16   `form:"Type" valid:"Required"`
	Description  string  `form:"Description"`
	RuleName  string  `orm:"size(64);form:"RuleName"`
	Data   string   `form:"Data"`
	CreatedAt  int32 `orm:"size(11);" form:"CreatedAt" valid:"Required"`
	UpdatedAt  int32 `orm:"size(11);" form:"UpdatedAt" valid:"Required"`
    Rule    *Rule   `orm:"null;rel(fk);on_delete(set_null);on_update(cascade)"`
	ItemChild   []*ItemChild  `orm:"reverse(many)"`
	Assignment  []*Assignment `orm:"reverse(many)"`
}

func (i *Item) TableName() string {
	return beego.AppConfig.String("rbac_item_table")
}

// 多字段索引
func (i *Item) TableIndex() [][]string {
	return [][]string{
		[]string{"Type"},
	}
}


func init() {
	orm.RegisterModel(new(Item))
}


func checkItem(i *Item) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&i)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//get Item list
func GetItemlist(page int64, page_size int64, sort string) (rules []orm.Params, count int64) {
	o := orm.NewOrm()
	item := new(Item)
	qs := o.QueryTable(item)
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

func AddItem(i *Item) (int64, error) {
	if err := checkItem(i); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	item := new(Item)
	item.Name = i.Name
	item.Data = i.Data
	item.CreatedAt = i.CreatedAt
	item.UpdatedAt = i.UpdatedAt
	id, err := o.Insert(item)
	return id, err
}

func UpdateItem(i *Item) (int64, error) {
	if err := checkItem(i); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	item := make(orm.Params)
	if len(i.Name) > 0 {
		item["Name"] = i.Name
	}
	if i.Data != "" {
		item["Data"] = i.Data
	}
	if i.CreatedAt != 0 {
		item["CreatedAt"] = i.CreatedAt
	}
	if i.UpdatedAt != 0 {
		item["UpdatedAt"] = i.UpdatedAt
	}
	if len(item) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Item
	num, err := o.QueryTable(table).Filter("Name", i.Name).Update(item)
	return num, err
}

func DelItemByName(Name string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Rule{Name: Name})
	return status, err
}

func ItemList() (items []orm.Params) {
	o := orm.NewOrm()
	item := new(Item)
	qs := o.QueryTable(item)
	qs.Values(&items, "name", "data")
	return items
}
