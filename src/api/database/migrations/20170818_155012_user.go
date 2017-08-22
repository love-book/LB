package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20170818_155012 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20170818_155012{}
	m.Created = "20170818_155012"
	migration.Register("User_20170818_155012", m)
}

// Run the migrations
func (m *User_20170818_155012) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `users` (`userid` bigint(20) NOT NULL COMMENT '用户id',`openid` char(20) NOT NULL DEFAULT '' COMMENT 'openid',`wnickname` char(30) NOT NULL DEFAULT '' COMMENT '微信昵称',`wimgurl` char(150) NOT NULL DEFAULT '' COMMENT '微信头像',`nickname` char(30) NOT NULL DEFAULT '' COMMENT '用户名',`imgurl` char(100) NOT NULL DEFAULT '' COMMENT '头像',`gender` char(1) NOT NULL DEFAULT '0' COMMENT '性别',`age` char(5) NOT NULL DEFAULT '0' COMMENT '年龄',`telphone` char(11) NOT NULL DEFAULT '0' COMMENT '电话',`qq` char(20) NOT NULL DEFAULT '' COMMENT 'QQ号',`weino` char(20) NOT NULL DEFAULT '' COMMENT '微博号',`signature` char(150) NOT NULL DEFAULT '' COMMENT '签名',`address` char(255) NOT NULL DEFAULT '' COMMENT '地址',`created_at` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',`updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',PRIMARY KEY (`userid`)) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';")
}

// Reverse the migrations
func (m *User_20170818_155012) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
