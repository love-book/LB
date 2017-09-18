package admin

import (
	"github.com/astaxie/beego"
	"github.com/kasoly/admin/src/rbac"
)

func router() {
	beego.Router("/rbac/item", &rbac.ItemController{}, "*:Index")
	beego.Router("/rbac/route/index", &rbac.RouteController{}, "*:Index")
	beego.Router("/rbac/route/addandedit", &rbac.RouteController{}, "*:AddAndEdit")
	beego.Router("/rbac/route/delroute", &rbac.RouteController{}, "*:DelRoute")

}

