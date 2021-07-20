package tests

import (
	"fmt"
	test "golog/core"
)

func TestSpyder() {
	responseBody := test.FetchUrl("https://wwww.baidu.com", "GET", nil)
	fmt.Printf(responseBody)
}
