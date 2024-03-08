package utils

import "fmt"

func GetYearAndTerm(year string, term string) string {
	yearIDMap := make(map[string]string)
	semesterMap := make(map[string]string)
	yearIDMap["2024"] = "44"
	yearIDMap["2014"] = "34"
	yearIDMap["2015"] = "35"
	yearIDMap["2016"] = "36"
	yearIDMap["2017"] = "37"
	yearIDMap["2018"] = "38"
	yearIDMap["2019"] = "39"
	yearIDMap["2020"] = "40"
	yearIDMap["2021"] = "41"
	yearIDMap["2022"] = "42"
	yearIDMap["2023"] = "43"
	yearIDMap["2025"] = "45"
	yearIDMap["2026"] = "46"
	yearIDMap["2027"] = "47"
	yearIDMap["2028"] = "48"
	yearIDMap["2029"] = "49"
	yearIDMap["2030"] = "50"
	yearIDMap["2031"] = "51"
	yearIDMap["2032"] = "52"
	yearIDMap["2033"] = "53"
	yearIDMap["2034"] = "54"
	semesterMap["春"] = "1"
	semesterMap["秋"] = "3"
	semesterMap["全部"] = ""
	year = yearIDMap[year]
	term = semesterMap[term]
	send := fmt.Sprintf("year=%s&term=%s&prop=&groupName=&para=0&sortColumn=&Submit=查询", year, term)
	return send
}
