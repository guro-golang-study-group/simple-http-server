package main

import (
	"log"
	"net"
	"testing"
)

// Server accept
func TestMain(t *testing.T) {
	ready(t)
	_, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func ready(t *testing.T) {
	server := Server{}
	readyChannel := make(chan bool)
	go server.Serve(readyChannel)
	if !<-readyChannel {
		t.Fail()
	}
}
