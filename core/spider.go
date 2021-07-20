package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// https://zhuanlan.zhihu.com/p/55039990
func FetchUrl (url string, method string, payload io.Reader) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Set("X-Authentication-Custom", "TSC BDP")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}


func TestSpyder() {
	responseBody := FetchUrl("https://wwww.baidu.com", "GET", nil)
	fmt.Printf(responseBody)
}