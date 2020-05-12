package main

import (
	"fmt"
	"time"
)

func forever(c chan byte) {
	var input byte
	for {
		fmt.Scanln(&input)
		c <- input
	}
}

func main() {

	c := make(chan byte)
	go forever(c)

	sleep := false

	for !sleep {
		select {
		case <-c:
			continue
		case <-time.After(time.Second * 5):
			fmt.Println("Screen timeout")
			sleep = true
		}
	}

}
