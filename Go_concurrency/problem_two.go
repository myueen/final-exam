package main

import (
	"fmt"
	"math/rand"
	"time"
)

// make two buffer channels to store data
var inCh = make(chan int, 5)
var outCh = make(chan int, 5)

// temporar value to compare to find maximum
var temp_val int = 0

func main() {
	go producer_one()
	go producer_two()
	go consumer_one()
	go consumer_two()
	go filter_number()

	time.Sleep(20 * time.Second)
}

func producer_one() {
	// Producer one generates odd integers from 1 to 29 (inclusive)
	for i := 1; i <= 29; i += 2 {
		inCh <- i
		fmt.Printf("Producer One generates: %d\n", i)
		r := rand.Intn(1500) // Random integer between 0 and 1500
		time.Sleep(time.Millisecond * time.Duration(r))
	}
}

func producer_two() {
	// Producer two generates even integers from 1 to 30 (inclusive)
	for i := 2; i <= 30; i += 2 {
		inCh <- i
		fmt.Printf("Producer 2 generates: %d\n", i)
		r := rand.Intn(1500) // Random integer between 0 and 1500
		time.Sleep(time.Millisecond * time.Duration(r))
	}
}

func consumer_one() {
	for {
		num := <-inCh
		outCh <- num * num
		fmt.Printf("Consumer 1 received: %d\n", num)
		r := rand.Intn(3) // Random integer between 0 and 3 second
		time.Sleep(time.Second * time.Duration(r))
	}
}

func consumer_two() {
	for {
		num := <-inCh
		outCh <- num * num
		fmt.Printf("Consumer 2 received: %d\n", num)
		r := rand.Intn(3) // Random integer between 0 and 3 second
		time.Sleep(time.Second * time.Duration(r))
	}
}

func filter_number() {
	pause_time := 20 * time.Second
	current_time := time.Now()

	for {
		select {
		case num := <-outCh:
			if num > temp_val {
				temp_val = num
				fmt.Printf("Numner is: %d\n", num)
			}
		default:
			if time.Since(current_time) > pause_time {
				fmt.Println("No new numbers for 40 seconds. Process terminated.")
				return
			}
		}
	}
}
