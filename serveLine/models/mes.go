package models

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 10)
	go func() {
		for {
			select {
			case mes := <-ch:
				fmt.Println(mes)
			}
		}
	}()

	ch <- "aaaa"
	ch <- "bb"
	ch <- "cc"
	ch <- "dd"
	ch <- "ee"
	ch <- "ff"
	ch <- "gg"
	<-time.After(2 * time.Second)
}
