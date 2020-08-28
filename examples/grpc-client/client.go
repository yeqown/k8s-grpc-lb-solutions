package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/yeqown/k8s-grpc-lb-solutions/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/balancer/roundrobin"
)

var (
	address string
	timeout int

	c pb.GreeterClient
)

func init() {
	log.SetOutput(os.Stdout)

	flag.IntVar(&timeout, "timeout", 1, "greet rpc call timeout")
	flag.StringVar(&address, "address", "localhost:50051", "grpc server addr")
}

//var tplServiceConfig = `{"LoadBalancingPolicy": "%s","MethodConfig": [{"Name": [{"Service": "hello.Greeter"}], "RetryPolicy": {"MaxAttempts":2, "InitialBackoff": "0.1s", "MaxBackoff": "1s", "BackoffMultiplier": 2.0, "RetryableStatusCodes": ["UNAVAILABLE"]}}]}`
var tplServiceConfig = `{"LoadBalancingPolicy": "%s"}`

func bootstrap() error {
	flag.Parse()

	// Set resolver
	//resolver.SetDefaultScheme("custom_dns")

	log.Println("connecting to:", address)

	bc := backoff.DefaultConfig
	bc.MaxDelay = time.Second
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(tplServiceConfig, roundrobin.Name)),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: bc,
		}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c = pb.NewGreeterClient(conn)

	return nil
}

func main() {
	_ = bootstrap()

	log.Println("client running")

	// Contact the server and print out its response.
	for range time.Tick(2 * time.Second) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "defaultName"})
		if err != nil {
			log.Printf("could not greet: %v\n", err)
		} else {
			log.Printf("Greeting: %s", r.Message)
		}
		cancel()
	}
}
