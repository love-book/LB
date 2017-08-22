package models

import(
	_ "common/conndatabase"
	"github.com/astaxie/beego/validation"
	"log"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"errors"
	"reflect"
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
	Telphone  	string  `json:"telphone"`
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

//根据id查询多个用户
func GetUsersByIds(ids []string)(u []*Users,err error){
	sql:= "select * from lb_users  where userid in(? ,?)"
	RawSeter := orm.NewOrm().Raw(sql,ids)
	num,err := RawSeter.QueryRows(&u)
	if err == nil{
		fmt.Println("受影响的行数",num)
		return u,nil
	}
	return nil,err
}


// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Userid: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Users
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Userid: m.Userid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id string) (err error) {
	o := orm.NewOrm()
	v := Users{Userid: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Userid: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
