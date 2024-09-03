package concurrency

import (
	"fmt"
	"time"
)

func greet(phrase string, channel chan bool) {
	fmt.Println("Hi", phrase)
	channel <- true
}

func slowGreet(phrase string, channel chan bool) {
	time.Sleep(3*time.Second)
	fmt.Println(phrase)
	channel <- true
}

func Main() {
	doneChans := make([]chan bool, 4)
	go greet("Mom!!!", doneChans[0])
	go greet("Dad!!!!", doneChans[1])
	go slowGreet("How... are ... you...??", doneChans[2])
	go greet("brother!!!", doneChans[3])

	for _, chans := range doneChans {
		<- chans
		close(chans)
	}
}