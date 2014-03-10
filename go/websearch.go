package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = search("web")
	Image = search("image")
	Video = search("video")
)

type Search func(query string) Result

type Result func(result, kind, query string)

func search(kind string) Search {
	return func(query string) {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	// results = append(results, Web(query))
	// results = append(results, Image(query))
	// results = append(results, Video(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
