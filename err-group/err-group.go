package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/*
	errgroup provides synchronization, error propagation, and Context cancellation for groups of goroutines working on subtasks.
	In other words you can use this in scenarios where typically sync.WaitGroup is used but this one also takes care of passing
	a Context into the subtasks and handling errors automatically for you.
*/
func main() {
	processTasks()
}

/*
	We get reference to Context which if we want can use it to cancel the tasks or time them out if needed.
	Only caveat to watch out here is that the first call to return a non-nil error cancels the group and its
	error will be returned by Wait.
*/
func processTasks() {
	grp, _ := errgroup.WithContext(context.TODO())
	for i := 0; i < 10; i++ {
		currentIdx := i
		grp.Go(func() error {
			return requestAPI(currentIdx)
		})
	}

	err := grp.Wait()
	if err != nil {
		fmt.Printf("Error executing tasks: %s \n", err.Error())
	} else {
		fmt.Println("All tasks completed with no error")
	}
}

func requestAPI(i int) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}
