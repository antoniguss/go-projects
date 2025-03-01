package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Learning concurrency
type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	orderChan := make(chan *Order)

	go func() {
		defer wg.Done()

		for _, order := range generateOrders(20) {
			orderChan <- order
		}

		close(orderChan)

		fmt.Println("Done generating orders")
	}()

	go processOrders(orderChan, &wg)

	wg.Wait()
}

func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}

	return orders
}
