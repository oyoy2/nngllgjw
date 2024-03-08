package controller

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"nngllgjw/config"
	"nngllgjw/utils"
	"strconv"
	"strings"
)

type StudentOwnScore struct {
	Coursename  string `json:"Coursename"`
	Gradepoints string `json:"Gradepoints"`
	Status      string `json:"Status"`
	Score       string `json:"Score"`
	Credits     string `json:"Credits"`
	Year        string `json:"Year"`
	Categories  string `json:"Categories"`
	Teacher     string `json:"teacher"`
}
type StudentOwnScoreExcel struct {
	Arithmetic string `json:"Arithmetic"`
	Weighted   string `json:"Weighted"`
	AverageGPA string `json:"AverageGPA"`
}

func GetAllStudentOwnScores(cookies *http.Cookie, year string, term string) ([]*StudentOwnScore, error, StudentOwnScoreExcel, int) {
	allScores := []*StudentOwnScore{}
	ScoreExcel := StudentOwnScoreExcel{}
	request := gorequest.New()
	resp, body, errs := request.Post(config.BaseURL+config.Personal_grades_inquiry).
		Set("Cookie", cookies.String()).
		Send("year=" + year + "&term=" + term + "&prop=&groupName=&para=0&sortColumn=&Submit=%E6%9F%A5%E8%AF%A2").
		End()
	if errs != nil {
		return nil, errs[0], ScoreExcel, 0
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err, ScoreExcel, 0
	}
	var totalCreditPoints float64
	var totalCredits float64
	var totalScores float64
	var totalScoresW float64
	var Fail int
	var total float64
	var Weighted float64

	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				if i != 0 {
					tds := s.Find("td")
					if utils.ExcludeGPA(strings.TrimSpace(tds.Eq(3).Text())) && strings.TrimSpace(tds.Eq(11).Text()) != "学位课" {
						credit, err := strconv.ParseFloat(strings.TrimSpace(tds.Eq(7).Text()), 64)
						if err != nil {
							fmt.Println("学分数转换错误:", err)
							return
						}
						gradePoint, err := strconv.ParseFloat(strings.TrimSpace(tds.Eq(6).Text()), 64)
						if err != nil {
							fmt.Println("绩点数转换错误:", err)
							return
						}
						if strings.Contains(strings.TrimSpace(tds.Eq(3).Text()), "习近平新时代中国特色社会主义思想概论") {
							tds.Eq(3).SetText("习概")
						} else if strings.Contains(strings.TrimSpace(tds.Eq(3).Text()), "马克思主义基本原理") {
							tds.Eq(3).SetText("马思")
						} else if strings.Contains(strings.TrimSpace(tds.Eq(3).Text()), "毛泽东思想和中国特色社会主义理论体系概论") {
							tds.Eq(3).SetText("毛概")
						}
						switch s := strings.TrimSpace(tds.Eq(5).Text()); {
						case s == "优":
							tds.Eq(5).SetText("95")
						case s == "良":
							tds.Eq(5).SetText("85")
						case s == "中":
							tds.Eq(5).SetText("75")
						case s == "及格":
							tds.Eq(5).SetText("65")
						case s == "合格":
							tds.Eq(5).SetText("65")
						case s == "不及格":
							tds.Eq(5).SetText("40")
						case s == "不合格":
							tds.Eq(5).SetText("40")
						}
						Score, err := strconv.ParseFloat(strings.TrimSpace(tds.Eq(5).Text()), 64)
						if err != nil {
							fmt.Println("总评成绩转换错误:", err)
							return
						}
						score := &StudentOwnScore{
							Coursename:  strings.TrimSpace(tds.Eq(3).Text()),
							Score:       strings.TrimSpace(tds.Eq(5).Text()),
							Gradepoints: strings.TrimSpace(tds.Eq(6).Text()),
							Credits:     strings.TrimSpace(tds.Eq(7).Text()),
							Status:      strings.TrimSpace(tds.Eq(12).Text()),
							Categories:  strings.TrimSpace(tds.Eq(11).Text()),
							Year:        strings.TrimSpace(tds.Eq(0).Text()),
							Teacher:     strings.TrimSpace(tds.Eq(4).Text()),
						}
						allScores = append(allScores, score)
						if strings.TrimSpace(tds.Eq(12).Text()) == "不及格" {
							Fail++
							return
						}
						totalCreditPoints += credit * gradePoint
						total++
						totalCredits += credit
						totalScores += Score
						totalScoresW += Score * credit
					}
				}
			})
		}
	})

	averageGPA := totalCreditPoints / totalCredits
	Weighted = totalScoresW / totalCredits
	Arithmetic := totalScores / total
	averageGPAString := strconv.FormatFloat(float64(averageGPA), 'f', 2, 64)
	WeightedString := strconv.FormatFloat(float64(Weighted), 'f', 2, 64)
	ArithmeticString := strconv.FormatFloat(float64(Arithmetic), 'f', 2, 64)
	ScoreExcel = StudentOwnScoreExcel{
		Arithmetic: ArithmeticString,
		Weighted:   WeightedString,
		AverageGPA: averageGPAString,
	}
	return allScores, nil, ScoreExcel, Fail
}
