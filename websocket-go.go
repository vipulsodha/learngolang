package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"log"
	"fmt"
	"time"
	"sync"
)

func handleReads(ws *websocket.Conn, wait * sync.WaitGroup)  {

	for  {

		var reply string;

		websocket.Message.Receive(ws, &reply)

		fmt.Println(reply)

		websocket.Message.Send(ws, "Got message buddy")

	}

	wait.Done()
}


func handleWs(ws *websocket.Conn)  {

	var wait sync.WaitGroup;

	wait.Add(2)

	go handleReads(ws, &wait)

	//go sendMessage(ws, &wait)

	wait.Wait()
}

func sendMessage(ws *websocket.Conn, wait *sync.WaitGroup)  {

	for {
		websocket.Message.Send(ws, "Timed messgage")
		time.Sleep(2 * time.Second)
	}

	wait.Done()
}


func main() {


	http.Handle("/", websocket.Handler(handleWs))

	log.Fatal(http.ListenAndServe(":8088", nil))

}
