package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Useradd",
			Router: `/useradd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Accesstoken",
			Router: `/accesstoken`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Getwxconfig",
			Router: `/getwxconfig`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Getusertokenbycode",
			Router: `/getusertokenbycode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Phonecode",
			Router: `/phonecode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AppController"] = append(beego.GlobalControllerRouter["api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Checkponecode",
			Router: `/checkphonecode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "Agreelibraryrequest",
			Router: `/agreelibraryrequest`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooknewsController"] = append(beego.GlobalControllerRouter["api/controllers:BooknewsController"],
		beego.ControllerComments{
			Method: "Refuselibraryrequest",
			Router: `/refuselibraryrequest`,
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
			Method: "Orderupdate",
			Router: `/orderupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BookorderController"] = append(beego.GlobalControllerRouter["api/controllers:BookorderController"],
		beego.ControllerComments{
			Method: "Orderdelete",
			Router: `/orderdelete`,
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
			Method: "Bookaddbycode",
			Router: `/bookaddbycode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Addconcern",
			Router: `/addconcern`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Delbookconcern",
			Router: `/delbookconcern`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "ConcernBookList",
			Router: `/concernbooklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "ConcernUserList",
			Router: `/concernuserlist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksController"] = append(beego.GlobalControllerRouter["api/controllers:BooksController"],
		beego.ControllerComments{
			Method: "Getuserbookinfo",
			Router: `/getuserbookinfo`,
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
			Method: "Bookrackdelete",
			Router: `/bookrackdelete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Mybookrack",
			Router: `/mybookrack`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Getbookusers",
			Router: `/getbookusers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Getbookinfo",
			Router: `/getbookinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Getuserinfo",
			Router: `/getuserinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Getmybooklist",
			Router: `/getmybooklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:BooksrackController"] = append(beego.GlobalControllerRouter["api/controllers:BooksrackController"],
		beego.ControllerComments{
			Method: "Getbookbysn",
			Router: `/getbookbysn`,
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
			Method: "Updateuser",
			Router: `/updateuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetUsersByLocaltion",
			Router: `/getusersbylocaltion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "AddLocaltionByID",
			Router: `/addlocaltionbyid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "AddOpinions",
			Router: `/addopinions`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:UsersController"] = append(beego.GlobalControllerRouter["api/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Uploadfile",
			Router: `/uploadfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:WechatController"] = append(beego.GlobalControllerRouter["api/controllers:WechatController"],
		beego.ControllerComments{
			Method: "Setwxmenu",
			Router: `/setwxmenu`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
