package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/yeqown/k8s-grpc-lb-solutions/proto"

	"google.golang.org/grpc"
)

var (
	d    time.Duration
	port string
)

// server is used to implement hello.GreeterServer.
type server struct{}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("i'm called")

	time.Sleep(d)
	return &pb.HelloReply{Message: "Hello " + in.Name + "! From " + ip()}, nil
}

func main() {
	flag.StringVar(&port, "port", ":50051", "default port")
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.Println("server running")

	// simulate busy server
	delay := rand.NewSource(time.Now().UnixNano()).Int63() % 2
	d = time.Duration(delay) * time.Second

	if d == time.Second {
		log.Println("I will stuck one second!!!")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server is running: %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func ip() string {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}
			if ip.String() != "127.0.0.1" {
				return ip.String()
			}
		}
	}
	return ""
}
