package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"common"
	"common/wechat"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Prepare() {
	//关闭xsrf校验
	this.EnableXSRF = false
	this.EnableRender =false
}

//json输出
func (this *ApiController) renderJson(Json interface{}) {
	//Json := map[string]interface{}{"code":0,"msg": "成功!","data":"1111"}
	this.Data["json"] = Json
	this.ServeJSON()
	this.StopRun()
}

func (this *ApiController) Index() {
	token  :=  beego.AppConfig.String("token") // 微信公众平台的Token
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	// 仅被动响应消息时可不填写appid、secret
	// 仅主动发送消息时可不填写token
	mp := wechat.New(token, appid, secret)
	// 检查请求是否有效
	// 仅主动发送消息时不用检查
	if !mp.Request.IsValid(this.Ctx.ResponseWriter, this.Ctx.Request) {
		return
	}
	fmt.Println("CreateContainer,RequestBody", string(this.Ctx.Input.RequestBody))
	// 文本消息
	if mp.Request.MsgType == wechat.MsgTypeText {
		// 回复消息
		mp.ReplyTextMsg(this.Ctx.ResponseWriter, "Hello, 世界")
	}
	// 图片消息
	if mp.Request.MsgType == wechat.MsgTypeImage {
		// 回复消息
		mp.ReplyTextMsg(this.Ctx.ResponseWriter, "图片消息")
	}

	//事件
	if mp.Request.MsgType == wechat.MsgTypeEvent{
		fmt.Print(mp.Request.Event)
		// 地理位置事件
		if mp.Request.Event == wechat.MsgTypeLocation {

		}
		// 关注事件
		if mp.Request.Event == wechat.EventSubscribe {
			mp.ReplyTextMsg(this.Ctx.ResponseWriter, "感谢关注恋书!")
		}
		// 取关事件
		if mp.Request.Event == wechat.EventSubscribe {

		}

		// 扫码提示事件scancode_waitmsg
		if mp.Request.Event == wechat.EventScanCodeWaitMsg {
			if(mp.Request.ScanCodeInfo.ScanType == wechat.ScanTypeBarcode){
				scanRes := strings.Split(mp.Request.ScanCodeInfo.ScanResult,",")
				res ,_:= common.GetBarcodeInfo(scanRes[1])
				if(res.Code != "10000"){
					mp.ReplyTextMsg(this.Ctx.ResponseWriter, res.Msg)
				}
				goodsName := res.Result.Showapi_res_body.GoodsName
				fmt.Printf(goodsName)
				if(goodsName == ""){
					mp.ReplyTextMsg(this.Ctx.ResponseWriter, "书名:"+goodsName)
				}else{
					mp.ReplyTextMsg(this.Ctx.ResponseWriter, "书名:"+goodsName)
				}
			}else {
				mp.ReplyTextMsg(this.Ctx.ResponseWriter, "暂不支持非条形码扫码!")
			}

		}


	}
}

func (this *ApiController) Login()  {

	token,err := common.GetToken("1982")
	if 0 == len(token) {
		fmt.Println(token,err.Error())
		fmt.Sprintf("get token failed")
		this.StopRun()
	}
	fmt.Println(token)
	this.Data["json"] = common.Actionsuccess
	this.ServeJSON()

	/*idworker,_ := common.NewIdWorker(12,5)
	id,_ := idworker.NextId()
	fmt.Println(id)*/
	//s,_:= common.GetAccessToken("1203939180183","os01")
	//common.VerifyAccessToken(s,"1203939180183","os01")

}

func (this *ApiController)Test()  {
	this.Data["json"] = common.ErrPermission
	this.ServeJSON()
	this.StopRun()
}

func (this *ApiController)StopRunning()  {
	param := this.Ctx.Input.Param(":type")
	if param == "1" {
		this.Data["json"] = common.ErrPermission
	}else{
		this.Data["json"] = common.ErrExpired
	}

	this.ServeJSON()
	this.StopRun()
}
