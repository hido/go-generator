package main

// Connect to JSONRPC Server and send command-line args to Echo

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:1234")
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not connect: %s\n", e)
		os.Exit(1)
	}
	client := jsonrpc.NewClient(conn)
	var reply string
	var arg string
	for _, s := range os.Args[2:] {
		arg += s + ", "
	}
	fmt.Printf("Sending: %s\n", arg)
	client.Call("RPCFunc."+os.Args[1], arg, &reply)
	fmt.Printf("Reply: %s\n", reply)
}
