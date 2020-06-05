package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/leeexing/douban-movie/spider"
	"github.com/leeexing/douban-movie/model"
)

var (
	// BaseURL 基础爬虫路径
	BaseURL = "https://movie.douban.com/top250"
)

// Start 开始执行
func Start() {
	fmt.Println("start ...")
	var movies []model.DoubanMovie

	// 有点多余，根据规律可以拼出来
	pages := spider.GetPages(BaseURL)
	for index, page := range pages {
		url := strings.Join([]string{BaseURL, page.URL}, "")
		movieArr := spider.GetMovies(url)
		movies = append(movies, movieArr...)
		if index == 0 {
			break
		}
	}

	model.SaveMovieToSQL(movies)

	fmt.Println("end !!!")
}

func fetch(url string, ch chan []model.DoubanMovie) {
	movieArr := spider.GetMovies(url)
	fmt.Printf("爬取地址：%s, 电影数量：%d \n", url, len(movieArr))
	ch <- movieArr
}

// UseChan 使用管道配合协程进行多线程爬取
func UseChan() {
	fmt.Println("useChan start ...")
	var movies []model.DoubanMovie

	pages := spider.GetPages(BaseURL)
	delayTime := time.Second * 6
	ch := make(chan []model.DoubanMovie)

	for _, page := range pages {
		url := strings.Join([]string{BaseURL, page.URL}, "")
		fmt.Println("协程爬取页面地址: ", url)
		go fetch(url, ch)
	}

L:
	for {
		select {
		case movieArr := <-ch:
			movies = append(movies, movieArr...)
		case <-time.After(delayTime):
			log.Println("timeout")
			break L
		}
	}

	fmt.Println("useChan end !!!")
}

func main() {
	start := time.Now()

	Start()
	// UseChan()

	log.Println("总共花费时间：", time.Since(start))

	defer model.DB.Close()
}
