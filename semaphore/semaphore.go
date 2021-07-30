package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

/*
	This package provides a weighted semaphore implementation. If you have worked with go channels before you will know
	that buffered channels kind of behaves like a semaphore. Capacity of the buffered channel is the number of resources
	we wish to synchronize, length of the channel is the number of resources current being used and capacity minus the
	length of the channel is the number of free resources.
	However in the case of buffered channels everything is equal weight and it becomes non-trivial to implement it in
	scenarios where a goroutine might pick up “heavy” task and you have to rate limit it accordingly.
*/
func main() {
	processTasks()
}

func processTasks() {
	sem := semaphore.NewWeighted(4)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go func() {
		for i := 0; i < 5; i++ {
			weight := int64(1)
			fmt.Printf("Task A.%d with weight %d acquiring semaphore \n", i, weight)
			err := sem.Acquire(context.TODO(), weight)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second * 1)

			sem.Release(1)
			wg.Done()
			fmt.Printf("Task A.%d with weight %d completed \n", i, weight)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			weight := int64(3)
			fmt.Printf("Task B.%d with weight %d acquiring semaphore \n", i, weight)
			err := sem.Acquire(context.TODO(), weight)
			if err != nil {
				panic(err)
			}

			time.Sleep(time.Second * 3)
			sem.Release(3)
			wg.Done()
			fmt.Printf("Task B.%d with weight %d completed \n", i, weight)
		}
	}()

	wg.Wait()
}
