package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type ItemChild struct {
	Parent  string  `orm:"union;size(64);" form:"Parent" valid:"Required"`
	Child   string  `orm:"union;size(64);" form:"Child" valid:"Required"`
	Item    *Item   `orm:"null;rel(fk);on_delete(cascade);on_update(cascade)"`
}

func (c *ItemChild) TableName() string {
	return beego.AppConfig.String("rbac_item_child_table")
}


func init() {
	orm.RegisterModel(new(ItemChild))
}


func checkItemChild(c *ItemChild) (err error) {
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


//get ItemChild list
func GetItemChildlist(page int64, page_size int64, sort string) (rules []orm.Params, count int64) {
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

func AddItemChild(c *ItemChild) (int64, error) {
	if err := checkItemChild(c); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	child := new(ItemChild)
	child.Parent = c.Parent
	child.Child = c.Child
	id, err := o.Insert(child)
	return id, err
}

func UpdateItemChild(c *ItemChild) (int64, error) {
	if err := checkItemChild(c); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	itemChild := make(orm.Params)
	if len(c.Parent) > 0 {
		itemChild["Parent"] = c.Parent
	}
	if len(c.Child) > 0 {
		itemChild["Child"] = c.Child
	}
	if len(itemChild) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table ItemChild
	num, err := o.QueryTable(table).Filter("Parent", c.Child).Update(itemChild)
	return num, err
}

func DelItemChildByName(Name string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&ItemChild{Parent: Name})
	return status, err
}

func ItemChildList() (itemChilds []orm.Params) {
	o := orm.NewOrm()
	itemChild := new(ItemChild)
	qs := o.QueryTable(itemChild)
	qs.Values(&itemChilds, "parent", "child")
	return itemChilds
}
