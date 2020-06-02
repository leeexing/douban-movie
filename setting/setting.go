package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	// Cfg 配置文件句柄
	Cfg *ini.File

	// RunMode 运行模式
	RunMode string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
}

// LoadBase 加载基本配置
func LoadBase() {
	sec, err := Cfg.GetSection("")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = sec.Key("RUN_MODE").MustString("debug")
}
