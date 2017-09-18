package models

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"common/conndatabase"
	"time"
)

/*
func init() {
dbhost := beego.AppConfig.String("dbhost")
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
}*/


func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}



//获取ID主键
func GetID() int64{
	time.Sleep(1 * time.Nanosecond)
	IdWorker,err := conndatabase.NewIdWorker(2,1)
	if err != nil{
		fmt.Println(nil)
	}
	id,error :=  IdWorker.NextId()
	if error != nil{
		fmt.Println(nil)
	}
	if id < 0{
		id = 0-id
	}
	return id
}
