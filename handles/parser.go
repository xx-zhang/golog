package controllers

import (
	"fmt"
	"regexp"
	"strings"
	//"time"
)


func getMonth(month string) int {
	monthMap := strings.Split("Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec", "|")
	for index, m := range monthMap{
		if month == m {
			i_month := index
			return i_month+1
		}
	}
	return 0
}


func PaserTime()  {


	// parse time
	// "Thu Jul  1 04:03:29 2021"
	demo := "Thu Jul  1 04:03:29 2021"
	partern := `\w{3}\s*(\w{3})\s*(\d+)\s*(\d+)\:(\d+)\:(\d+)\s*(\d+)`
	rx := regexp.MustCompile(partern)
	matchs := rx.FindStringSubmatch(demo)

	month := getMonth(matchs[2])
	//year := matchs[6]
	//day := matchs[3]
	//hour := matchs[3]
	//minute := matchs[4]
	//second := matchs[5]
	fmt.Println(month)
}


func ParseMsg() {

}