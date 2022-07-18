package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go timer()
	time.Sleep(20 * time.Second)
	fmt.Println("\nstoped")
}

func timer() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		fmt.Printf("%s\r", time.Since(start))
		time.Sleep(1 * time.Millisecond)
	}

}
