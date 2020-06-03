package spider

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	downloader "github.com/leeexing/douban-movie/download"
	"github.com/leeexing/douban-movie/model"
)

// Page 页码
type Page struct {
	Page int
	URL  string
}

var digitCompile = regexp.MustCompile("\\d+")

// GetPages 获取分页
func GetPages(url string) []Page {
	body := downloader.GetGbk(url)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	return ParserPages(doc)
}

// ParserPages 解析页面数据
func ParserPages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, URL: ""})
	doc.Find("#content .paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			URL:  url,
		})
	})

	return pages
}

// GetMovies 解析网页中的电影信息
func GetMovies(url string) (movies []model.DoubanMovie) {
	body := downloader.GetGbk(url)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatalf("load movie detail err: %v", err)
	}

	return parserMovie(doc)
}

// parserMovie 解析网页中电影信息
func parserMovie(doc *goquery.Document) (movies []model.DoubanMovie) {
	doc.Find("#content .article li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.Replace(subtitle, "/", "", 1)
		subtitle = strings.Replace(subtitle, "\u00a0", "", 2)

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.Replace(other, "/", "", 1)
		other = strings.Replace(other, "\u00a0", "", 2)

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		desc = strings.Replace(desc, "\u00a0", " ", -1)
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := digitCompile.FindString(movieDesc[0])
		area := strings.TrimSpace(movieDesc[len(movieDesc)-2])
		tag := strings.TrimSpace(movieDesc[len(movieDesc)-1])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		comment = digitCompile.FindString(comment)
		// compile := regexp.MustCompile("[0-9]")
		// comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := model.DoubanMovie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		movies = append(movies, movie)
	})

	return movies
}
