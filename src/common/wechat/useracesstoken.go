package wechat

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

//用户授权登陆后返回的accesstoken,与基础的accesstoken有区别，不限制调用次数
type Useracesstoken struct {
	AppId             string
	Secret            string
	AccessToken       string
	Code              string
	Openid			  string
	RefreshToken	  string
}

const (
	OAUTH2Prefix = "https://api.weixin.qq.com/sns/oauth2/"
	UserinfoPrefix  = "https://api.weixin.qq.com/sns/userinfo"
)


func NewUseracesstoken(appid,secret,code string)*Useracesstoken{
	return &Useracesstoken{
		AppId:appid,
		Secret:secret,
		Code:code,
	}
}

type  UserRespose struct {
	AccessToken     string     `json:"access_token"`
	ExpiresIn       int        `json:"expires_in"`
	RefreshToken    string     `json:"refresh_token"`
	Openid           string		`json:"openid"`
	Scope            string		`json:"scope"`
	Errcode			 int64		`json:"errcode"`
	Errmsg			 string		`json:"errmsg"`
	WebUserInfo
}

type WebUserInfo struct {
	Openid        string `json:"openid"`
	Nickname      string `json:"nickname"`
	Sex           int64  `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	Headimgurl    string `json:"headimgurl"`
	Privilege  	  []string `json:"privilege"`
	Unionid		  string `json:"unionid"`
}

func (this *Useracesstoken) Fetch() (rtn *UserRespose,err error) {
	url := fmt.Sprintf("%saccess_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", OAUTH2Prefix,this.AppId,this.Secret,this.Code)
	rtn, err = acesstokenget(url)
	if err != nil {
		return  nil, err
	}
	if rtn.Errcode>0{
		return  nil,errors.New("code 已失效或已使用!")
	}
	return rtn, nil
}

func  (this *Useracesstoken) GetUserinfo()(rtn *UserRespose,err error){
	url := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", UserinfoPrefix,this.AccessToken,this.Openid)
	rtn, err = acesstokenget(url)
	if err != nil {
		return  nil, err
	}
	if rtn.Errcode>0{
		return  nil,errors.New("获取用户信息失败!")
	}
	return rtn, nil
}


func acesstokenget(url string) (*UserRespose, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rtn UserRespose
	if err := json.Unmarshal(data, &rtn); err != nil {
		return nil, err
	}
	return &rtn, nil
}
