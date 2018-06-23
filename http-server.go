package main

import (
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r * http.Request) {

	//fmt.Fprintf(w, "hello http");


	w.Write([]byte("Hello world"))


}


func main() {

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8088", nil))
}
