package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:BooknewsController"] = append(beego.GlobalControllerRouter["api/controllers:BooknewsController"],
		beego.ControllerComments{
			Method: "Newslist",
			Router: `/newsList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooknewsController"] = append(beego.GlobalControllerRouter["api/controllers:BooknewsController"],
		beego.ControllerComments{
			Method: "Libraryrequest",
			Router: `/libraryrequest`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooknewsController"] = append(beego.GlobalControllerRouter["api/controllers:BooknewsController"],
		beego.ControllerComments{
			Method: "Libraryrequestupdate",
			Router: `/libraryrequestupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BookorderController"] = append(beego.GlobalControllerRouter["api/controllers:BookorderController"],
		beego.ControllerComments{
			Method: "Orderlist",
			Router: `/orderList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BookorderController"] = append(beego.GlobalControllerRouter["api/controllers:BookorderController"],
		beego.ControllerComments{
			Method: "Orderadd",
			Router: `/orderadd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BookorderController"] = append(beego.GlobalControllerRouter["api/controllers:BookorderController"],
		beego.ControllerComments{
			Method: "Orderupdate",
			Router: `/orderupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "Bookracklist",
			Router: `/bookracklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Bookrackadd",
			Router: `/bookrackadd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Bookrackaddbysn",
			Router: `/bookrackaddbysn`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Bookrackupdate",
			Router: `/bookrackupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Userinfo",
			Router: `/userinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Useradd",
			Router: `/useradd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Userupdate",
			Router: `/userupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
