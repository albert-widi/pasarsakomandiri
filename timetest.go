package main

import (
	"time"
	"fmt"
	"github.com/jinzhu/now"
)

func main() {
	time := time.Now()
	format1 := time.Format("2006-01-02 15:04:05")
	fmt.Println(format1)

	now.TimeFormats = append(now.TimeFormats, "2006-01-02 23:04:05")
	//asd := now.MustParse(time)
}
