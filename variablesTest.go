package main

import "fmt"

func main()  {

	//const test =5
	//
	//fmt.Println("Value : ", test)
	//fmt.Printf("Type %T", test)
	//fmt.Println("")
	//
	//var a [3]int
	//
	//fmt.Println(len(a))
	//
	//
	//var atest []int;
	//
	//var aarray [5]int
	//
	//fmt.Printf("Type %T", atest)
	//fmt.Println("")
	//fmt.Printf("Type %T", aarray)



	var a =  [...]int{1,2,3,4,5,6,7,8}



	var bigSlice = a[:]

	bigSlice = append(bigSlice, 100,200)

	fmt.Println("Capacity bigSlice : ", cap(bigSlice))
	fmt.Println("Capacity a :", cap(a))

	bigSlice[1]=99;

	fmt.Println(a)
	fmt.Println(bigSlice)

	//var aSlice = a[:3]

	//newSlice := append(aSlice, 100,20,300)
	//
	//fmt.Println(newSlice)
	//
	////newSlice[1]=99
	//
	//
	//fmt.Println(newSlice)
	//
	//fmt.Println(aSlice)
	//
	//fmt.Println("Capacity a: ", cap(a))
	//
	//fmt.Println("Capacity newSlice: ", cap(newSlice))

	//fmt.Println("Capacity aSlice: ", cap(aSlice))

	//var len10 [10]int;
	//var lenTen [10]int;
	//
	//len10 = lenTen;
	//
	//fmt.Println(len10)




}
