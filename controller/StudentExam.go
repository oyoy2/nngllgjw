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

type StudentExam struct {
	CourseId   string `json:"CourseId"`
	CourseName string `json:"CourseName"`
	ExamTime   string `json:"ExamTime"`
	ExamAddr   string `json:"ExamAddr"`
	ExamType   string `json:"ExamType"`
}

func GetStudentExam(cookies *http.Cookie) ([]*StudentExam, error) {
	AllStudentExam := []*StudentExam{}
	request := gorequest.New()
	resp, _, errs := request.Get(config.BaseURL+config.Student_exam_arrangement).
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
					StudentExam := &StudentExam{
						CourseId:   strings.TrimSpace(tds.Eq(0).Text()),
						CourseName: strings.TrimSpace(tds.Eq(1).Text()),
						ExamTime:   strings.TrimSpace(tds.Eq(2).Text()),
						ExamAddr:   strings.TrimSpace(tds.Eq(3).Text()),
						ExamType:   strings.TrimSpace(tds.Eq(4).Text()),
					}
					AllStudentExam = append(AllStudentExam, StudentExam)
				}

			})
		}
	})
	return AllStudentExam, nil
}
