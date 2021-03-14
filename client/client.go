package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/rakiasomai/Zenly/proto"

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

	func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)			
			} 
			log.Printf("Resp received: %s", resp.Value)
			break
		}
		log.Printf("finished")
	}()
	
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
    func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}	
			log.Printf("Resp received: %s", resp.Value)
			break
		}
		log.Printf("finished")
		
	}()
}


func choice() {
	var c string
	fmt.Printf("Choose the function add or get : ")
	fmt.Scanf("%s", &c)
	switch {
	case c == "add":
		user_add()
	case c == "get":
		user_get()
	}
}

func main() {   
	// dail server
	conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	} 
	choice()
}
