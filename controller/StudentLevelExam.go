package controller

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"net/http"
	"nngllgjw/config"
	"strings"
)

type StudentLExam struct {
	ExamName  string `json:"ExamName"`
	ExamTime  string `json:"ExamTime"`
	ExamScore string `json:"ExamScore"`
}

func GetStudentLExam(cookies *http.Cookie) ([]*StudentLExam, error) {
	AllStudentLExam := []*StudentLExam{}
	request := gorequest.New()
	resp, _, errs := request.Get(config.BaseURL+config.Level_exam_scores).
		Set("Cookie", cookies.String()).
		End()
	if errs != nil {
		return nil, errs[0]
	}
	defer resp.Body.Close()
	gbkReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(gbkReader)
	if err != nil {
		return nil, err
	}
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i != 0 {
					tds := s.Find("td")
					StudentLExam := &StudentLExam{
						ExamName:  strings.TrimSpace(tds.Eq(0).Text()),
						ExamTime:  strings.TrimSpace(tds.Eq(1).Text()),
						ExamScore: strings.TrimSpace(tds.Eq(2).Text()),
					}
					AllStudentLExam = append(AllStudentLExam, StudentLExam)
				}

			})
		}
	})
	return AllStudentLExam, nil
}
