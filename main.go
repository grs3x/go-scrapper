package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("dados.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string

		row.Find("th, td").Each(func(j int, col *goquery.Selection) {
			cellData := strings.TrimSpace(col.Text())

			if !containsIgnoredKeyword(cellData) {
				rowData = append(rowData, cellData)
			}
		})

		if len(rowData) > 0 {
			err := writer.Write(rowData)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	fmt.Println("Dados gravados no arquivo dados.csv.")
}

func containsIgnoredKeyword(str string) bool {
	ignoredKeywords := []string{"DIAG", "CONF", "TOOL", "CMD"}
	for _, keyword := range ignoredKeywords {
		if strings.Contains(str, keyword) {
			return true
		}
	}
	return false
}
