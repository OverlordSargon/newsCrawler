package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "../../model/icorating"
	"../../writer"
)

type ICORatingWorker struct {
	id       int
	finished bool
	pageType string
	links    []string
}

func (worker *ICORatingWorker) Start() error {
	for _, link := range worker.links {
		entity, _ := worker.GetDetails(link)
		outputPath := "./data/icorating/"
		outFilename := entity.Title + ".json"
		writer.WriteToFS(outputPath, outFilename, entity)
	}
	return nil
}

func (worker *ICORatingWorker) GetDetails(detailsLink string) (model.ICORatingCompany, error) {
	doc, err := goquery.NewDocument(detailsLink)
	if err != nil {
		return model.ICORatingCompany{}, err
	}
	result := model.ICORatingCompany{}
	titleNode := doc.Find("h1")
	if len(titleNode.Nodes) > 0 {
		result.Title = titleNode.Text()
	}
	tableCells := doc.Find("td")
	for i := range tableCells.Nodes {
		cell := tableCells.Eq(i)
		text := cell.Text()
		if text == "Product Type:" {
			result.Type = clearText(cell.Siblings().Text())
		}
		if text == "Industry:" {
			result.Industry = clearText(cell.Siblings().Text())
		}
		if text == "Description:" {
			result.Description = clearText(cell.Siblings().Text())
		}
		if text == "Features:" {
			result.Features = clearText(cell.Siblings().Text())
		}
	}
	return result, nil
}

func clearText(input string) string {
	output := strings.Replace(input, "\n", "", -1)
	output = strings.TrimSpace(output)
	return output
}

// func (worker *ICORatingWorker) GetNews(link string) (interface{}, error) {

// }

// func (worker *ICORatingWorker) GetReview(link string) (interface{}, error) {

// }
