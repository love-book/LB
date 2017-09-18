package models

import (
	"fmt"
	"encoding/json"
	"time"
)

var(
  	BotMsgOne    =  make(chan []byte,500)
  	BotMsgTwo    =  make(chan []byte,500)
  	BotMsgThree  =  make(chan []byte,500)
  	BotMsgFour   =  make(chan []byte,500)
  	BotMsgFive   =  make(chan []byte,500)
  	BotMsgSex    =  make(chan []byte,500)
  	BotMsgSeven  =  make(chan []byte,500)
	BotMsgEigth  =  make(chan []byte,500)
  	BotMsgNine   =  make(chan []byte,500)
    BotMsgTen    =  make(chan []byte,500)
)

type Channel struct {
	Text       string
	Img        string
    Verify     string
    Action     string
	MsgId      string
	RemarkName string
	Bin        string
	MsgType    int
	Channel    int
}

func (c *Channel) ChannelOne()(ok bool){
   loop1:
	for {
		select {
		case m := <-BotMsgOne:
			fmt.Println("BotMsgOne")
			go Consumer(m)
			ok  = true
			break loop1
		case <- time.After(15 * time.Second):
			ok  = false
			println("timeout-BotMsgOne")
			break loop1
		}
	}
	return
}
func (c *Channel) ChannelTwo()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgTwo:
			fmt.Println("BotMsgTwo")
			ok  = true
			go Consumer(m)
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgTwo")
			ok  = false
			break loop1
		}
	}
	return
}
func  (c *Channel) ChannelThree()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgThree:
			fmt.Println("BotMsgThree")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgThree")
			ok  = false
			break loop1
		}
	}
	return
}
func (c *Channel) ChannelFour()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgFour:
			fmt.Println("BotMsgFour")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgFour")
			break loop1
		}
	}
	return
}
func (c *Channel)ChannelFive()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgFive:
			fmt.Println("BotMsgFive")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgFive")
			ok = false
			break loop1
		}
	}
	return
}
func (c *Channel)ChannelSex()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgSex:
			fmt.Println("BotMsgSex")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgSex")
			break loop1
		}
	}
	return
}
func (c *Channel) ChannelSeven()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgSeven:
			fmt.Println("BotMsgSeven")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgSeven")
			break loop1
		}
	}
	return
}
func (c *Channel)ChannelEigth()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgEigth:
			fmt.Println("BotMsgEigth")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgEigth")
			break loop1
		}
	}
	return
}
func (c *Channel)ChannelNine()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgNine:
			fmt.Println("BotMsgNine")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgNine")
			break loop1
		}
	}
	return
}
func (c *Channel)ChannelTen()(ok bool){
loop1:
	for {
		select {
		case m := <-BotMsgTen:
			fmt.Println("BotMsgTen")
			go Consumer(m)
			ok = true
			break loop1
		case <- time.After(15 * time.Second):
			println("timeout-BotMsgTen")
			break loop1
		}
	}
	return
}


func Consumer(msg []byte)(channel *Channel){
	fmt.Println(string(msg))
	return
}

func Producter(msg chan []byte) {
	Js := map[string]interface{}{"code": 200, "msg": "成功!", "data": BotConfAndAction()}
	body, err := json.Marshal(Js)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	msg <- body
}

//机器人配置信息和机器人动作
func BotConfAndAction()(msg []map[string]interface{}){
	conf := map[string]interface{}{
		"Text":"感谢关注lb!",
		"Img":"http://n.sinaimg.cn/default/4_img/uplaod/3933d981/20170908/i9ew-fykuffc4351895.jpg",
		"Verify":"已添加小助手为好友!请按操作继续完成绑定信息!",
		"Action":"action",
		"MsgId":"",
		"RemarkName":"A1505032200",
		"Bin":"",
		"MsgType":1990,
	}
	msg = append(msg,conf)
	return msg
}
