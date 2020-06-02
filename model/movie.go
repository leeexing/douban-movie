package model

import "log"

// DoubanMovie 豆瓣音乐数据结构体
type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

// AddMovie 数据入库
func AddMovie(movie *DoubanMovie) {
	if err := DB.Create(&movie).Error; err != nil {
		log.Fatalf("db.Create movie: %s", err)
	}
}
