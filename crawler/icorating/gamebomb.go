package crawler

import (
	//"../../src/github.com/PuerkitoBio/goquery"
	"github.com/PuerkitoBio/goquery"
	"../../model"
	"../../writer"
	"strings"
)

type NewsPages struct {
	id    int
	title string
	done  bool
	link  string
}

func (news *NewsPages) StartCrawling() error {
	println("newsLink")
	println(news.link)
	resultOfCarawling, _ := news.GetDetails()
	outputPath := "./data/games/"
	outFilename := resultOfCarawling.Title + ".json"
	writer.WriteToFS(outputPath, outFilename, resultOfCarawling)
	return nil
}

func (news *NewsPages) GetDetails() (model.News, error) {
	doc, err := goquery.NewDocument(news.link)
	if err != nil {
		return model.News{}, err
	}
	result := model.News{}
	titleNode := doc.Find("h1")
	if len(titleNode.Nodes) > 0 {
		result.Title = titleNode.Text()
	}

	views, _ := doc.Find(".views-bubble").Html()
	if (len(views) > 0) {
		result.Views = views
	}

	commentsNumber := doc.Find(".h1").Find(".count").Text()
	if (len(commentsNumber) > 0) {
		result.Comments = strings.Replace(strings.Replace(commentsNumber, "(", "", -1), ")", "", -1)
	}

	result.Link = news.link

	return result, nil
}
