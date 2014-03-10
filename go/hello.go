// go generator
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func generator(msg string) <-chan *Message {
	ch := make(chan *Message)
	go func() {
		for i := 0; i < 10; i++ {
			waitForIt := make(chan bool)
			ch <- &Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			// wait to advance
			<-waitForIt
		}
		close(ch)
	}()
	return ch
}

func fanIn(ch1, ch2 <-chan *Message) <-chan *Message {
	ch := make(chan *Message)
	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()
	// go func() {
	// 	for {
	// 		select {
	// 		case ch <- <-ch1:
	// 		case ch <- <-ch2:
	// 		}
	// 	}
	// }()
	return ch
}

func run() {
	fanInCh := fanIn(generator("joe"), generator("ann"))

	for i := 0; i < 10; i++ {
		// 	fmt.Println(<-fanInCh)
		msg1 := <-fanInCh
		fmt.Println(msg1.str)
		msg2 := <-fanInCh
		fmt.Println(msg2.str)
		// signal to advance
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("tired!")
}

func main() {
	run()

	// for v := range generator("foo") {
	// 	fmt.Println(v)
	// }
}
