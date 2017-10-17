package rabbitmq

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	"common/wechat"
)

func wechatConf()*wechat.Weixinmp{
	token  :=  beego.AppConfig.String("token") // 微信公众平台的Token
	appid  :=  beego.AppConfig.String("appid")  // 微信公众平台的AppID
	secret :=  beego.AppConfig.String("secret")  // 微信公众平台的AppSecret
	// 仅主动发送消息时可不填写token
	return  wechat.New(token, appid, secret)
}

func ConcernNoticeRegister(d *Consumer) {
	d.HandlerRegister.Add(d.HandlerTag, Handler(concernNotice), "concernNotice")
	if err := d.HandlerRegister.EnableByName("concernNotice"); err != nil {
		fmt.Println(err)
	}
}

// 收藏推送
func concernNotice(d *Consumer) {
	var body map[string]string
	json.Unmarshal([]byte(d.Body),&body)
	mp:= wechatConf()
	var articleMap []wechat.Article
	article := wechat.Article{
          Title:"ABC：物联网从虚到实的地图",
		  Description:"所谓IoT命题，已经被提出了很多年。但现状是相对底层的基础硬件和技术市场处在高度活跃状态。但在应用端，IoT进入垂直行业、在工业生产体系中大规模部署的比率却偏低。相对而言，传统企业在选择物联网之路时依旧趋向保守。",
          PicUrl:"http://imgs.ebrun.com/resources/2017_09/2017_09_30/2017093020015067517909651.png",
		  Url:"http://www.ebrun.com/20170930/248731.shtml",
	}
	articleMap = append(articleMap,article)
	mp.SendNewsMsg(body["OpenIdFrom"],&articleMap)
	mp.SendTextMsg(body["OpenIdFrom"],body["UserName"]+"关注了你,快去看看吧!")
}



func LibraryrequestRegister(d *Consumer) {
	d.HandlerRegister.Add(d.HandlerTag, Handler(MLibraryRequest), "MLibraryRequest")
	if err := d.HandlerRegister.EnableByName("MLibraryRequest"); err != nil {
		fmt.Println(err)
	}
}

// 借书请求
func MLibraryRequest(d *Consumer) {
	var body map[string]string
	json.Unmarshal([]byte(d.Body),&body)
	mp:= wechatConf()
	var articleMap []wechat.Article
	article := wechat.Article{
		Title:"ABC：物联网从虚到实的地图",
		Description:"所谓IoT命题，已经被提出了很多年。但现状是相对底层的基础硬件和技术市场处在高度活跃状态。但在应用端，IoT进入垂直行业、在工业生产体系中大规模部署的比率却偏低。相对而言，传统企业在选择物联网之路时依旧趋向保守。",
		PicUrl:"http://imgs.ebrun.com/resources/2017_09/2017_09_30/2017093020015067517909651.png",
		Url:"http://www.ebrun.com/20170930/248731.shtml",
	}
	articleMap = append(articleMap,article)
	mp.SendNewsMsg(body["OpenIdFrom"],&articleMap)
	mp.SendTextMsg(body["OpenIdFrom"],body["UserName"]+"向你发起借书请求,快去看看吧!")
}



func   AgreelibraryrequestRegister(d *Consumer) {
	d.HandlerRegister.Add(d.HandlerTag, Handler(MAgreelibraryrequest), "MAgreelibraryrequest")
	if err := d.HandlerRegister.EnableByName("MAgreelibraryrequest"); err != nil {
		fmt.Println(err)
	}
}

// 同意借书
func MAgreelibraryrequest(d *Consumer) {
	var body map[string]string
	json.Unmarshal([]byte(d.Body),&body)
	mp:= wechatConf()
	var articleMap []wechat.Article
	article := wechat.Article{
		Title:"ABC：物联网从虚到实的地图",
		Description:"所谓IoT命题，已经被提出了很多年。但现状是相对底层的基础硬件和技术市场处在高度活跃状态。但在应用端，IoT进入垂直行业、在工业生产体系中大规模部署的比率却偏低。相对而言，传统企业在选择物联网之路时依旧趋向保守。",
		PicUrl:"http://imgs.ebrun.com/resources/2017_09/2017_09_30/2017093020015067517909651.png",
		Url:"http://www.ebrun.com/20170930/248731.shtml",
	}
	articleMap = append(articleMap,article)
	mp.SendNewsMsg(body["OpenIdFrom"],&articleMap)
	mp.SendTextMsg(body["OpenIdFrom"],body["UserName"]+"同意你的借书请求,快去看看吧!")
}

