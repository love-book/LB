package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["app/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["app/controllers/admin:UserController"],
		beego.ControllerComments{
			Method: "Userlist",
			Router: `/userlist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["app/controllers/admin:UserController"],
		beego.ControllerComments{
			Method: "Useradd",
			Router: `/useradd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["app/controllers/admin:UserController"],
		beego.ControllerComments{
			Method: "Userdelete",
			Router: `/userdelete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["app/controllers/admin:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/userupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
