package main

import (
	"fmt"
	file "golog/logsource"
	localTests "golog/tests"
	//rdfile "golog/handles"
	modsec "golog/dao"
	parser "golog/handles"
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
	modsec.GetRandLine()
}

func test4(){
	parser.PaserTime()

}


func randC() {
	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(150)
	fmt.Println(k)
}

func main() {
	//randC()
	test4()
}
