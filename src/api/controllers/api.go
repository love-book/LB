package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"common"
	"common/wechat"
	"encoding/json"
	"crypto/md5"
	comm "common/conndatabase"
	"models"
	"time"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Prepare() {
	//关闭xsrf校验
	this.EnableXSRF = false
	this.EnableRender =false
}

//提交数据重复校验
func  repeatCommit(m interface{}) (err error){
	byte,_:= json.Marshal(m)
	hash := md5.New()
	hash.Write(byte)
	md5 := fmt.Sprintf("%x", hash.Sum(nil))
	// 从池里获取连接
	rc := comm.Pool.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	var v interface{}
	v,err = rc.Do("EXISTS",md5)
	fmt.Println(err)
	fmt.Println(v)
	if err != nil {
		v,_= rc.Do("SETEX",md5,6,md5)
		fmt.Println(v)
	}
	return
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
		// 地理位置事件
		fmt.Println("Event:", mp.Request.Event)
		if mp.Request.Event == wechat.EventLocation {
			//设置15分钟更新一次地理位置信息
			rc := comm.Pool.Get()
			defer rc.Close()
			v,err := redis.Int64(rc.Do("EXISTS",mp.Request.FromUserName))
			fmt.Println(err)
			fmt.Println(v)
            if v == 0{
				l,err := common.GetGps(mp.Request.Latitude,mp.Request.Longitude)
				fmt.Println(l)
				if err == nil {
					var long,lat float64
					for _,vs := range l.Result{
						long = vs.X
						lat  = vs.Y
						break
					}
					//将用户的经纬度信息加入redis GEO
					rc.Do("GEOADD",comm.LocationGeo,long,lat,mp.Request.FromUserName)
					rc.Do("SETEX",mp.Request.FromUserName, 60*15,long)
				}
			}
		}
		// 关注事件
		if mp.Request.Event == wechat.EventSubscribe {
			//获取微信用户信息
			userInfo,err:= mp.GetUserInfo(mp.Request.FromUserName)
			//为用户注册账号
			if err ==nil{
				FromUserName := []string{mp.Request.FromUserName}
				_,errs := models.GetUsersByOpenId(FromUserName)
				fmt.Println(errs)
				if errs != nil{
					user := models.Users{}
					id := models.GetID()
					user.Userid    =   fmt.Sprintf("%d", id)
					user.Nickname  =   userInfo.Nickname
					user.Openid    =   mp.Request.FromUserName
					user.Wnickname =   userInfo.Nickname
					user.Wimgurl   =   userInfo.Headimgurl
					user.Imgurl    =   userInfo.Headimgurl
					user.Gender    =   userInfo.Sex
					user.Age       =   0
					user.Telphone  = ""
					user.Qq    = ""
					user.Weino = ""
					user.Signature  = ""
					user.Address    =  userInfo.Province+","+userInfo.City
					user.Created_at =  time.Now().Unix()
					user.Updated_at = 0
					err := user.InsertValidation()
					if err != nil {
						mp.ReplyTextMsg(this.Ctx.ResponseWriter, err.Error())
					}
					_,err = models.AddUsers(&user)
					if err != nil {
						mp.ReplyTextMsg(this.Ctx.ResponseWriter, err.Error())
					}
				}
			}
			mp.ReplyTextMsg(this.Ctx.ResponseWriter, "感谢关注恋书!")
		}
		// 取关事件
		if mp.Request.Event == wechat.EventSubscribe {

		}
		// 扫码提示事件scancode_waitmsg
		if mp.Request.Event == wechat.EventScanCodeWaitMsg {
			if(mp.Request.ScanCodeInfo.ScanType == wechat.ScanTypeBarcode){
				scanRes := strings.Split(mp.Request.ScanCodeInfo.ScanResult,",")
				isbn:= scanRes[1]
				//查询数据库是否存在该条形码
				model,err:= models.GetIbsn(isbn)
				if err != nil {
					//state 1:扫码正常添加，3:扫码没有查询到结果待补充
					var state uint8 = 1
					res, err:= common.GetBarcodeInfo(isbn)
					//查询失败
					if err != nil{
						mp.ReplyTextMsg(this.Ctx.ResponseWriter, "网络繁忙!")
					}
					//查询不到信息
					if  res.Charge == false{
						state = 3
						res.Result.Showapi_res_body.GoodsName = "未知"
						res.Result.Showapi_res_body.ManuName = ""
						res.Result.Showapi_res_body.Img = ""
						res.Result.Showapi_res_body.Price = "0"
					}
					body :=  &res.Result.Showapi_res_body
					model := models.Books{}
					Uid := models.GetID()
					model.Bookid   = fmt.Sprintf("%d", Uid)
					model.Bookname = body.GoodsName
					model.Author   = body.ManuName
					model.Imagehead = ""
					model.Imageback = ""
					model.Imageurl  = body.Img
					model.Describe =body.GoodsName
					price, _:= strconv.ParseUint(body.Price, 10, 16)
					model.Price  = uint16(price)
					model.Isbn   = isbn
					model.Depreciation = 0
					model.State = state
					//根据openid获取用户id
					FromUserName := []string{mp.Request.FromUserName}
					UserInfo,errs := models.GetUsersByOpenId(FromUserName)
					if errs !=nil{
						mp.ReplyTextMsg(this.Ctx.ResponseWriter,"暂无信息!请稍后重试")
					}
					model.Userid = UserInfo.Userid
					err = comm.Insert(model)
					if err != nil{
						mp.ReplyTextMsg(this.Ctx.ResponseWriter,"没有当前图书信息!")
					}
				}
				book := models.Bookrack{}
				book.Userid  = model.Userid
				book.Bookid  = model.Bookid
				book.Book_state = "1"
				book.Is_borrow  = "1"
				_,err = models.GetBookByUidAndBookId(model.Userid,book.Bookid)
				if err != nil {
					//加入用户书架
					id := models.GetID()
					book.Bookqid = fmt.Sprintf("%d",id)
					t:=time.Now().Unix()
					book.Create_time = t
					book.Update_time = t
					err := comm.Insert(&book)
					if err == nil {
						mp.ReplyTextMsg(this.Ctx.ResponseWriter,"图书成功添加!请查看个人书架")
					}
				}else{
					mp.ReplyTextMsg(this.Ctx.ResponseWriter,"您的个人书架存在当前图书")
				}
				mp.ReplyTextMsg(this.Ctx.ResponseWriter,"暂无信息!请稍后重试")
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
