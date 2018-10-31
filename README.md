# liteurl-url短链
[![Build Status](https://travis-ci.org/wusphinx/liteurl.svg?branch=master)](https://travis-ci.org/wusphinx/liteurl)
[![codecov](https://codecov.io/gh/wusphinx/liteurl/branch/master/graph/badge.svg)](https://codecov.io/gh/wusphinx/liteurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/wusphinx/liteurl)](https://goreportcard.com/report/github.com/wusphinx/liteurl)
[![GoDoc](https://godoc.org/github.com/wusphinx/liteurl?status.svg)](https://godoc.org/github.com/wusphinx/liteurl)

# Step

1. 获取url哈希(md5)
2. 将url及哈希值入库得到`F_id`
3. 将`F_id`按`baseStr`（0-9a-zA-Z-+）进制表示返回

数据库结构参考如下

```
CREATE TABLE IF NOT EXISTS `t_short_url` (
  `F_id` bigint(64) unsigned NOT NULL AUTO_INCREMENT,
  `F_hash` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `F_url` varchar(512) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`F_id`),
  KEY `hash` (`F_hash`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8
```
