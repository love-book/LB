package conndatabase

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type mysqlConn struct {
	dbhost		string
	dbport		string
	dbuser		string
	dbpassword	string
	dbname		string
	dbcharset	string
	dbconnstring	string

}


func init() {
	var conn *mysqlConn  = new(mysqlConn)
	conn.dbhost = beego.AppConfig.String("dbhost")
	conn.dbport = beego.AppConfig.String("dbport")
	conn.dbuser = beego.AppConfig.String("dbuser")
	conn.dbpassword = beego.AppConfig.String("dbpassword")
	conn.dbname = beego.AppConfig.String("dbname")
	conn.dbcharset = beego.AppConfig.String("dbset")

	conn.dbconnstring= conn.dbuser + ":" +conn.dbpassword + "@tcp(" + conn.dbhost + ":" + conn.dbport + ")/" + conn.dbname + "?charset="+conn.dbcharset
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",conn.dbconnstring,100,20)
	orm.Debug = true
	//orm.SetMaxOpenConns("default",100)
	//orm.SetMaxIdleConns("default",20)
/*	orm.RegisterModel(new (models.User))
	orm.RegisterModel(new (models.BookInfo))*/

}


func Insert(o interface{}) error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func Read(o interface{}) error {
	if err := orm.NewOrm().Read(o); err != nil {
		return err
	}
	return nil
}

func Update(o interface{}) error {
	if _, err := orm.NewOrm().Update(o); err != nil {
		return err
	}
	return nil
}

func Delete(o interface{}) error {
	if _, err := orm.NewOrm().Delete(o); err != nil {
		return err
	}
	return nil
}

func Query(sql string ,obj interface{},args ...interface{}) (int64,error) {
	num, err := orm.NewOrm().Raw(sql,args).QueryRows(&obj)
	if err == nil {
		return num,err
	}
	return 0, err
}


func  RawSeter(sql string ,args ...interface{}) orm.RawSeter{
	var rs orm.RawSeter
	rs = orm.NewOrm().Raw(sql,args)
	return  rs
}

func FileUrl() string{
	FilesUrl := beego.AppConfig.String("FilesUrl")
	return FilesUrl
}
