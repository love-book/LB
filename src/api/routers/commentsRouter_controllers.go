package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "BookList",
			Router: `/booklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Bookadd",
			Router: `/bookadd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Bookupdate",
			Router: `/bookupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Uploadfile",
			Router: `/uploadfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Bookaddbycode",
			Router: `/bookaddbycode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "BookrackList",
			Router: `/booklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Bookadd",
			Router: `/bookadd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Bookupdate",
			Router: `/bookupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Userinfo",
			Router: `/userinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Useradd",
			Router: `/useradd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/userupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
