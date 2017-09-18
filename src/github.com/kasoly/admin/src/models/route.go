package models
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	"strconv"
)
type Route struct {
	Id     int64    `orm:"size(11)" form:"Id"`
	Name   string   `orm:"size(128)" form:"Name" valid:"Required"`
}

func (r *Route) TableName() string {
	return beego.AppConfig.String("rbac_route_table")
}

func init() {
	orm.RegisterModel(new(Route))
}


func checkRoute(m *Route) (err error) {
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

//get Route list
func GetRoutelist(start int64, length int64,conditions string) (routes []*Route, count int64) {
	/*o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	qs.Limit(length,start).OrderBy(sort).Values(&routes)
	count, _ = qs.Count()
	return routes, count*/
	countSql:= "select count(*) from  auth_route  where true "+conditions
	rowSql  := "select * from  auth_route  where true "+conditions+" limit "+strconv.FormatInt(start,10)+","+strconv.FormatInt(length,10)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&count)
	o.Raw(rowSql).QueryRows(&routes)
	return routes, count
}

func AddRoute(r *Route) (int64, error) {
	if err := checkRoute(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	route := new(Route)
	route.Name = r.Name
	id, err := o.Insert(route)
	return id, err
}

func UpdateRoute(r *Route) (int64, error) {
	if err := checkRoute(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	route := make(orm.Params)
	if len(r.Name) > 0 {
		route["Name"] = r.Name
	}
	if len(route) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Route
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(route)
	return num, err
}

func DelRouteById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Route{Id: Id})
	return status, err
}

func RouteList() (routes []orm.Params) {
	o := orm.NewOrm()
	route := new(Route)
	qs := o.QueryTable(route)
	qs.Values(&routes, "name")
	return routes
}
