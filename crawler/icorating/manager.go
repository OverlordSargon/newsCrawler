package crawler

import (
	"../../src/github.com/PuerkitoBio/goquery"
	"../../misc"
	"time"
	"strings"
)

var siteMainLink = "http://gamebomb.ru/"

type GamebombCrawler struct {
	news []*NewsPages
}

type RawNews struct {
	Title string
	Link  string
}

func (manager *GamebombCrawler) Init(config misc.Configuration) error {
	rawNews, err := manager.GetEntitiesLinks(siteMainLink)
	if err != nil {
		return err
	}

	for index, element := range rawNews {
		if (strings.Contains(element.Title, config.SearchWord) || len(config.SearchWord) == 0) {
			worker := &NewsPages{
				id:    index,
				title: element.Title,
				link:  element.Link,
			}
			manager.news = append(manager.news, worker)
			go func() {
				worker.StartCrawling()
			}()
		}
	}
	timeout, err := time.ParseDuration(config.UpdateTimeout)
	if err != nil {
		timeout, _ = time.ParseDuration("5m")
	}
	for {
		time.Sleep(timeout)
		workersFinished := true
		for _, worker := range manager.news {
			if !worker.done {
				workersFinished = false
				break
			}
		}
		if workersFinished {
			break
		}
	}
	return nil
}

/**
	Find string in array
 */
func contains(slice []string, search string) bool {
	for _, value := range slice {
		if value == search {
			return true
		}
	}
	return false
}

/**
	Gathering link for further crawling
 */
func (crawler *GamebombCrawler) GetEntitiesLinks(mainPageLink string) ([]RawNews, error) {
	doc, err := goquery.NewDocument(mainPageLink)
	if err != nil {
		return nil, err
	}
	result := []RawNews{}
	allNews := doc.Find(".gbnews-listShort")

	titles := make([]string, allNews.Length())
	links := make([]string, allNews.Length())

	allNews.Find("h3").Each(func(i int, s *goquery.Selection) {
		var resultHeader, error = s.Html()
		if len(resultHeader) != 0 && !contains(titles, resultHeader) {
			titles = append(titles, resultHeader)
		}
		if (error != nil) {
			print("Error while parsing")
		}
	})
	allNews.Find("td").Each(func(i int, s *goquery.Selection) {
		href, found := s.Find("a").Attr("href")
		if found && !contains(links, href) {
			links = append(links, href)
		}
	})

	iterator := 1
	//Проходимся по собранным ссылкам, собираем объекты
	for iterator < allNews.Length()*2 {
		currentTitle := string(titles[iterator])
		currentLink := string(links[iterator])

		if (len(currentTitle) != 0 && len(currentLink) != 0) {
			currentNews := RawNews{
				Title: currentTitle,
				Link:  currentLink,
			}
			result = append(result, currentNews)
		}
		iterator += 1
	}

	return result, nil
}
