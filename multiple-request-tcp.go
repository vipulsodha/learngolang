package main

import (
	"net"
	"log"
	"fmt"
)

func handleMutipleRequest(con net.Conn)  {
	defer con.Close()


	data := make([]byte, 1028)

	con.Write([]byte("HI Please input\n"))

		_, err := con.Read(data)

		if err != nil {
			fmt.Println("ERror")
		} else {
			fmt.Println(string(data))
		}

}


func main()  {




	server, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatal("Error", err)
	}

	for {

		conn, err := server.Accept()
		if err == nil {
			go handleMutipleRequest(conn)

		} else {
			log.Fatal("Error in connection")
		}
	}
}
