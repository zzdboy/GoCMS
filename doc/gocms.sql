-- phpMyAdmin SQL Dump
-- version 4.0.8
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2014-05-26 20:29:52
-- 服务器版本: 5.7.1-m11
-- PHP 版本: 5.5.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- 数据库: `gocms`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE IF NOT EXISTS `admin` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(32) DEFAULT NULL COMMENT '密码',
  `roleid` smallint(5) DEFAULT '0' COMMENT '角色',
  `lastloginip` varchar(15) DEFAULT '0.0.0.0' COMMENT '最后登陆地址PI',
  `lastlogintime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登陆时间',
  `email` varchar(40) DEFAULT NULL COMMENT '邮箱',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `lang` varchar(6) NOT NULL DEFAULT 'zh-cn' COMMENT '语言',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1:允许登陆 0:禁止登陆 ',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`,`roleid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='管理员表' AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`id`, `username`, `password`, `roleid`, `lastloginip`, `lastlogintime`, `email`, `realname`, `lang`, `status`, `createtime`) VALUES
(1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 1, '192.168.1.4', '2014-05-24 10:22:22', 'zzdboy1616@163.com', 'admin', 'zh-cn', 1, '2014-01-17 23:58:58');

-- --------------------------------------------------------

--
-- 表的结构 `admin_panel`
--

CREATE TABLE IF NOT EXISTS `admin_panel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `name` varchar(40) DEFAULT '' COMMENT '菜单名称',
  `url` varchar(255) DEFAULT '' COMMENT '菜单url',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  UNIQUE KEY `uid` (`id`,`aid`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='快捷面板' AUTO_INCREMENT=5 ;

-- --------------------------------------------------------

--
-- 表的结构 `announce`
--

CREATE TABLE IF NOT EXISTS `announce` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `starttime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `hits` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `cateid` (`status`,`endtime`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `cid` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '栏目id category表相对应id',
  `aid` int(11) NOT NULL COMMENT '管理员ID',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `color` char(24) NOT NULL DEFAULT '' COMMENT '标题颜色',
  `font` char(24) NOT NULL DEFAULT '' COMMENT '标题加粗',
  `thumb` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
  `content` text NOT NULL COMMENT '内容',
  `copyfrom` varchar(100) NOT NULL DEFAULT '' COMMENT '来源',
  `keywords` varchar(100) NOT NULL DEFAULT '' COMMENT '关键字',
  `description` varchar(250) NOT NULL COMMENT '描述',
  `relation` varchar(255) NOT NULL DEFAULT '' COMMENT '相关文章',
  `pagetype` tinyint(1) NOT NULL DEFAULT '0' COMMENT '分页方式',
  `maxcharperpage` mediumint(6) NOT NULL DEFAULT '0' COMMENT '分页字符数',
  `istop` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶 0:不置顶 1:置顶',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否发布 0:不发布 1:发布',
  `hits` tinyint(5) NOT NULL DEFAULT '0' COMMENT '点击数',
  `iscomment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否允许评论',
  `createtime` datetime NOT NULL COMMENT '发布时间',
  `updatetime` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `cid` (`cid`,`id`),
  KEY `istop` (`istop`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='内容表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `category`
--

CREATE TABLE IF NOT EXISTS `category` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '栏目id',
  `pid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '类别 0:栏目 1:单网页',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '栏目名称',
  `enname` varchar(30) NOT NULL DEFAULT '' COMMENT '栏目英文名称',
  `desc` mediumtext NOT NULL COMMENT '描述',
  `url` varchar(100) NOT NULL COMMENT '链接地址',
  `hits` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点击数量',
  `setting` mediumtext NOT NULL COMMENT '栏目配置',
  `order` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `ismenu` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示，1 显示',
  PRIMARY KEY (`id`),
  KEY `module` (`pid`,`order`,`id`),
  KEY `type` (`type`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='内容分类表' AUTO_INCREMENT=17 ;

--
-- 转存表中的数据 `category`
--

INSERT INTO `category` (`id`, `pid`, `type`, `name`, `enname`, `desc`, `url`, `hits`, `setting`, `order`, `ismenu`) VALUES
(1, 0, 0, '网站介绍', 'About', '网站介绍', 'javascript:;', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"网站介绍","meta_keywords":"网站介绍","meta_title":"网站介绍"}', 0, 1),
(2, 1, 1, '发展历程', 'Events', '发展历程', '/About/Fzlc/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"发展历程","meta_keywords":"发展历程","meta_title":"发展历程"}', 0, 1),
(3, 1, 1, '企业理念', 'Concept', '企业理念', '/About/Qyln/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"企业理念","meta_keywords":"企业理念","meta_title":"企业理念"}', 0, 1),
(4, 1, 1, '招兵买马', 'Job', '招兵买马', '/About/Job/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"招兵买马","meta_keywords":"招兵买马","meta_title":"招兵买马"}', 0, 1),
(5, 1, 1, '相关证书', 'Certificate', '相关证书', '/About/Xgzs/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"相关证书","meta_keywords":"相关证书","meta_title":"相关证书"}', 0, 1),
(6, 1, 1, '合作伙伴', 'Partners', '合作伙伴', '/About/Hzhb/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"合作伙伴","meta_keywords":"合作伙伴","meta_title":"合作伙伴"}', 0, 1),
(7, 1, 1, '联系我们', 'Contact us', '联系我们', '/About/Contactus/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"联系我们","meta_keywords":"联系我们","meta_title":"联系我们"}', 0, 1),
(8, 1, 1, '汇款方式', 'Remittance way', '汇款方式', '/About/Pay/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"汇款方式","meta_keywords":"汇款方式","meta_title":"汇款方式"}', 0, 1),
(9, 1, 1, '投诉建议', 'Complaints Suggestions', '投诉建议', '/About/Complaints/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"投诉建议","meta_keywords":"投诉建议","meta_title":"投诉建议"}', 0, 1),
(10, 0, 0, '新闻', 'News', '新闻', '/News/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"新闻","meta_keywords":"新闻","meta_title":"新闻"}', 0, 1),
(11, 10, 0, '行业趋势', 'Industry trends', '行业趋势', '/News/trends/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"行业趋势","meta_keywords":"行业趋势","meta_title":"行业趋势"}', 0, 1),
(12, 10, 0, '公司新闻', 'Company news', '公司新闻', '/News/company/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"公司新闻","meta_keywords":"公司新闻","meta_title":"公司新闻"}', 0, 1),
(13, 0, 0, '产品', 'Product', '产品', '/Product/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"产品","meta_keywords":"产品","meta_title":"产品"}', 0, 1),
(14, 0, 0, '案例', 'Case', '案例', '/Case/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"案例","meta_keywords":"案例","meta_title":"案例"}', 0, 1),
(15, 0, 0, '方案', 'Solution', '方案', '/Solution/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"方案","meta_keywords":"方案","meta_title":"方案"}', 0, 1),
(16, 1, 1, '公司信息', 'Company information', '公司信息', '/About/', 0, '{"content_ishtml":"1","ishtml":"1","meta_desc":"公司信息","meta_keywords":"公司信息","meta_title":"公司信息"}', 0, 1);

-- --------------------------------------------------------

--
-- 表的结构 `comment`
--

CREATE TABLE IF NOT EXISTS `comment` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `replyid` int(10) NOT NULL DEFAULT '0' COMMENT '回复ID',
  `aid` int(10) NOT NULL DEFAULT '0' COMMENT '新闻ID',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
  `uid` int(10) NOT NULL DEFAULT '0' COMMENT '会员uid',
  `agree` int(10) unsigned DEFAULT '0' COMMENT '同意',
  `against` int(10) unsigned DEFAULT '0' COMMENT '反对',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT 'IP',
  `createtime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '时间',
  PRIMARY KEY (`id`),
  KEY `replyid` (`replyid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='内容评论表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `copyfrom`
--

CREATE TABLE IF NOT EXISTS `copyfrom` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `sitename` varchar(30) NOT NULL COMMENT '网站名称',
  `siteurl` varchar(100) NOT NULL COMMENT '网站url',
  `thumb` varchar(100) NOT NULL COMMENT '缩略图',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='来源表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `logs`
--

CREATE TABLE IF NOT EXISTS `logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` int(10) unsigned NOT NULL COMMENT 'uid',
  `module` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '模型',
  `url` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的url',
  `action` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的action',
  `ip` varchar(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '操作者所在IP',
  `desc` text COLLATE utf8_unicode_ci NOT NULL COMMENT '操作说明',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`module`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='后台操作日志表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `menu`
--

CREATE TABLE IF NOT EXISTS `menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` char(40) NOT NULL DEFAULT '' COMMENT '名称',
  `enname` char(40) NOT NULL DEFAULT '' COMMENT '英文名称',
  `url` char(100) NOT NULL DEFAULT '' COMMENT '功能地址',
  `data` char(100) DEFAULT '' COMMENT '附加参数',
  `order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `display` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示，1:显示 0:不显示',
  PRIMARY KEY (`id`),
  KEY `listorder` (`order`),
  KEY `parentid` (`pid`),
  KEY `module` (`url`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='后台菜单表' AUTO_INCREMENT=34 ;

--
-- 转存表中的数据 `menu`
--

INSERT INTO `menu` (`id`, `pid`, `name`, `enname`, `url`, `data`, `order`, `display`) VALUES
(1, 0, '我的面板', 'Panel', 'Panel', '', 10000, 1),
(2, 0, '设置', 'Settings', 'Setting', '', 20000, 1),
(3, 0, '模块', 'Modules', 'Module', '', 30000, 1),
(4, 0, '内容', 'Content', 'Content', '', 40000, 1),
(5, 0, '用户', 'Users', 'User', '', 50000, 1),
(6, 0, '扩展', 'Extensions', 'Extend', '', 60000, 1),
(7, 0, '界面', 'Templates', 'Style', '', 70000, 1),
(8, 0, '应用', 'Plugin', 'Plugin', '', 80000, 1),
(9, 2, '菜单设置', 'Menu Settings', 'javascript:;', '', 20100, 1),
(10, 9, '菜单管理', 'Menu management', 'Menu', '', 20101, 1),
(11, 1, '个人设置', 'Personal Settings', 'javascript:;', '', 10100, 1),
(12, 11, '个人信息', 'Personal information', 'EditInfo', '', 10101, 1),
(13, 11, '修改密码', 'Change password', 'EditPwd', '', 10102, 1),
(14, 2, '管理员管理', 'Administrator manager', 'javascript:;', '', 20200, 1),
(15, 14, '管理员管理', 'Administrator manager', 'Admin', '', 20201, 1),
(16, 14, '角色管理', 'Role management', 'Role', '', 20202, 1),
(17, 2, '日志管理', 'Log management', 'javascript:;', '', 20300, 1),
(18, 17, '日志管理', 'Log management', 'Logs', '', 20301, 1),
(19, 1, '快捷面板', 'Shortcut panel', 'javascript:;', '', 10200, 1),
(20, 4, '内容管理', 'Content management', 'javascript:;', '', 40100, 1),
(21, 4, '相关设置', 'Related settings', 'javascript:;', '', 40200, 1),
(22, 20, '栏目管理', 'Manage column', 'Category', '', 40101, 1),
(23, 20, '内容管理', 'Manage content', 'Content', '', 40102, 1),
(24, 3, '模块管理', 'Manage module', 'javascript:;', '', 30100, 1),
(25, 24, '公告', 'Announcement', 'Announce', '', 30101, 1),
(26, 6, '扩展', 'Extensions', 'javascript:;', '', 60100, 1),
(27, 26, '来源管理', 'Source management', 'Copyfrom', '', 60101, 1),
(28, 5, '会员管理', 'Manage user', 'javascript:;', '', 50100, 1),
(29, 5, '会员组管理', 'Manage user group', 'javascript:;', '', 50200, 1),
(30, 28, '会员管理', 'Manage user', 'User', '', 50101, 1),
(31, 29, '管理会员组', 'Manage user group', 'Group', '', 50201, 1),
(32, 7, '模板管理', 'Manage template', 'javascript:;', '', 70100, 1),
(33, 32, '模板风格', 'Style template', 'Style', '', 70101, 1);

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

CREATE TABLE IF NOT EXISTS `role` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '角色说明',
  `data` text NOT NULL COMMENT '菜单列表',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `roleid` (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表' AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `role`
--

INSERT INTO `role` (`id`, `rolename`, `desc`, `data`, `status`, `createtime`) VALUES
(1, '超级管理员', '超级管理员', '1,11,12,13,2,9,10,14,15,16,17,18,3,4,5,6,7,8', 1, '2014-01-18 00:09:09'),
(2, '网站管理员', '网站管理员', '1,11,12,13,2,9,10,14,15,16,17,18,3,5,6,7,8', 1, '2014-02-10 22:18:18');

-- --------------------------------------------------------

--
-- 表的结构 `template`
--

CREATE TABLE IF NOT EXISTS `template` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `identity` varchar(50) NOT NULL DEFAULT '' COMMENT '风格标识',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '风格中文名',
  `author` varchar(50) NOT NULL DEFAULT '' COMMENT '风格作者',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '风格版本',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='模板风格' AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `template`
--

INSERT INTO `template` (`id`, `identity`, `name`, `author`, `version`, `status`, `createtime`) VALUES
(1, 'default', '默认模板', 'GOCMS TEAM', '1.0', 1, '2014-05-08 21:47:47');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `email` char(32) NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `username` char(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `encrypt` char(6) NOT NULL DEFAULT '' COMMENT '随机码',
  `nickname` char(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `regip` char(15) NOT NULL DEFAULT '' COMMENT '注册ip',
  `regdate` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '注册时间',
  `lastdate` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `lastip` char(15) NOT NULL DEFAULT '' COMMENT '上次登录ip',
  `loginnum` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '登陆次数',
  `groupid` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户组id',
  `areaid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '地区id',
  `amount` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金钱总额',
  `point` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `ismessage` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否有短消息',
  `islock` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否定锁',
  `vip` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'vip等级',
  `overduedate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'vip过期时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核状态 5:用户名已存在;4:拒绝;3:删除:2:忽略;0:未审核;1:通过',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `email` (`email`(20))
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=4 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `email`, `username`, `password`, `encrypt`, `nickname`, `mobile`, `birthday`, `regip`, `regdate`, `lastdate`, `lastip`, `loginnum`, `groupid`, `areaid`, `amount`, `point`, `ismessage`, `islock`, `vip`, `overduedate`, `status`, `createtime`) VALUES
(1, 'zzdboy@163.com', 'zzdboy', 'e10adc3949ba59abbe56e057f20f883e', 'ODTFUP', 'zzdboy', '13426456330', '1994-05-19', '192.168.1.5', '2014-05-05 22:18:18', '2014-05-05 22:18:18', '192.168.1.5', 0, 1, 0, '0.00', 34, 0, 2, 1, 2015, 0, '2014-05-05 22:18:18'),
(2, 'zjdboy@163.com', 'zjdboy', 'e10adc3949ba59abbe56e057f20f883e', 'GLWDBW', 'zjdboy', '13426456330', '2014-05-22', '192.168.1.4', '2014-05-06 21:40:40', '2014-05-06 21:40:40', '192.168.1.4', 0, 1, 0, '0.00', 45, 0, 2, 0, 2014, 0, '2014-05-06 21:40:40'),
(3, 'demo@163.com', 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'WGUSHK', 'wqewqe', '13426456330', '2014-05-29', '192.168.1.3', '2014-05-12 21:32:32', '2014-05-12 21:32:32', '192.168.1.3', 0, 1, 0, '0.00', 34, 0, 2, 1, 2014, 0, '2014-05-12 21:32:32');

-- --------------------------------------------------------

--
-- 表的结构 `user_group`
--

CREATE TABLE IF NOT EXISTS `user_group` (
  `id` tinyint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员组id',
  `name` char(15) NOT NULL DEFAULT '' COMMENT '用户组名称',
  `issystem` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是系统组',
  `star` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '会员组星星数',
  `point` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '积分范围',
  `allowmessage` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '允许发短消息数量',
  `allowvisit` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许访问',
  `allowpost` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许发稿',
  `allowpostverify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否投稿不需审核',
  `allowsearch` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许搜索',
  `allowupgrade` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许自主升级',
  `allowsendmessage` tinyint(1) unsigned NOT NULL COMMENT '允许发送短消息',
  `allowpostnum` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '每天允许发文章数',
  `allowattachment` tinyint(1) NOT NULL COMMENT '是否允许上传附件',
  `priceyear` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包年价格',
  `pricemonth` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包月价格',
  `priceday` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包天价格',
  `icon` char(30) NOT NULL COMMENT '用户组图标',
  `usernamecolor` char(7) NOT NULL COMMENT '用户名颜色',
  `desc` char(100) NOT NULL COMMENT '描述',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否禁用',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=8 ;

--
-- 转存表中的数据 `user_group`
--

INSERT INTO `user_group` (`id`, `name`, `issystem`, `star`, `point`, `allowmessage`, `allowvisit`, `allowpost`, `allowpostverify`, `allowsearch`, `allowupgrade`, `allowsendmessage`, `allowpostnum`, `allowattachment`, `priceyear`, `pricemonth`, `priceday`, `icon`, `usernamecolor`, `desc`, `status`, `createtime`) VALUES
(1, '游客', 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, '0.00', '0.00', '0.00', '', '#66CC00', '', 1, '0000-00-00 00:00:00'),
(2, '新手上路', 1, 1, 50, 100, 1, 1, 0, 0, 0, 1, 0, 0, '50.00', '10.00', '1.00', '', '', '', 1, '0000-00-00 00:00:00'),
(3, '注册会员', 1, 2, 100, 150, 0, 1, 0, 0, 1, 1, 0, 0, '300.00', '30.00', '1.00', '', '', '', 1, '0000-00-00 00:00:00'),
(4, '中级会员', 1, 3, 150, 500, 1, 1, 0, 1, 1, 1, 0, 0, '500.00', '60.00', '1.00', '', '', '', 1, '0000-00-00 00:00:00'),
(5, '高级会员', 1, 5, 300, 999, 1, 1, 0, 1, 1, 1, 0, 0, '360.00', '90.00', '5.00', '', '', '', 1, '0000-00-00 00:00:00'),
(6, '禁止访问', 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, '0.00', '0.00', '0.00', '', '', '0', 1, '0000-00-00 00:00:00'),
(7, '邮件认证', 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '0.00', '0.00', '0.00', 'images/group/vip.jpg', '#000000', '', 1, '0000-00-00 00:00:00');
