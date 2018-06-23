package main

import "fmt"

type address struct {
	pincode int
	lane string
}


type Person struct {
	name string
	age int
	address
}


func main()  {


	vipul := Person{
		"Vipul",
		42,
		address{
			12,
			"borivali",
		},
	}


	fmt.Println(vipul.address.pincode)



}
