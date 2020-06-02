package downloader

import (
	"io"
	"log"
	"net/http"

	"github.com/leeexing/douban-movie/util"
)

// Get 普通下载方法，不包含设置请求头的操作
func Get(url string) io.Reader {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body
}

// GetGbk 获取转码之后的内容
func GetGbk(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("http.Request err: %s", err)
	}

	req.Header.Add("User-Agent", util.GetUserAgent())
	req.Header.Add("Host", "movie.douban.com")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.do err: %v", err)
	}

	// mahonia 转码字符串编码
	// mah := mahonia.NewDecoder("gbk")
	// return mah.NewReader(res.Body)

	return res.Body
}
