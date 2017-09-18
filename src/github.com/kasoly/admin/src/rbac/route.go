package rbac

import (
	m "github.com/kasoly/admin/src/models"
)

type RouteController struct {
	BaseController
}

func (this *RouteController) Index() {
	if this.IsAjax() {
		start, _ := this.GetInt64("start",0)
		length, _ := this.GetInt64("length",10)
		draw, _ := this.GetInt64("draw",1)
		conditions:=""
		if Id := this.GetString("Id");Id!=""{
			conditions += " AND id ="+Id
		}
		if Name := this.GetString("Name");Name!=""{
			conditions += " AND name like '%"+Name+"%'"
		}
		conditions += " ORDER BY id DESC "
		routes,count := m.GetRoutelist(start,length,conditions)
		if len(routes) < 1 {
			routes = []*m.Route{}
		}
		this.Data["json"] = &map[string]interface{}{"draw":draw,"recordsTotal": count,"recordsFiltered":count,"data":&routes}
		this.ServeJSON()
		return
	} else {
		this.Layout = this.GetTemplatetype() + "/layout/main.tpl"
		this.TplName = this.GetTemplatetype() + "/rbac/route.tpl"
	}
}
func (this *RouteController) AddAndEdit() {
	n := m.Route{}
	if err := this.ParseForm(&n); err != nil {
		this.Rsp(false, err.Error(),"")
		return
	}
	var id int64
	var err error
	Nid, _ := this.GetInt64("Id")
	if Nid > 0 {
		id, err = m.UpdateRoute(&n)
	} else {
		id, err = m.AddRoute(&n)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success","")
		return
	} else {
		this.Rsp(false, err.Error(),"")
		return
	}
}
func (this *RouteController) DelRoute() {
	Id,_:= this.GetInt64("Id")
	status, err := m.DelRouteById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success","")
		return
	} else {
		this.Rsp(false, err.Error(),"")
		return
	}
}