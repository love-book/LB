-- ----------------------------
-- Table structure for lb_booknews
-- ----------------------------
DROP TABLE IF EXISTS `lb_booknews`;
CREATE TABLE `lb_booknews` (
  `newid` bigint(20) NOT NULL COMMENT '消息编号',
  `userid` int(20) NOT NULL DEFAULT '0' COMMENT '用户编号',
  `from_to` tinyint(1) NOT NULL DEFAULT '0' COMMENT '消息人类型 from:1书属于的人,to:2借书人',
  `books` json DEFAULT NULL COMMENT 'bookid:图书编号,bookname:书名,author:作者,imgurl:书本组图,describe:简介,state:图书状态,flag:图书标签,create_time:上架时间',
  `users` json DEFAULT NULL COMMENT 'from:书主人信息,to:借阅者信息,nickname:昵称,gender:性别,yimage_url:用户头像',
  `order_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态1:同意2:拒绝,3:完成，0：借书请求',
  `pushtime` int(11) NOT NULL DEFAULT '0' COMMENT '发送时间',
  PRIMARY KEY (`newid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息表';

-- ----------------------------
-- Records of lb_booknews
-- ----------------------------

-- ----------------------------
-- Table structure for lb_bookorder
-- ----------------------------
DROP TABLE IF EXISTS `lb_bookorder`;
CREATE TABLE `lb_bookorder` (
  `orderid` bigint(20) NOT NULL COMMENT '订单编号',
  `userid_from` int(11) NOT NULL DEFAULT '0' COMMENT '书主人编号',
  `userid_to` int(11) NOT NULL DEFAULT '0' COMMENT '借书人编号',
  `books` json DEFAULT NULL COMMENT 'bookid:图书编号,bookname:书名,author:作者,imgurl:书本组图,describe:简介,state:图书状态,flag:图书标签,create_time:上架时间',
  `users` json DEFAULT NULL COMMENT 'from:书主人信息,to:借阅者信息',
  `order_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态1:接收2:拒绝,3:完成',
  `pushtime` int(11) NOT NULL DEFAULT '0' COMMENT '发送时间',
  PRIMARY KEY (`orderid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';

-- ----------------------------
-- Records of lb_bookorder
-- ----------------------------

-- ----------------------------
-- Table structure for lb_bookrack
-- ----------------------------
DROP TABLE IF EXISTS `lb_bookrack`;
CREATE TABLE `lb_bookrack` (
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `books` json DEFAULT NULL COMMENT 'bookid:图书编号,bookname:书名,author:作者,imgurl:书本组图,describe:简介,state:图书状态,flag:图书标签,create_time:上架时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='书架表';

-- ----------------------------
-- Records of lb_bookrack
-- ----------------------------

-- ----------------------------
-- Table structure for lb_books
-- ----------------------------
DROP TABLE IF EXISTS `lb_books`;
CREATE TABLE `lb_books` (
  `bookid` bigint(20) NOT NULL COMMENT '图书编号',
  `bookname` char(50) NOT NULL DEFAULT '' COMMENT '书名',
  `author` char(20) NOT NULL DEFAULT '' COMMENT '作者',
  `imageurl` char(150) NOT NULL DEFAULT '' COMMENT '图书封面图',
  `imagehead` char(150) NOT NULL DEFAULT '' COMMENT '图书正面图',
  `imaegback` char(150) NOT NULL DEFAULT '' COMMENT '图书背面图',
  `isbn` char(18) NOT NULL DEFAULT '',
  `depreciation` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '图书折旧',
  `price` int(8) NOT NULL DEFAULT '0' COMMENT '标价',
  `describe` text COMMENT '图书简介',
  `state` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0非锁定状态 1：锁定状态',
  PRIMARY KEY (`bookid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图书表';

-- ----------------------------
-- Records of lb_books
-- ----------------------------

-- ----------------------------
-- Table structure for lb_concern
-- ----------------------------
DROP TABLE IF EXISTS `lb_concern`;
CREATE TABLE `lb_concern` (
  `userid` bigint(20) NOT NULL,
  `books` json DEFAULT NULL,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='收藏表';

-- ----------------------------
-- Records of lb_concern
-- ----------------------------

-- ----------------------------
-- Table structure for lb_loginlog
-- ----------------------------
DROP TABLE IF EXISTS `lb_loginlog`;
CREATE TABLE `lb_loginlog` (
  `userid` bigint(20) NOT NULL DEFAULT '0',
  `logintime` int(11) NOT NULL DEFAULT '0' COMMENT '登录时间',
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录时间表';

-- ----------------------------
-- Records of lb_loginlog
-- ----------------------------

-- ----------------------------
-- Table structure for lb_users
-- ----------------------------
DROP TABLE IF EXISTS `lb_users`;
CREATE TABLE `lb_users` (
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `openid` char(20) NOT NULL DEFAULT '' COMMENT 'openid',
  `wnickname` char(30) NOT NULL DEFAULT '' COMMENT '微信昵称',
  `wimgurl` char(150) NOT NULL DEFAULT '' COMMENT '微信头像',
  `nickname` char(30) NOT NULL DEFAULT '' COMMENT '用户名',
  `imgurl` char(100) NOT NULL DEFAULT '' COMMENT '头像',
  `regtime` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `gender` char(1) NOT NULL DEFAULT '0' COMMENT '性别',
  `age` char(5) NOT NULL DEFAULT '0' COMMENT '年龄',
  `telphone` int(11) DEFAULT '0' COMMENT '电话',
  `qq` char(20) NOT NULL DEFAULT '' COMMENT 'QQ号',
  `weino` char(20) NOT NULL DEFAULT '' COMMENT '微博号',
  `signature` char(150) NOT NULL DEFAULT '' COMMENT '签名',
  `address` char(255) NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of lb_users
-- ----------------------------
