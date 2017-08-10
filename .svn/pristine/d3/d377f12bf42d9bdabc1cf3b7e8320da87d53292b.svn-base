package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
	"strconv"
	"app/common/conndatabase"
)


func init() {
	/*dbhost := beego.AppConfig.String("dbhost")
dbport := beego.AppConfig.String("dbport")
dbuser := beego.AppConfig.String("dbuser")
dbpassword := beego.AppConfig.String("dbpassword")
dbname := beego.AppConfig.String("dbname")
dbprefix := beego.AppConfig.String("dbprefix")
if dbport == "" {
	dbport = "3306"
}
dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
orm.RegisterDataBase("default", "mysql", dburl,30,30)
/*
beego必须个别名为default的数据库，作为默认使用。
orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)
第一个参数是数据库的别名，用来切换数据库使用。
第二个是driverName，在RegisterDriver时注册的
第三是数据库连接字符串:test:123456@/test?charset=utf8相对于用户名:密码@数据库地址+名称?字符集
第四个参数相当于:
orm.SetMaxIdleConns("default", 30)
设置数据库的最大空闲连接。
第五个参数相当于：
orm.SetMaxOpenConns("default", 30)
设置数据库的最大数据库连接。
第四个参数和第五个参数也可以不传值，会使用数据库默认值

	orm.RegisterModelWithPrefix(dbprefix,new(User))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
		// 自动建表
		//orm.RunSyncdb("default", false, true)
	}
  */
}


func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

//返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}


/**
 * 分页函数，适用任何表
 * 返回 总记录条数,总页数,以及当前请求的数据RawSeter,调用中需要"rs.QueryRows(&tblog)"就行了  --tblog是一个Tb_log对象
 * 参数：表名，当前页数，页面大小，条件（查询条件,格式为 " and name='zhifeiya' and age=12 "）
 *
	start, _ := this.GetInt("start") //获取起始位置
	length, _ := this.GetInt("length") //获取分页步长
	draw, _ := this.GetInt("draw") //获取请求次数
	var user []models.User
	var conditions string = " order by userid desc"
	var  TableName = "lb_users"
	totalItem, rs :=models.GetPagesInfo(TableName,start,length,conditions)
	rs.QueryRows(&user)
	Json := map[string]interface{}{"draw":draw,"recordsTotal": totalItem,"recordsFiltered":totalItem,"data":user}
	this.renderJson(Json)

 */
func GetPagesInfo(tableName string, start int, pagesize int, conditions string) (int, orm.RawSeter) {
	if start <= 1 {
		start = 1
	}
	if pagesize == 0 {
		pagesize = 15
	}
	var rs orm.RawSeter
	o := orm.NewOrm()
	var totalItem  int = 0                                                          //总条数
	o.Raw("SELECT count(*) FROM " + tableName + "  where true " + conditions).QueryRow(&totalItem) //获取总条数
	rs = o.Raw("select *  from  " + tableName + "  where true " + conditions + " LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(pagesize))
	return totalItem,  rs
}

//获取ID主键
func GetID() int64{
	IdWorker,err := conndatabase.NewIdWorker(1,1)
	if err != nil{
		fmt.Println(nil)
	}
	id,error :=  IdWorker.NextId()
	if error != nil{
		fmt.Println(nil)
	}
/*	if id < 0{
		id = 0-id
	}*/
	return id
}

