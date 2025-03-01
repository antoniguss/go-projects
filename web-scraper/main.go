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
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	wg.Wait()

	reportOrderStatus(orders)

	fmt.Println("All operations completed")
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)

		// Get random status
		status := []string{
			"Processing", "Shipped", "Delivered",
		}[rand.Intn(3)]

		order.Status = status

		fmt.Printf("Updated order %d status: %s\n", order.ID, status)
	}
}

func processOrders(orders []*Order) {
	for _, order := range orders {
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

func reportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")

		for _, order := range orders {
			fmt.Printf(
				"Order %d: %s\n",
				order.ID, order.Status,
			)
		}

		fmt.Println("-------------")
		fmt.Println()
	}
}
