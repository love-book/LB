package rabbitmq

import (
	"fmt"
)


func Register(Consum *Delivery) {
	Consum.HandlerRegister.Add(555, Handler(appMsgText), "appMsgText")
	if err := Consum.HandlerRegister.EnableByName("appMsgText"); err != nil {
		fmt.Println(err)
	}
}

// 微信文字消息
func appMsgText(Consum *Delivery) {
	fmt.Println(Consum,"微信文字消息")
}
