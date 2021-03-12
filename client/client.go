package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"

	pb "github.com/rakiasomai/Zenly/proto"

	"time"

	"google.golang.org/grpc"
)

var conn grpc.ClientConnInterface
var err error


// start user add function
func user_add() {
	client := pb.NewKVClient(conn)
	var k string
	var v string

	fmt.Printf("Write The key : ")
	fmt.Scanf("%s", &k)

	fmt.Printf("Write The value : ")
	fmt.Scanf("%s", &v)
	
	in := &pb.KeyValue{Key: k, Value: v}
	stream, err := client.Add(context.Background(), in)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	//ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Value)
		}
	}()
	<-done
	log.Printf("finished")
	stream.CloseSend()
}

// start user get function 
func user_get() {
	client := pb.NewKVClient(conn)
	var k string

	fmt.Printf("Write The key that you want to get : ")
	fmt.Scanf("%s", &k)

	
	in := &pb.KeyRequest{Key: k}
	
	stream, err := client.Get(context.Background(), in)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	//ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for {

			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Value)
		}
	}()
	<-done
	log.Printf("finished")
	main()
}



func main() {
	rand.Seed(time.Now().Unix())
	var c string
	fmt.Printf("choose the function : ")
	fmt.Scanf("%s", &c)

	// dail server
	conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	} else if c == "a" {
		user_get()
		} else if c == "b" {
		user_add()
		} else {
	
	}
}
