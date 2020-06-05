package model

import (
	"github.com/leeexing/douban-movie/logging"
)

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
		logging.Fatal("db.Create movie: ", err)
	}
}

// SaveMovieToSQL 保存数据到sql数据库
func SaveMovieToSQL(movies []DoubanMovie)  {
	for index, movie := range movies {
		logging.Info("index:", index, " movie:", movie)
		// log.Printf("index: %d >>> movie: %#v ", index, movie)
		// AddMovie(&movie)
		// SaveMovieToMongo(&movie)
	}
}