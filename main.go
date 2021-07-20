package main

import (
	"fmt"
	dao "golog/core"
	file "golog/logsource"
	localTests "golog/tests"

	"math/rand"
	"time"
)

func testTime() {
	fmt.Printf("3333")

}

func tset11(){
	file.TailExample()
}

func test2() {
	localTests.TestSpyder()
}

func test3()  {
	dao.GetRandLine()
}

func randC() {
	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(150)
	fmt.Println(k)
}

func main() {
	localTests.TestChannel()
}
