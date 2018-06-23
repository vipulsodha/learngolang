package main

import (
	"sync"
	"fmt"
)

var sharedResource = 0;

func incrementMutex(m *sync.Mutex, wait *sync.WaitGroup)  {

	m.Lock()

	sharedResource = sharedResource + 1;

	m.Unlock()

	wait.Done()
}



func main()  {

	var m sync.Mutex;
	var wait sync.WaitGroup;

	for i:=0;i<1000 ;i++  {
		wait.Add(1)
		go incrementMutex(&m, &wait)
	}

	wait.Wait()

	fmt.Println(sharedResource)
}