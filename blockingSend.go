package main

import (
	"fmt"
	"time"
	"sync"
)

func test(ch chan bool, group *sync.WaitGroup)  {


	fmt.Println("Test : pushing message")
	ch<- true
	fmt.Println("Test: Pushed message")

	group.Done()

}

func test1(ch chan bool, group * sync.WaitGroup)  {
	fmt.Println("Test1 : pushing message")
	ch<- true
	fmt.Println("Test1 : Pushed message")
	group.Done()
}

func main()  {


	ch := make(chan bool)
	var wait sync.WaitGroup;

	wait.Add(2);


	go test(ch, &wait)
	go test1(ch, &wait)

	time.Sleep(2 * time.Second)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	wait.Wait()
}
