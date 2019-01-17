package main

import (
	"log"
	"net"
	"time"
)

type Server struct{}

func (Server) Serve(ready chan bool) {
	time.Sleep(3000 * time.Millisecond)
	log.Println("Start Serve")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	ready <- true
	log.Println("Listen: 8080")
	_, err = ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
