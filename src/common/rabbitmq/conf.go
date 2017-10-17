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
	Exchange := key
	ExchangeType := ExchangeMap[key]
	Producer:= Producer{
		Exchange,
		ExchangeType,
		Exchange,
		Exchange,
		body,
		true,
	}
	Producer.Push()
}



func init(){
	actionConcernNotice()
	actionLibraryrequest()
	actionAgreelibraryrequest()
}

//mq.Push(mq.ConcernNotice,"mykasoly")
const(
	ConcernNotice = "concernNotice" //收藏
	LibraryRequest = "libraryrequest" //发起借书
	Agreelibraryrequest = "agreelibraryrequest" //同意借书
)
//mq消息类型
var ExchangeMap = map[string]string{
	//KASOLY:"x-delayed-message",
	ConcernNotice:"direct",
	LibraryRequest:"direct",
	Agreelibraryrequest:"direct",
}


func actionConcernNotice(){
	Exchange:= ConcernNotice
	ExchangeType:= ExchangeMap[ConcernNotice]
	consumer:= &Consumer{
		Exchange:Exchange,
		ExchangeType:ExchangeType,
		Queue :Exchange,
		BindingKey :Exchange,
		ConsumerTag :Exchange,
		HandlerTag :Exchange,
		HandlerRegister:CreateHandlerRegister(),
	}
	ConcernNoticeRegister(consumer)
	consumer.Sub()
}


func actionLibraryrequest(){
	Exchange:= LibraryRequest
	ExchangeType:= ExchangeMap[LibraryRequest]
	consumer:= &Consumer{
		Exchange:Exchange,
		ExchangeType:ExchangeType,
		Queue :Exchange,
		BindingKey :Exchange,
		ConsumerTag :Exchange,
		HandlerTag :Exchange,
		HandlerRegister:CreateHandlerRegister(),
	}
	LibraryrequestRegister(consumer)
	consumer.Sub()
}


func actionAgreelibraryrequest(){
	Exchange:= Agreelibraryrequest
	ExchangeType:= ExchangeMap[Agreelibraryrequest]
	consumer:= &Consumer{
		Exchange:Exchange,
		ExchangeType:ExchangeType,
		Queue :Exchange,
		BindingKey :Exchange,
		ConsumerTag :Exchange,
		HandlerTag :Exchange,
		HandlerRegister:CreateHandlerRegister(),
	}
	AgreelibraryrequestRegister(consumer)
	consumer.Sub()
}