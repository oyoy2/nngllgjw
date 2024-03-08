package controller

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"nngllgjw/config"
	"regexp"
	"strings"
)

func GetCurrcourseAsExcel(cookie *http.Cookie) ([][]string, error) {
	request := gorequest.New()
	resp, body, _ := request.Get(config.BaseURL+config.Current_semester_schedule).
		Set("Cookie", cookie.String()).
		End()
	defer resp.Body.Close()
	utf8Body, _ := ioutil.ReadAll(transform.NewReader(strings.NewReader(body), simplifiedchinese.GBK.NewDecoder()))
	body = string(utf8Body)
	pattern := `window\.open\('(\.\./\.\./manager/coursearrange/showTimetable\.do\?id=\d+&yearid=\d+&termid=\d+&timetableType=STUDENT&sectionType=BASE)'\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(body)
	url := matches[1][6:]
	resp, body, _ = request.Get(config.BaseURL+url).
		Set("Cookie", cookie.String()).
		End()
	defer resp.Body.Close()
	utf8Body, _ = ioutil.ReadAll(transform.NewReader(strings.NewReader(body), simplifiedchinese.GBK.NewDecoder()))
	body = string(utf8Body)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return [][]string{}, err
	}
	table := doc.Find("#timetable")
	tableData := [][]string{}

	table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
		rowData := []string{}

		row.Find("th, td").Each(func(colIdx int, cell *goquery.Selection) {
			rowData = append(rowData, cell.Text())
		})

		tableData = append(tableData, rowData)
	})

	return tableData, nil
}
