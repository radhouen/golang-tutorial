package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func createConnection()  {
	con,err := net.Listen(CONN_TYPE,CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer con.Close() // Close the listener when the application closes.
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for   {
		req,err := con.Accept()
		if err != nil {

		}
			go handleRequest(req)
	}

}
// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	fmt.Println("reqlenght", reqLen)
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
func main()  {
	createConnection()

}
