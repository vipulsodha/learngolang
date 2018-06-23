package main

import (
	"net"
	"log"
)


func tcpHandler (con net.Conn) {


	defer con.Close()

	hello := "hello world\n"

	con.Write([]byte(hello))



}


func main() {

	server, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatal("Error", err)
	}

	for {

		conn, err := server.Accept()
		if err == nil {
			go tcpHandler(conn)

		} else {
			log.Fatal("Error in connection")
		}
	}

}
