package main

import (
	"time"
	"fmt"
)

type CustomWaitGroupTest struct {

	threadCount int

	ch chan bool

}

func (w *CustomWaitGroupTest) add(i int)  {

	w.threadCount = w.threadCount + i
}


func (w *CustomWaitGroupTest) done()  {

	w.threadCount = w.threadCount - 1

	if w.ch != nil {
		w.ch<- true
	}
}

func (w CustomWaitGroupTest) waitLen()  int {

	return  w.threadCount
}

func (w *CustomWaitGroupTest) wait()  {

	w.ch = make(chan bool, w.threadCount)

	for i := w.threadCount; i > 0; i-- {
		<- w.ch
	}
}

func testFunc(wg *CustomWaitGroupTest)  {

	for  i := 0; i< 5; i++ {

		time.Sleep(2 * time.Second)
		fmt.Println("Task completed : ", i+1)
		wg.done()

	}
}

func main()  {

	var test CustomWaitGroupTest

	test.add(5)

	go testFunc(&test)

	test.wait()

	fmt.Println("All done")

}