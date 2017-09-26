package rabbitmq

import (
	"github.com/astaxie/beego"
)

var (
	mqhost  = beego.AppConfig.String("mqhost")
	mqport  = beego.AppConfig.String("mqport")
	mquser  = beego.AppConfig.String("mquser")
	mqpass  = beego.AppConfig.String("mqpass")
	mqvhost = beego.AppConfig.String("mqvhost")
	uri = "amqp://"+mquser+":"+mqpass+"@"+mqhost+":"+mqport+"/"+mqvhost
)

func Push(key string,body string){
	Exchange:=beego.AppConfig.String(key+"::Exchange")
	ExchangeType:=beego.AppConfig.String(key+"::ExchangeType")
	Producer:= Producer{
		Exchange,
		ExchangeType,
		Exchange,
		Exchange,
		body,
		true}
	Producer.Push()
}

func init(){
	subkasoly()
	subgo()
}

func subkasoly(){
	Exchange:=beego.AppConfig.String("kasoly::Exchange")
	ExchangeType:=beego.AppConfig.String("kasoly::ExchangeType")
	Consum:= Consum{
		Exchange,
		ExchangeType,
		Exchange,
		Exchange,
		Exchange}
	Consum.Sub()
}

func subgo(){
	Exchange:=beego.AppConfig.String("go::Exchange")
	ExchangeType:=beego.AppConfig.String("go::ExchangeType")
	Consum:= Consum{
		Exchange,
		ExchangeType,
		Exchange,
		Exchange,
		Exchange}
	Consum.Sub()
}