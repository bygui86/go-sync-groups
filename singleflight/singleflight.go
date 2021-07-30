package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
	singleflight provides a duplicate function call suppression mechanism, which means if multiple goroutines call a same
	function concurrently this package ensures that only one execution is in-flight for a given key at a time.
	If a duplicate comes in, the duplicate caller waits for the original to complete and receives the same results.
*/
func main() {
	processTasks()
}

/*
	You will observe that this prints "heavyTask()" only once even though there are 5 goroutines calling the same function.
	This is particularly useful when interacting with the “slow” function which are being called concurrently such as with
	database, reading files or making HTTP calls.
*/
func processTasks() {
	sf := singleflight.Group{}
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			value, _, _ := sf.Do("my_task", heavyTask)
			fmt.Printf("Task %d value: %s \n", id, value.(string))
		}(i)
	}

	wg.Wait()
}

func heavyTask() (interface{}, error) {
	println("executing heavy task")
	time.Sleep(time.Second * 5)
	return "heavy-task-result", nil
}
