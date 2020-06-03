package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func migration(DB *gorm.DB) {
	// 自动迁移模式
	fmt.Println("正在进行自动化迁移中...")
	DB.Set("gorm.table_options", "charset=utf8mb4")
	DB.AutoMigrate(&DoubanMovie{})
}
