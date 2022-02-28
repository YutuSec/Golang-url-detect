package main

import (
	"awesomeProject3/Gettitle/DATA"
	"fmt"
	"log"
	"sync"
	time2 "time"
)

func main() {
	timenow := time2.Now()
	urls, err := DATA.ReadLinefile("url.txt")
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup
	var webinfo DATA.Webinfo
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			webinfo.GETWEBINFO(url, &wg)
		}(url)
	}
	wg.Wait()
	fmt.Printf("探测历时%v秒", time2.Since(timenow))
}
