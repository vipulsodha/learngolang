package main

import "fmt"

func main()  {

	fmt.Println("hello world")


	var age int  = 42


	fmt.Println(age)



	var testAge = 42;


	fmt.Println(testAge)


	// short hand assignment

	name := "Vipul"

	fmt.Println(name)


	//multiple assignment

	var newAge, phone = 11, 1122344;

	fmt.Println(newAge, " " , phone);

	var (
		newName  = "rohan"
		rohanAge = 15
	)

	fmt.Println(rohanAge, newName)



	helloAge, helloName := 42, "Vipul"

	fmt.Println(helloAge, helloName)

}