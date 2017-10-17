package controllers

import (
	"github.com/astaxie/beego"
	"common/wechat"
	"fmt"
)

// Wechat API
type WechatController struct {
	ApiController
}


// @Title       设置公众号菜单
// @Summary     设置公众号菜单
// @Description 设置公众号菜单
// @Param   token   header    string  true     "token"
// @Param	body	body  models.WxConfigForm  true  "{<br/>"openid":"openid"<br/>}
// @Success 200 "{<br/>"appid":"微信公众平台的AppID"<br/>,"secret":"微信公众平台的AppSecret"<br/>}
// @Failure 403 :openid is empty
// @router /setwxmenu [post]
func (this *WechatController) Setwxmenu()  {
	token  :=  beego.AppConfig.String("token") // 微信公众平台的Token
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	mp := wechat.New(token, appid, secret)
	//e:=mp.DeleteCustomMenu()
	//fmt.Println(e)
	var Button  []wechat.Button
	var SubButton  []wechat.Button
	SubButton1:= wechat.Button{
		Name:"书城",
		Type:"view",
		Url:"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa54fdefbe87687f7&redirect_uri=http://api.kasoly.com/&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect",
	}
	SubButton = append(SubButton,SubButton1)
	Button1 := wechat.Button{
		Name:"LoveBooks",
		SubButton:SubButton,
	}
	Button = append(Button,Button1)
	var SubButton2  []wechat.Button
	SubButton22 := wechat.Button{
		Name:"扫码推事件",
		Type:"scancode_push",
		Key:"scancode_push",
	}
	SubButton2 = append(SubButton2,SubButton22)
	Button2 := wechat.Button{
		Name:"扫码",
		SubButton:SubButton2,
	}
	Button = append(Button,Button2)
	b:=mp.CreateCustomMenu(&Button)
	fmt.Println(b)
	c,_:=mp.GetCustomMenu()
	this.Rsp(true,"成功!",&c)
}
