package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"time"
	"sort"
	"strings"
	"crypto/sha1"
	"strconv"
)

const (
	// request message types
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeVideo      = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation   = "location"
	MsgTypeLink       = "link"
	MsgTypeEvent      = "event"
	// event types
	EventSubscribe   = "subscribe"
	EventUnsubscribe = "unsubscribe"
	EventScan        = "SCAN"
	EventLocation    = "LOCATION"
	EventClick       = "CLICK"
	EventView        = "VIEW"
	EventScanCodeWaitMsg  = "scancode_waitmsg"
	EventScanCodePush     = "scancode_push"

	// media types
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
	// button types
	ButtonTypeClick = "click"
	ButtonTypeView  = "view"
	// environment constants
	UrlPrefix      = "https://api.weixin.qq.com/cgi-bin/"
	MediaUrlPrefix = "http://file.api.weixin.qq.com/cgi-bin/media/"
	retryNum       = 3

	// scanType
	ScanTypeQrcode  = "qrcode"   //二维码
	ScanTypeBarcode = "barcode"  //条形码
)

type Weixinmp struct {
	Request     Request
	AccessToken AccessToken
	TicketToken TicketToken
}

func New(token, appId, appSecret string) *Weixinmp {
	accessToken:=AccessToken{AppId: appId, AppSecret: appSecret}
	access,_:=accessToken.Fresh()
	ticketToken:=TicketToken{AppId:appId,AccessToken:access}
	return &Weixinmp{
		Request:     Request{Token: token},
		AccessToken: accessToken,
		TicketToken: ticketToken,
	}
}

// message structs
type msgHeader struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	ToUserName   string   `json:"touser"`
	FromUserName string   `json:"-"`
	CreateTime   int64    `json:"-"`
	MsgType      string   `json:"msgtype"`
}

type textMsg struct {
	msgHeader
	Content string `json:"-"`
	Text    struct {
		Content string `xml:"-" json:"content"`
	} `xml:"-" json:"text"`
}

type imageMsg struct {
	msgHeader
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type voiceMsg struct {
	msgHeader
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

type videoMsg struct {
	msgHeader
	Video *Video `json:"video"`
}

type musicMsg struct {
	msgHeader
	Music *Music `json:"music"`
}

type newsMsg struct {
	msgHeader
	ArticleCount int `json:"-"`
	Articles     struct {
		Item *[]Article `xml:"item" json:"articles"`
	} `json:"news"`
}

// 群发图文消息
type newsGroupMsg struct {
	Filter struct {
		IsToAll bool   `json:"is_to_all"`
		GroupId string `json:"group_id"`
	} `json:"filter"`

	Mpnews struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
	MsgType string `json:"msgtype"`
}

type Video struct {
	MediaId     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Music struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HQMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"picurl"`
	Url         string `json:"url"`
}

// reply text message
func (this *Weixinmp) ReplyTextMsg(rw http.ResponseWriter, content string) error {
	var msg textMsg
	msg.MsgType = "text"
	msg.Content = content
	return this.replyMsg(rw, &msg)
}

// reply image message
func (this *Weixinmp) ReplyImageMsg(rw http.ResponseWriter, mediaId string) error {
	var msg imageMsg
	msg.MsgType = "image"
	msg.Image.MediaId = mediaId
	return this.replyMsg(rw, &msg)
}

// reply voice message
func (this *Weixinmp) ReplyVoiceMsg(rw http.ResponseWriter, mediaId string) error {
	var msg voiceMsg
	msg.MsgType = "voice"
	msg.Voice.MediaId = mediaId
	return this.replyMsg(rw, &msg)
}

// reply video message
func (this *Weixinmp) ReplyVideoMsg(rw http.ResponseWriter, video *Video) error {
	var msg videoMsg
	msg.MsgType = "video"
	msg.Video = video
	return this.replyMsg(rw, &msg)
}

// reply music message
func (this *Weixinmp) ReplyMusicMsg(rw http.ResponseWriter, music *Music) error {
	var msg musicMsg
	msg.MsgType = "music"
	msg.Music = music
	return this.replyMsg(rw, &msg)
}

// reply news  message
func (this *Weixinmp) ReplyNewsMsg(rw http.ResponseWriter, articles *[]Article) error {
	var msg newsMsg
	msg.MsgType = "news"
	msg.ArticleCount = len(*articles)
	msg.Articles.Item = articles
	return this.replyMsg(rw, &msg)
}

// reply message
func (this *Weixinmp) replyMsg(rw http.ResponseWriter, msg interface{}) error {
	v := reflect.ValueOf(msg).Elem()
	v.FieldByName("ToUserName").SetString(this.Request.FromUserName)
	v.FieldByName("FromUserName").SetString(this.Request.ToUserName)
	v.FieldByName("CreateTime").SetInt(time.Now().Unix())
	data, err := xml.Marshal(msg)
	if err != nil {
		return err
	}
	if _, err := rw.Write(data); err != nil {
		return err
	}
	return nil
}

// send text message
func (this *Weixinmp) SendTextMsg(touser string, content string) error {
	var msg textMsg
	msg.MsgType = "text"
	msg.Text.Content = content
	return this.sendMsg(touser, &msg)
}

// send image message
func (this *Weixinmp) SendImageMsg(touser string, mediaId string) error {
	var msg imageMsg
	msg.MsgType = "image"
	msg.Image.MediaId = mediaId
	return this.sendMsg(touser, &msg)
}

// send voice message
func (this *Weixinmp) SendVoiceMsg(touser string, mediaId string) error {
	var msg voiceMsg
	msg.MsgType = "voice"
	msg.Voice.MediaId = mediaId
	return this.sendMsg(touser, &msg)
}

// send video message
func (this *Weixinmp) SendVideoMsg(touser string, video *Video) error {
	var msg videoMsg
	msg.MsgType = "video"
	msg.Video = video
	return this.sendMsg(touser, &msg)
}

// send music message
func (this *Weixinmp) SendMusicMsg(touser string, music *Music) error {
	var msg musicMsg
	msg.MsgType = "music"
	msg.Music = music
	return this.sendMsg(touser, &msg)
}

// send news message
func (this *Weixinmp) SendNewsMsg(touser string, articles *[]Article) error {
	var msg newsMsg
	msg.MsgType = "news"
	msg.Articles.Item = articles
	return this.sendMsg(touser, &msg)
}

// send message
func (this *Weixinmp) sendMsg(touser string, msg interface{}) error {
	v := reflect.ValueOf(msg).Elem()
	v.FieldByName("ToUserName").SetString(touser)
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%smessage/custom/send?access_token=", UrlPrefix)
	buf := bytes.NewBuffer(data)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		if _, err := post(url+token, "text/plain", buf); err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		break // success
	}
	return nil
}

// 向全部用户群发图文消息
func (this *Weixinmp) sendNewsToALl(mediaId string) error {
	var news newsGroupMsg
	news.MsgType = "mpnews"
	news.Filter.IsToAll = true
	news.Mpnews.MediaId = mediaId
	return this.sendGroupMsg(news)
}

// 向特定GroupId用户群发图文消息
func (this *Weixinmp) sendNewsToGroup(groupId string, mediaId string) error {
	var news newsGroupMsg
	news.MsgType = "mpnews"
	news.Filter.IsToAll = false
	news.Filter.GroupId = groupId
	news.Mpnews.MediaId = mediaId
	return this.sendGroupMsg(news)
}

// 群发消息
func (this *Weixinmp) sendGroupMsg(msg interface{}) error {
	url := fmt.Sprintf("%smessage/mass/sendall?access_token=", UrlPrefix)
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		if _, err := post(url+token, "text/plain", buf); err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		break // success
	}
	return nil
}

type qrScene struct {
	ExpireSeconds int64  `json:"expire_seconds,omitempty"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneId int64 `json:"scene_id"`
		} `json:"scene"`
	} `json:"action_info"`
}

// get qrcode url
func (this *Weixinmp) GetQRCodeURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + ticket
}

// create permanent qrcode
func (this *Weixinmp) CreateQRScene(sceneId int64) (string, error) {
	var inf qrScene
	inf.ActionName = "QR_SCENE"
	inf.ActionInfo.Scene.SceneId = sceneId
	return this.createQRCode(&inf)
}

// create temporary qrcode
func (this *Weixinmp) CreateQRLimitScene(expireSeconds, sceneId int64) (string, error) {
	var inf qrScene
	inf.ExpireSeconds = expireSeconds
	inf.ActionName = "QR_LIMIT_SCENE"
	inf.ActionInfo.Scene.SceneId = sceneId
	return this.createQRCode(&inf)
}

func (this *Weixinmp) createQRCode(inf *qrScene) (string, error) {
	data, err := json.Marshal(inf)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%sqrcode/create?access_token=", UrlPrefix)
	buf := bytes.NewBuffer(data)
	ticket := ""
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return "", err
		}
		rtn, err := post(url+token, "text/plain", buf)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return "", err
		}
		ticket = rtn.Ticket
		break // success
	}
	return ticket, nil
}

// download media to file
func (this *Weixinmp) DownloadMediaFile(mediaId, fileName string) error {
	url := fmt.Sprintf("%sget?media_id=%s&access_token=", MediaUrlPrefix, mediaId)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		resp, err := http.Get(url + token)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		// json
		if resp.Header.Get("Content-Type") == "text/plain" {
			var rtn response
			if err := json.Unmarshal(data, &rtn); err != nil {
				if i < retryNum-1 {
					continue
				}
				return err
			}
			if i < retryNum-1 {
				continue
			}
			return errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
		}
		// media
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		break // success
	}
	return nil
}

// upload media to file
func (this *Weixinmp) UploadMediaFile(mediaType, fileName string) (string, error) {
	var buf bytes.Buffer
	bw := multipart.NewWriter(&buf)
	defer bw.Close()
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fw, err := bw.CreateFormFile("filename", f.Name())
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(fw, f); err != nil {
		return "", err
	}
	f.Close()
	bw.Close()
	url := fmt.Sprintf("%supload?type=%s&access_token=", MediaUrlPrefix, mediaType)
	mime := bw.FormDataContentType()
	mediaId := ""
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return "", err
		}
		rtn, err := post(url+token, mime, &buf)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return "", err
		}
		mediaId = rtn.MediaId
		break // success
	}
	return mediaId, nil
}

type Button struct {
	Type      string   `json:"type,omitempty"`
	Name      string   `json:"name"`
	Key       string   `json:"key,omitempty"`
	Url       string   `json:"url,omitempty"`
	SubButton []Button `json:"sub_button,omitempty"`
}

// create custom menu
func (this *Weixinmp) CreateCustomMenu(btn *[]Button) error {
	var menu struct {
		Button *[]Button `json:"button"`
	}
	menu.Button = btn
	data, err := json.Marshal(&menu)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)
	url := fmt.Sprintf("%smenu/create?access_token=", UrlPrefix)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		if _, err := post(url+token, "text/plain", buf); err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		break // success
	}
	return nil
}

// get custom menu
func (this *Weixinmp) GetCustomMenu() ([]Button, error) {
	var menu struct {
		Menu struct {
			Button []Button `json:"button"`
		} `json:"menu"`
	}
	url := fmt.Sprintf("%smenu/get?access_token=", UrlPrefix)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return nil, err
		}
		resp, err := http.Get(url + token)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return nil, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return nil, err
		}
		// has error?
		var rtn response
		if err := json.Unmarshal(data, &rtn); err != nil {
			if i < retryNum-1 {
				continue
			}
			return nil, err
		}
		// yes
		if rtn.ErrCode != 0 {
			if i < retryNum-1 {
				continue
			}
			return nil, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
		}
		// no
		if err := json.Unmarshal(data, &menu); err != nil {
			if i < retryNum-1 {
				continue
			}
			return nil, err
		}
		break // success
	}
	return menu.Menu.Button, nil
}

// delete custom menu
func (this *Weixinmp) DeleteCustomMenu() error {
	url := UrlPrefix + "menu/delete?access_token="
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		if _, err := get(url + token); err != nil {
			if i < retryNum-1 {
				continue
			}
			return err
		}
		break // success
	}
	return nil
}

type UserInfo struct {
	Subscribe     int64  `json:"subscribe"`
	Openid        string `json:"openid"`
	Nickname      string `json:"nickname"`
	Sex           int64  `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	Headimgurl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
}

// get user info
func (this *Weixinmp) GetUserInfo(openId string) (UserInfo, error) {
	var uinf UserInfo
	url := fmt.Sprintf("%suser/info?lang=zh_CN&openid=%s&access_token=", UrlPrefix, openId)
	// retry
	for i := 0; i < retryNum; i++ {
		token, err := this.AccessToken.Fresh()
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		resp, err := http.Get(url + token)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		// has error?
		var rtn response
		if err := json.Unmarshal(data, &rtn); err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		// yes
		if rtn.ErrCode != 0 {
			if i < retryNum-1 {
				continue
			}
			return uinf, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
		}
		// no
		if err := json.Unmarshal(data, &uinf); err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		break // success
	}
	return uinf, nil
}


type Wxconf struct {
	AppId         string `json:"appId"`
	Timestamp     int    `json:"timestamp"`
	Noncestr      string `json:"nonceStr"`
	Signature     string `json:"signature"`
	JsApiList     []string `json:"jsApiList"`
}
var JsApiList =[]string{
	"onMenuShareTimeline","onMenuShareAppMessage","onMenuShareQQ",
	"onMenuShareWeibo","onMenuShareQZone","chooseImage",
	"uploadImage","downloadImage","startRecord","stopRecord",
	"onVoiceRecordEnd","playVoice","pauseVoice","stopVoice",
	"translateVoice","openLocation","getLocation","hideOptionMenu",
	"showOptionMenu","closeWindow","hideMenuItems","showMenuItems",
	"showAllNonBaseMenuItem","hideAllNonBaseMenuItem","scanQRCode",
	"chooseWXPay",
}
func (this *Weixinmp) GetSignature(w map[string]string) *Wxconf {
	s := []string{}
	for k,_:= range w{
		s = append(s,k)
	}
	sort.Strings(s)
	var signatureStr = ""
	for _,v:= range s{
		signatureStr+=v+"="+w[v]+"&"
	}
	k := strings.Trim(signatureStr, "&")
	h := sha1.New()
	io.WriteString(h,k)
	sign:= fmt.Sprintf("%x", h.Sum(nil))
	t,_:=strconv.Atoi(w["timestamp"])
	return &Wxconf{
		AppId:this.AccessToken.AppId,
		Timestamp:t,
		Noncestr:w["noncestr"],
		Signature: sign,
		JsApiList:[]string{},
	}
}

