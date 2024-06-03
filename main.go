package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	clientFlag, serverFlag bool
	port                   = "6543"
)

func init() {
	flag.BoolVar(&clientFlag, "client", false, "-client")
	flag.BoolVar(&serverFlag, "server", false, "-server")
}

func main() {
	flag.Parse()
	if clientFlag {
		runClient()
	} else {
		runServer()
	}
}

func runClient() {
	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to server at %s\n", addr)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server:", err.Error())
			return
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err.Error())
			return
		}
		fmt.Print("Server response: ", response)
	}
}

func runServer() {
	fmt.Println("welcome to remdics")
	// port := "6543"
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("Server is listening on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Print("Message received:", string(message))
		conn.Write([]byte(message))
	}
}
