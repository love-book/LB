package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
)
type Assignment struct {
	ItemName  string  `orm:"union;size(64);" form:"ItemName" valid:"Required"`
	UserId   string  `orm:"union;size(64);" form:"UserId" valid:"Required"`
	CreatedAt  int32 `orm:"size(11);" form:"CreatedAt" valid:"Required"`
	Item    *Item   `orm:"null;rel(fk);on_delete(cascade);on_update(cascade)"`
}

func (a *Assignment) TableName() string {
	return beego.AppConfig.String("rbac_assignment_table")
}


func init() {
	orm.RegisterModel(new(Assignment))
}


func checkAssignment(a *Assignment) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&a)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//get Assignment list
func GetAssignmentlist(page int64, page_size int64, sort string) (assignments []orm.Params, count int64) {
	o := orm.NewOrm()
	assignment := new(Item)
	qs := o.QueryTable(assignment)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&assignments)
	count, _ = qs.Count()
	return assignments, count
}

func AddAssignment(a *Assignment) (int64, error) {
	if err := checkAssignment(a); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	assignment := new(Assignment)
	assignment.ItemName = a.ItemName
	assignment.UserId = a.UserId
	id, err := o.Insert(assignment)
	return id, err
}

func UpdateAssignment(a *Assignment) (int64, error) {
	if err := checkAssignment(a); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	assignment := make(orm.Params)
	if len(a.ItemName) > 0 {
		assignment["ItemName"] = a.ItemName
	}
	if len(a.UserId) > 0 {
		assignment["UserId"] = a.UserId
	}
	if len(assignment) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Assignment
	num, err := o.QueryTable(table).Filter("ItemName", a.ItemName).Update(assignment)
	return num, err
}

func DelAssignmentByName(Name string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Assignment{ItemName: Name})
	return status, err
}

func AssignmentList() (assignments []orm.Params) {
	o := orm.NewOrm()
	assignment := new(Assignment)
	qs := o.QueryTable(assignment)
	qs.Values(&assignments, "ItemName", "UserId")
	return assignments
}
