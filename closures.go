package main

import "fmt"

func testClosure(a int) func() {

	b:= 10

	test := func() {


		fmt.Println(a+b)
	}


	return test;
}



func main()  {


	b:= testClosure(5)

	b()

}
