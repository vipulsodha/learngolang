package simpleSync


type CustomWaitGroupTest struct {

	threadCount int

	ch chan bool

}

func (w *CustomWaitGroupTest) Add(i int)  {

	w.threadCount = w.threadCount + i
}


func (w *CustomWaitGroupTest) Done()  {

	w.threadCount = w.threadCount - 1

	if w.ch != nil {
		w.ch<- true
	}
}

func (w CustomWaitGroupTest) WaitLen()  int {

	return  w.threadCount
}

func (w *CustomWaitGroupTest) Wait()  {

	w.ch = make(chan bool, w.threadCount)

	for i := w.threadCount; i > 0; i-- {
		<- w.ch
	}
}

