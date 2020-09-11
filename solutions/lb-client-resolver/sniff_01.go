package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/yeqown/k8s-grpc-lb-solutions/resolver"
)

var address01 string

func init() {
	flag.StringVar(&address01, "address01", "passthrough", "grpc server addr")
}

func sniffTest() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			got := resolver.Sniff(address01)
			fmt.Println("sniff remote: ", got)
		}
	}()
}

func main() {
	flag.Parse()

	fmt.Println("sniffTest running")
	sniffTest()

	select {}
}
