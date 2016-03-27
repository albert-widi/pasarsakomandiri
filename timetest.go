package main

import (
	"fmt"
	"time"
)

func test1(c chan int) {
	time.Sleep(100)
	c <- 4
}

func test2() int {
	time.Sleep(40)
	return 2
}

func main() {
	chan1 := make(chan int)
	go test1(chan1)
	asd := test2()

	//fmt.Println(<-chan1)
	//fmt.Println(asd)
	fmt.Println(<-chan1, " ", asd)

	shift1 := 8

	fmt.Println(shift1 >> 3)

}
