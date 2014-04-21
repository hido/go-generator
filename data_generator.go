package main

// Connect to JSONRPC Server and send command-line args to Echo

import (
	"./server"
	"flag"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	var stream = flag.String("stream", "normal", "Name of stream in MQ")
	var speed = flag.Int("speed", 1, "Number of samples per second")
	var count = flag.Int("count", 100, "Number of samples to be generated")
	var filename = flag.String("filename", "test.csv", "Name of csv file")
	var seed = flag.Int("seed", 0, "Rand seed")
	var silent = flag.Bool("silent", false, "Rand seed")
	flag.Parse()
	fmt.Println(*stream, *filename, *count, *speed, *seed, *silent)

	conn, e := net.Dial("tcp", "localhost:1234")
	if e != nil {
		go server.StartServer()
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
