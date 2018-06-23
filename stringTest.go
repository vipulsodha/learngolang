package main

import "fmt"

func main()  {


	var test string = "Señor";

	fmt.Println("Length of stirng : ", len(test))
	for i:=0;i<len(test) ;i++  {

		fmt.Printf("%c", test[i])
		fmt.Println()
	}

	//runes := []rune(test)
	//fmt.Println("------")
	//
	//fmt.Println("Runes length : ", len(runes))
	//
	//for i:=0;i<len(runes) ;i++  {
	//
	//	fmt.Printf("%x", runes[i])
	//	fmt.Println()
	//}


}
