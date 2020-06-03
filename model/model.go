package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/leeexing/douban-movie/setting"

	// 私密
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB 数据库连接
	DB *gorm.DB
)

func init() {
	var (
		err                                             error
		makeMigration                                   bool
		dbType, dbName, user, passwd, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("datebase")
	if err != nil {
		log.Fatalf("Fail to get section 'datebase': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	passwd = sec.Key("PASSWD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	makeMigration = sec.Key("MAKE_MIGRATION").MustBool(false)

	DB, err = gorm.Open(dbType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		passwd,
		host,
		dbName))
	if err != nil {
		log.Fatalf("gorm.Open.err: %v", err)
	}

	// 禁用复数. 就是表名后面不添加 s
	DB.SingularTable(true)

	// gorm 默认表名的一个处理函数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	if makeMigration {
		migration(DB)
	}
}

// CloseDB 暴露关闭数据库的方法
func CloseDB() {
	defer DB.Close()
}
