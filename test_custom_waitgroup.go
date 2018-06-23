package main

import (
	"learngo/main/simpleSync"
	"fmt"
	"time"
)




func process(i int, wg *simpleSync.CustomWaitGroupTest) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main()  {

	no := 3

	var wg simpleSync.CustomWaitGroupTest;

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")

}