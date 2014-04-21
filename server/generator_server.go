package server

import (
	"../estimate"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
)

type RPCFunc uint8

type FeatureParameter struct {
	name  string
	value float64
}

type FeatureParameterSet struct {
	method string
	params map[string]map[string]float64
}

var param_set = make(map[string]*FeatureParameterSet, 100)

func (*RPCFunc) Echo(arg *string, result *string) error {
	log.Print("Arg passed: " + *arg)
	*result = ">" + *arg + "<"
	return nil
}

func (*RPCFunc) Get(arg *string, result *string) error {
	args := strings.Split(*arg, ", ")
	stream_name := args[0]

	if _, key_exist := param_set[stream_name]; key_exist {
		json_map := map[string]interface{}{}
		switch param_set[stream_name].method {
		case "gaussian":
			for variable_name, params := range param_set[stream_name].params {
				mean := params["mean"]
				std := params["std"]
				json_map[variable_name] = rand.NormFloat64()*std + mean
			}
		default:
			fmt.Println("NON SUPPORTED METHOD")
		}
		json_string, _ := json.Marshal(json_map)
		*result = string(json_string)
	} else {
		log.Print("No existing stream name: " + stream_name)
		*result = "false"
	}
	return nil
}

func (*RPCFunc) Register(arg *string, result *string) error {
	log.Print("Arg passed: " + *arg)
	*result = ">" + *arg + "<"
	args := strings.Split(*arg, ", ")
	stream_name, method, file_name := args[0], args[1], args[2]
	param_set[stream_name] = new(FeatureParameterSet)
	param_set[stream_name].method = method

	switch method {
	case "gaussian":
		param_set[stream_name].params = estimate.EstimateGaussian(file_name)
	default:
		fmt.Println("NON SUPPORTED METHOD")
	}

	*result = "true"
	return nil
}

func main() {
	StartServer()
}

func StartServer() {
	log.Print("Starting Server...")
	l, err := net.Listen("tcp", "localhost:1234")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("listening on: ", l.Addr())
	rpc.Register(new(RPCFunc))
	for {
		log.Print("waiting for connections ...")
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %s", conn)
			continue
		}
		log.Printf("connection started: %v", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
