package rbac

import (
	//m "github.com/kasoly/admin/src/models"
)

type ItemController struct {
	BaseController
}

func (this *ItemController) Index() {
	this.Layout = this.GetTemplatetype() + "/layout/main.tpl"
	this.TplName = this.GetTemplatetype() + "/rbac/item.tpl"
}

