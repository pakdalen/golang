package main

import (
	"fmt"
)

func a(done chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
	done <- true
}

func b(done chan bool) {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
	done <- true
}

func main() {
	done := make(chan bool)
	go a(done)
	go b(done)
	// time.Sleep(time.Second)
	<-done
	<-done

	fmt.Println("end main()")
}
