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
	"net/http"
	"encoding/base64"
	"math/rand"
)

type ApiController struct {
	Userid string
	Openid string
	Province string
	Imgurl   string
	Nickname string
	beego.Controller
}

func (this *ApiController) Prepare() {
	//关闭xsrf校验
	this.EnableXSRF = false
	this.EnableRender =false
	this.SetUserId()
}

func (this *ApiController) SetUserId() {
	auth := this.Ctx.Request.Header.Get("token")
	if u,ok:=common.GetToken(auth);ok{
		split :=strings.Split(u.Appid,";")
		this.Userid   = split[0]
		this.Openid   = split[1]
		this.Province = split[2]
		this.Imgurl   = split[3]
		this.Nickname = split[4]
	}else{
		this.Userid   =  ""
		this.Openid   =  ""
		this.Province =  ""
		this.Imgurl   =  ""
		this.Nickname =  ""
	}
}

//提交数据重复校验
func  repeatCommit(m interface{}) (err error){
	bt,_:= json.Marshal(m)
	hash := md5.New()
	hash.Write(bt)
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
//获取验证码
func GetPhoneCode(size int)string{
	result:=map[int]string{
		1:"1",2:"2",3:"3",4:"4",
		5:"5",6:"6",7:"7",8:"8",9:"9",
	}
	var temp string = ""
	for i:=0 ; i<size ;  {
		temp += string(result[rand.Intn(9)])
		i++
	}
	return temp
}

func (this *ApiController) Rsp(status bool, msg string,data interface{}) {
	this.Data["json"] = &map[string]interface{}{"status": status, "msg": msg,"data":data}
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
		mp.ReplyTextMsg(this.Ctx.ResponseWriter, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa54fdefbe87687f7&redirect_uri=http://api.kasoly.com/&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect")
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
			//设置半个小时更新一次地理位置信息
			rc := comm.Pool.Get()
			defer rc.Close()
			v,err := redis.Int64(rc.Do("EXISTS",mp.Request.FromUserName))
			fmt.Println(err)
            if v == 0{
				l,err := common.GetLocation(mp.Request.Latitude,mp.Request.Longitude)
				if err == nil {
					province:=l.Result.AddressComponent.Province
					city:=l.Result.AddressComponent.City
					//将用户的经纬度信息加入redis GEO
					geoKey := province
					rc.Do("GEOADD",geoKey,mp.Request.Longitude,mp.Request.Latitude,mp.Request.FromUserName)
					rc.Do("SETEX",mp.Request.FromUserName, 60*30,mp.Request.Longitude)
					FromUserName := []string{mp.Request.FromUserName}
					user,err := models.GetUsersByOpenId(FromUserName)
					if err==nil{
						user.Long   = mp.Request.Longitude
						user.Lat    = mp.Request.Latitude
						user.Province = province
						user.City     = city
						user.Logintime = time.Now().Unix()
						models.UpdateUsersById(user)
					}
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
					user.Province	=  userInfo.Province
					user.City	    =  userInfo.City
					user.Address    =  ""
					user.Logintime	=  time.Now().Unix()
					user.Created_at =  time.Now().Unix()
					user.Updated_at =  time.Now().Unix()
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


func (this *ApiController)Test()  {
	go  models.Producter(models.BotMsgThree)
	//go models.Producter(models.BotMsgTwo)
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


//绑定微信用户
func (this *ApiController) BindWxUser() {
	n := models.Wxusers{}
	if err := this.ParseForm(&n); err != nil {
		this.Rsp(false, err.Error(),"")
		return
	}
	var id int64
	var err error
	var Nid int64
	var condtion = "AND bin = '"+fmt.Sprintf("%d",n.Bin)+"'"
	condtion += "AND remark_name = '"+n.RemarkName+"'"
	if user,err := models.GetWxuser(condtion);err == nil{
		fmt.Println(err)
		fmt.Println(user)
		Nid = int64(user.Id)
	}else{
		Nid, _= this.GetInt64("Id")
	}
	fmt.Println(Nid)
	if Nid > 0 {
		id, err = models.UpdateWxuser(&n)
	} else {
		id, err = models.AddWxuser(&n)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success","")
		return
	} else {
		this.Rsp(false, err.Error(),"")
		return
	}
}



func (this *ApiController) BotMsg() {
	if v := CheckAuth(this.Ctx.Request);!v {
		fmt.Println(v)
		this.Data["json"] = &map[string]interface{}{"code":0,"msg": "无操作权限!","data":""}
		this.ServeJSON()
	}
	var  helper models.Helper
	if v,_ := this.GetInt("Uin");v !=0 {
		helper.Uin = v
	}else{
		Json := map[string]interface{}{"code":200,"msg": "成功!","data":BotConfAndAction()}
		this.Data["json"] = &Json
		this.ServeJSON()
		return
	}
	Uin := strconv.Itoa(helper.Uin)
	//models.Cache.Delete(Uin)
	channel := models.Cache.Get(Uin)
    if channel == "" {
		if v := this.GetString("UserName");v !="" {
			helper.UserName = v
		}
		if v := this.GetString("NickName");v !="" {
			helper.NickName = v
		}
		var condtion = "AND uin = '"+Uin+"'"
		h,err := models.GetHelperByUserName(condtion)
		if err != nil{
			cond := " ORDER BY channel DESC"
			c,err := models.GetHelperByUserName(cond)
			if err != nil{
				fmt.Println(err)
				Json := map[string]interface{}{"code":0,"msg": "成功!","data":BotConfAndAction()}
				this.Data["json"] = &Json
				this.ServeJSON()
				return
			}
			//为新的小助手分配channel
			helper.Channel = c.Channel+1
			id,err := models.AddHelper(&helper)
			fmt.Println(err)
			fmt.Println(id)
		}else{
			//数据库查询到放入本地缓存
			err = models.Cache.Put(Uin,h.Channel)
			helper.Channel = h.Channel
		}
	}else{
		if value, ok := channel.(int);ok{
			helper.Channel = value
		}
	}
	var ok bool = false
	var ch  models.Channel
	switch helper.Channel {
	    case 1:
			ok  = ch.ChannelOne()
			goto MsgChannel
			break
		case 2:
			ok  = ch.ChannelTwo()
			goto MsgChannel
			break
		case 3:
			ok = ch.ChannelThree()
			goto MsgChannel
			break
		case 4:
			ok = ch.ChannelFour()
			goto MsgChannel
			break
		case 5:
			ok = ch.ChannelFive()
			goto MsgChannel
			break;
		case 6:
			ok = ch.ChannelSex()
			goto MsgChannel
			break
		case 7:
			ok = ch.ChannelSeven()
			goto MsgChannel
			break
		case 8:
			ok = ch.ChannelEigth()
			goto MsgChannel
			break
	    case 9:
			ok = ch.ChannelNine()
			goto MsgChannel
			break
		case 10:
			ok = ch.ChannelTen()
			goto MsgChannel
			break
	   default:
		   Json := map[string]interface{}{"code":999,"msg": "退出!","data":BotConfAndAction()}
		   this.Data["json"] = &Json
		   this.ServeJSON()
		   break
	}
	MsgChannel:
	   var Json   map[string]interface{}
	   fmt.Println(ok)
	   if ok == true {
		   Json = map[string]interface{}{"code":200,"msg": "成功!","data":BotConfAndAction()}
	   }else{
		   Json = map[string]interface{}{"code":0,"msg": "成功!","data":BotConfAndAction()}
	   }
	  this.Data["json"] = &Json
	  this.ServeJSON()
}

func CheckAuth(r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}
	//查询缓存或数据库
	if pair[0]=="kasoly" {
		return true
	}
	return false
}

//机器人配置信息和机器人动作
func BotConfAndAction() (msg []map[string]interface{}){
	conf := map[string]interface{}{
		"Text":`您好！关于您购买的【商业计划书模板案例融资大学生创业企业职业生涯规划项目策划研究】已发货：
        【解压版】链接: https://pan.baidu.com/s/1bpH45cj 密码: 8zfb
		【压缩版】链接：https://pan.baidu.com/s/1hsMJzIO 密码：6b3v
 		手机用户可以直接打开解压版，压缩版请下载到本地电脑解压，如文件超数量限制请分批保存。
		评价10字、5★ 好 评，自动赠送2000G 好 评 大礼包！给个 好 评 吧~~亲~~~！`,
		"Img":"http://n.sinaimg.cn/default/4_img/uplaod/3933d981/20170908/i9ew-fykuffc4351895.jpg",
		"Verify":"已添加小助手为好友!请按操作继续完成绑定信息!",
		"Action":"action",
		"MsgId":"",
		"RemarkName":"A1505124398",
		"Bin":"",
		"Channel":0,
		"MsgType":1990,
	}
	msg = append(msg,conf)
	return msg
}

