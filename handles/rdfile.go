package controllers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const filename string  = "D:\\home\\test\\modsec_audit.log"

func ReadFile()  [] string {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	results := make([] string, 100, 20000)
	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		results  = append(results, string(line))
		//fmt.Println(string(line))
	}
	return results

}
