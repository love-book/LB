package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
	"log"
	"fmt"
	"strings"
	"errors"
	"reflect"
	comm "common/conndatabase"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type  Users struct {
	Userid		string	 `json:"userid" orm:"pk;size(20);column(userid);"`
	Openid  	string	 `json:"openid"`
	Wnickname  	string	 `json:"wnickname"`
	Wimgurl    	string 	 `json:"wimgurl"`
	Nickname  	string 	 `json:"nickname"`
	Imgurl    	string 	 `json:"imgurl" `
	Gender  	int64  	 `json:"gender"`
	Age   		int64 	 `json:"age"`
	Telphone  	string   `json:"telphone"`
	Password	string   `json:"password"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
	Weino  		string   `json:"weibo"`
	Signature  	string   `json:"signature"`
	Constellation string `json:"constellation"`
	Province  	string   `json:"province"`
	City  	    string   `json:"city"`
	Address  	string   `json:"address"`
	Long  	    float64  `json:"long"`
	Lat  	    float64  `json:"lat"`
	Logintime	int64  	 `json:"logintime"`
	Created_at  int64  	 `json:"created_at"`
	Updated_at  int64  	 `json:"updated_at"`
}

type  UsersList struct {
	Userid		string	 `json:"userid"`
	Openid  	string	 `json:"openid"`
	Wnickname  	string	 `json:"wnickname"`
	Wimgurl    	string 	 `json:"wimgurl"`
	Nickname  	string 	 `json:"nickname"`
	Imgurl    	string 	 `json:"imgurl" `
	Gender  	int64  	 `json:"gender"`
	Age   		int64 	 `json:"age"`
	Telphone  	string   `json:"telphone"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
	Weino  		string   `json:"weibo"`
	Signature  	string   `json:"signature"`
	Constellation string `json:"constellation"`
	Province  	string   `json:"province"`
	City  	    string   `json:"city"`
	Address  	string   `json:"address"`
	Long  	    float64  `json:"long"`
	Lat  	    float64  `json:"lat"`
	Logintime	int64  	 `json:"logintime"`
	Radius     	string	 `json:"radius"`
}

type  UserInfo struct {
	Userid		string	 `json:"userid"`
	Openid  	string	 `json:"openid"`
	Wnickname  	string	 `json:"wnickname"`
	Wimgurl    	string 	 `json:"wimgurl"`
	Nickname  	string 	 `json:"nickname"`
	Imgurl    	string 	 `json:"imgurl" `
	Gender  	int64  	 `json:"gender"`
	Age   		int64 	 `json:"age"`
	Telphone  	string   `json:"telphone"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
	Weino  		string   `json:"weibo"`
	Signature  	string   `json:"signature"`
	Constellation string `json:"constellation"`
	Province  	string   `json:"province"`
	City  	    string   `json:"city"`
	Address  	string   `json:"address"`
	Long  	    float64  `json:"long"`
	Lat  	    float64  `json:"lat"`
	Logintime	int64  	 `json:"logintime"`
	Created_at  int64  	 `json:"created_at"`
	Updated_at  int64  	 `json:"updated_at"`
	Msgtips_num int64 	 `json:"msgtips_num"`
}

func init()  {
	orm.RegisterModel(new(Users))
}

func (c *Users) TableName() string {
	return beego.AppConfig.String("table_users")
}


//添加用户验证
func (a Users) InsertValidation() error {
	valid := validation.Validation{}
	valid.Required(a.Userid,  "userid").Message("用户编号不能为空！")
	valid.Required(a.Nickname, "nickname").Message("用户昵称不能为空！")
	valid.MaxSize(a.Nickname,30,"nickname").Message("用户昵称不大于30个字符！")
	valid.MinSize(a.Nickname,1,"nickname").Message("用户昵称不小于2个字符！")
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


//根据OpenId查询用户
func GetUsersByOpenId(id []string)(u *Users,err error){
	sql:= "select * from lb_users  where openid =? limit 1"
	RawSeter := orm.NewOrm().Raw(sql,id)
	err = RawSeter.QueryRow(&u)
	return
}

//查询用户
func GetUsersBypass(pram []string)(u *Users,err error){
	sql:= "select * from lb_users  where telphone =? and password =? limit 1"
	RawSeter := orm.NewOrm().Raw(sql,pram)
	err = RawSeter.QueryRow(&u)
	return
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

//根据openid获取用户经纬度
func  GetLocationByID(openid string)(l map[string]interface{},err error){
	rc := comm.Pool.Get()
	defer rc.Close()
	v,err := redis.Positions(rc.Do("GEOPOS",comm.LocationGeo,openid))
	l = make(map[string]interface{},len(v))
	for _,vs := range v{
		if vs != nil{
			l["long"] = vs[0]
			l["lat"]  = vs[1]
		}
		break
	}
	return
}


//根据openid获取附近的人
func  GetUsersByLocaltion(openid string,geokey string,radius int64,count int)(re []map[string]string,err error){
	rc := comm.Pool.Get()
	defer rc.Close()
	re,err = RadiusByMember(rc.Do("GEORADIUSBYMEMBER",geokey,openid,radius,"m","WITHDIST","ASC","COUNT",count))
	return
}

func RadiusByMember(result interface{}, err error) ([]map[string]string,error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return  nil,err
	}
	radiusMap := make([]map[string]string, len(values))
	for i := range values {
		if values[i] == nil {
			continue
		}
		p, ok := values[i].([]interface{})
		if !ok {
			return nil,fmt.Errorf("redigo: unexpected element type for interface slice, got type %T", values[i])
		}
		member, err := redis.String(p[0], nil)
		if err != nil {
			return nil, err
		}
		radius, err := redis.Float64(p[1], nil)
		if err != nil {
			return nil, err
		}
		res := make(map[string]string,len(p))
		res["member"] = member
		res["radius"] = strconv.FormatFloat(radius,'f', -1, 64)
		radiusMap[i] = res
	}
	return radiusMap,nil
}

//用户上传地理位置
func  AddLocationByID(openid ,LocationGeo string,lang float64,lat float64)(err error){
	rc := comm.Pool.Get()
	defer rc.Close()
	v,err := rc.Do("GEOADD",LocationGeo,lang,lat,openid)
	fmt.Println(v)
	return
}


//计算两个用户之间的距离
func  GetUsersRadiusByMembers(geokey string,member1 string,member2 string)(radius string,err error){
	rc := comm.Pool.Get()
	defer rc.Close()
	radius = GeoByMember(rc.Do("GEODIST",geokey,member1,member2,"m"))
	return
}
func GeoByMember(result interface{}, err error) (radius string) {
	values, err := redis.Bytes(result, err)
	if err != nil {
		return "0"
	}
	radius = string(values)
	return
}



func GetUserList(start int,length int,conditions string) (user []*UsersList,totalItem int){
	var  countSql = "select count(*) from lb_users  where true "+conditions
	var  rowsSql  = "select * from lb_users  where true "+conditions+"  order by logintime desc  limit " + strconv.Itoa(start) + "," + strconv.Itoa(length)
	o := orm.NewOrm()
	o.Raw(countSql).QueryRow(&totalItem) //获取总条数
	o.Raw(rowsSql).QueryRows(&user)
	return  user,totalItem
}