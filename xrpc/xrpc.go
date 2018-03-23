package main

import (
	"log"
	"net"
	"time"
	"net/rpc"
	"net/http"
)

type Echo int


func (e *Echo) Hi(args string, reply *string) error {
	*reply = "echo: " + args
	return nil
}


func main() {
	rpc.Register(new(Echo))
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)

	time.Sleep(2 * time.Second)

	c, e := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
	if e != nil {
		log.Fatal("client dail error:", e)
	}

	reqs := "a"
	var resp string

	c.Call("Echo.Hi", reqs, &resp)

	log.Println("Server response: ", resp)
}