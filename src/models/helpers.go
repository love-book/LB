package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	"time"
)

type Helper struct {
	Id                int		`orm:"size(10);" form:"Id"`
	Sid               int		`orm:"size(10);" form:"Sid"`
	Channel			  int 	    `orm:"size(10);" form:"Channel"`
	Uin               int    	`orm:"size(64);" form:"Uin" valid:"Required"`
	UserName          string  	`orm:"size(255);" form:"UserName" valid:"Required"`
	NickName          string	`orm:"size(255);" form:"NickName" valid:"Required"`
	HeadImgUrl        string	`orm:"size(255);" form:"HeadImgUrl"`
	ContactFlag       int		`orm:"size(10);"  form:"ContactFlag"`
	MemberCount       int		`orm:"size(20);"  form:"MemberCount"`
	MemberList        string	`orm:"size(255);" form:"MemberList"`
	RemarkName        string	`orm:"size(255);" form:"RemarkName" `
	PYInitial         string	`orm:"size(255);" form:"PYInitial" `
	PYQuanPin         string	`orm:"size(255);" form:"PYQuanPin" `
	RemarkPYInitial   string	`orm:"size(255);" form:"RemarkPYInitial" `
	RemarkPYQuanPin   string	`orm:"size(255);" form:"RemarkPYQuanPin" `
	HideInputBarFlag  int		`orm:"size(64);" form:"HideInputBarFlag"`
	StarFriend        int		`orm:"size(10);" form:"StarFriend"`
	Sex               int		`orm:"size(2);"  form:"Sex" `
	Signature         string	`orm:"size(255);" form:"Signature"`
	AppAccountFlag    int		`orm:"size(6);"  form:"AppAccountFlag"`
	Statues           int		`orm:"size(2);"  form:"Statues"`
	AttrStatus        uint32	`orm:"size(64);" form:"AttrStatus"`
	Province          string	`orm:"size(64);" form:"Province"`
	City              string	`orm:"size(64);" form:"City"`
	Alias             string	`orm:"size(64);" form:"Alias"`
	VerifyFlag        int		`orm:"size(10);" form:"VerifyFlag"`
	OwnerUin          int		`orm:"size(64);" form:"OwnerUin"`
	WebWxPluginSwitch int		`orm:"size(64);" form:"WebWxPluginSwitch"`
	HeadImgFlag       int		`orm:"size(64);" form:"HeadImgFlag"`
	SnsFlag           int		`orm:"size(64);" form:"SnsFlag"`
	UniFriend         int		`orm:"size(64);" form:"UniFriend"`
	DisplayName       string	`orm:"size(64);" form:"DisplayName"`
	ChatRoomId        int		`orm:"size(64);" form:"ChatRoomId"`
	KeyWord           string	`orm:"size(255);" form:"KeyWord"`
	EncryChatRoomId   string	`orm:"size(255);" form:"EncryChatRoomId"`
	IsOwner           int		`orm:"size(2);" form:"IsOwner"`
	MemberStatus      int		`orm:"size(2);" form:"MemberStatus"`
	CreatedAt         int64		`orm:"size(10);" form:"CreatedAt"`
	UpdatedAt         int64		`orm:"size(10);" form:"UpdatedAt"`
}

func (h *Helper) TableName() string {
	return beego.AppConfig.String("table_helper")
}

func init() {
	orm.RegisterModel(new(Helper))
}


func checkHelper(h *Helper) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&h)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//get Helper list
func GetHelperlist(page int64, page_size int64, sort string) (helpers []orm.Params, count int64) {
	o := orm.NewOrm()
	helper := new(Helper)
	qs := o.QueryTable(helper)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&helpers)
	count, _ = qs.Count()
	return helpers, count
}

func AddHelper(h *Helper) (int64, error) {
	if err := checkHelper(h); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	helper := new(Helper)
	helper.Sid = 0
	helper.Uin  = h.Uin
	helper.UserName = h.UserName
	helper.NickName = h.NickName
	helper.Channel  = h.Channel
	helper.CreatedAt = time.Now().Unix()
	helper.UpdatedAt = 0
	id, err := o.Insert(helper)
	return id, err
}

func UpdateHelper(h *Helper) (int64, error) {
	if err := checkHelper(h); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	helper := make(orm.Params)
	if h.CreatedAt != 0 {
		helper["CreatedAt"] = h.CreatedAt
	}
	if h.UpdatedAt != 0 {
		helper["UpdatedAt"] = h.UpdatedAt
	}
	if len(helper) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Helper
	num, err := o.QueryTable(table).Filter("Uin", h.Uin).Update(helper)
	return num, err
}

func DelHelperByName(uin int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Helper{Uin: uin})
	return status, err
}

func HelperList() (helpers []orm.Params) {
	o := orm.NewOrm()
	helper := new(Helper)
	qs := o.QueryTable(helper)
	qs.Values(&helpers, "name", "data")
	return helpers
}

func GetHelperByUserName(conditions string) (helper *Helper,err error) {
	table := beego.AppConfig.String("table_helper")
	rowSql  := "select * from  "+table+"  where true "+conditions+" limit 1"
	o := orm.NewOrm()
	err = o.Raw(rowSql).QueryRow(&helper)
	return
}