module github.com/leeexing/douban-movie

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/go-ini/ini v1.57.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
)

replace (
	github.com/leeexing/douban-movie/conf => E:/Leeing/go/repository/douban-movie/conf
	github.com/leeexing/douban-movie/download => E:/Leeing/go/repository/douban-movie/download
	github.com/leeexing/douban-movie/model => E:/Leeing/go/repository/douban-movie/model
	github.com/leeexing/douban-movie/setting => E:/Leeing/go/repository/douban-movie/setting
	github.com/leeexing/douban-movie/spider => E:/Leeing/go/repository/douban-movie/spider
	github.com/leeexing/douban-movie/util => E:/Leeing/go/repository/douban-movie/util
)
