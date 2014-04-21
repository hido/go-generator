package main

// Connect to JSONRPC Server and send command-line args to Echo

import (
	"./server"
	"flag"
	"fmt"
	amqp "github.com/streadway/amqp"
	"net"
	"net/rpc/jsonrpc"
	"strconv"
	"time"
)

func main() {
	var stream = flag.String("stream", "normal", "Name of stream in MQ")
	var speed = flag.Int("speed", 1, "Number of samples per second")
	var count = flag.Int("count", 100, "Number of samples to be generated")
	var filename = flag.String("filename", "test.csv", "Name of csv file")
	flag.Parse()

	conn, e := net.Dial("tcp", "localhost:1234")
	if e != nil {
		fmt.Println("No server running")
		go server.StartServer()
		time.Sleep(time.Second)
		conn, e = net.Dial("tcp", "localhost:1234")
	}
	client := jsonrpc.NewClient(conn)

	mq_connect, err := amqp.Dial("amqp://jubatus:jubatus@localhost")
	if err != nil {
		fmt.Println("Error!: cannot connect to MQ")
	}
	defer mq_connect.Close()

	arg_register := *stream + ", " + "gaussian" + ", " + *filename
	var reply_register string
	client.Call("RPCFunc.Register", arg_register, &reply_register)

	mq_channel, err2 := mq_connect.Channel()
	defer mq_channel.Close()
	if err2 != nil {
		fmt.Println("Error!: cannot create MQ channel")
	}

	for i := 0; i < *count; i++ {
		arg_get := *stream
		var row string
		client.Call("RPCFunc.Get", arg_get, &row)
		body := strconv.Itoa(i) + " " + row
		fmt.Println(body)
		msg := amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			ContentType:  "text/plain",
			Body:         []byte(body),
		}
		err_msg := mq_channel.Publish("", *stream, false, false, msg)
		if err_msg != nil {
			fmt.Println("Error!: ", err_msg)
		}
		time.Sleep(time.Second / time.Duration(*speed))
	}
}
