package main

import (
	"sync"
	"fmt"
)

var commonResource int = 0;

func testRace(wait *sync.WaitGroup)  {

	commonResource = commonResource + 1;
	wait.Done()

}

func main()  {

	var wait sync.WaitGroup;
	for i:=0;i<1000 ;i++  {
		wait.Add(1)
		go testRace(&wait)
	}

	wait.Wait()

	fmt.Println(commonResource)

}
