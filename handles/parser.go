package controllers

import (
	"regexp"
	"strconv"
	"time"
)


func getMonth(month string) int {
	monthMap := map[string]int {"Jan": 1, "Feb":2, "Mar":3, "Apr":4, "May":5, "Jun":6, "Jul":7, "Aug":8, "Sep":9, "Oct": 10, "Nov":11, "Dec":12}
	return monthMap[month]
}

/*data TIME_LAYOUT = "Thu Jul  1 04:03:29 2021"
//data C_TIME_LAYLOUT = "2006-01-02 15:04:05"*/
func ParseDateTime(dateStr string)  time.Time {
	// parse time
	// "Thu Jul  1 04:03:29 2021"
	partern := `(?P<sWeek>\w{3})\s*(?P<sMonth>\w{3})\s*(?P<day>\d+)\s*(?P<hour>\d+)\:(?P<minute>\d+)\:(?P<second>\d+)\s*(?P<year>\d+)`
	rx := regexp.MustCompile(partern)
	compiles := rx.FindStringSubmatch(dateStr)
	groupNames := rx.SubexpNames()
	result := make(map[string]string)
	if len(compiles) == len(groupNames) {
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				result[name] = compiles[i]
			}
		}
	}
	month := getMonth(result["sMonth"])
	year,_  := strconv.Atoi(result["year"])
	day,_  := strconv.Atoi(result["day"])
	hour,_ := strconv.Atoi(result["hour"])
	minute,_ := strconv.Atoi(result["minute"])
	second,_ := strconv.Atoi(result["second"])
	dt := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	return dt
}