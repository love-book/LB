package models

import(
	_ "common/conndatabase"
	"github.com/astaxie/beego/validation"
	"log"
	"github.com/astaxie/beego/orm"
)

type  Users struct {
	Userid		string	`json:"userid" orm:"pk;size(20);column(userid);"`
	Openid  	string	`json:"openid"`
	Wnickname  	string	`json:"wnickname"`
	Wimgurl    	string 	`json:"wimgurl"`
	Nickname  	string 	`json:"nickname"`
	Imgurl    	string 	`json:"imgurl" `
	Gender  	int8  	`json:"gender"`
	Age   		int32 	`json:"age"`
	Telphone  	int32  	`json:"telphone"`
	Qq  		string  `json:"qq"`
	Weino  		string  `json:"weibo"`
	Signature  	string  `json:"signature"`
	Address  	string  `json:"address"`
	Created_at  int64  	`json:"created_at"`
	Updated_at  int64  	`json:"updated_at"`
}

func init()  {
	orm.RegisterModelWithPrefix("lb_",new(Users))
}

//添加用户验证
func (a Users) InsertValidation() error {
	valid := validation.Validation{}
	valid.Required(a.Userid,  "userid").Message("用户编号不能为空！")
	valid.Required(a.Nickname, "nickname").Message("用户昵称不能为空！")
	valid.MaxSize(a.Nickname,50,"nickname").Message("用户昵称不大于50个字符！")
	valid.MinSize(a.Nickname,5,"nickname").Message("用户昵称不小于5个字符！")
	//valid.Range(a.Age, 0, 100, "age").Message("年龄不符合范围！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return err
		}
	}
	return nil
}
//验证userid
func (a Users) UserValidation() error {
	valid := validation.Validation{}
	valid.Required(a.Userid,  "userid").Message("用户编号不能为空！")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return err
		}
	}
	return nil
}