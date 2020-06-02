package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/leeexing/douban-movie/model"
	"github.com/leeexing/douban-movie/spider"
)

var (
	// BaseURL 基础爬虫路径
	BaseURL = "https://movie.douban.com/top250"
)

// Add 添加数据
func Add(movies []model.DoubanMovie) {
	for index, movie := range movies {
		log.Printf("index: %d >>> movie: %#v ", index, movie)
		model.AddMovie(&movie)
	}
}

// Start 开始执行
func Start() {
	fmt.Println("start ...")
	var movies []model.DoubanMovie

	pages := spider.GetPages(BaseURL)
	for _, page := range pages {
		url := strings.Join([]string{BaseURL, page.URL}, "")
		movieArr := spider.GetMovies(url)
		movies = append(movies, movieArr...)
	}

	Add(movies)

	fmt.Println("end !!!")
}

func main() {
	start := time.Now()

	Start()

	log.Println("总共花费时间：", time.Since(start))

	defer model.DB.Close()
}
