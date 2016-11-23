package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	tn := time.Now()
	out := make(chan string)
	for i := 1; i <= 10; i++ {
		go htmlTime(out, i)
		//fmt.Print(<-out) I could print this here... but
	}
	// if you don't make two separate loops you'll never get the go routines
	// running concurrently as <- out will block
	for i := 1; i <= 10; i++ {
		fmt.Print(<-out)
	}
	dur := time.Since(tn)
	fmt.Printf("Done in %v\n", dur)
}

func htmlTime(out chan<- string, i int) {
	tn := time.Now()
	res, err := http.Get("http://www.apple.com")
	if err != nil {
		log.Fatal(err)
	}
	dur := time.Since(tn)
	out <- fmt.Sprintf("%v, URL Download Time: %v, Loop Number: %v\n", res.Status, dur, i)
	res.Body.Close()
}
