package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	running := true
	for running {
		running = false

	}

}

func runTimerFor20Sec() {
	forever := make(chan time.Duration)
	ctx, cancel := context.WithCancel(context.Background())

	go timer(ctx, forever)

	fmt.Println("start")

	time.Sleep(20 * time.Second)
	cancel()

	<-forever
	fmt.Println("\nstoped")
}

func timer(ctx context.Context, forever chan time.Duration) {
	start := time.Now()

	for {
		timevalue := time.Since(start)

		select {
		case <-ctx.Done():
			forever <- timevalue
			return
		default:
			fmt.Printf("%s\r", timevalue)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
