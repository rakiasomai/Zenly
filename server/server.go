package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/rakiasomai/Zenly/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s server) Add(in *pb.KeyValue, srv pb.KV_AddServer) error {
	log.Printf("fetch response for key : %d", in.Key)
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func(count int64) {

			resp := pb.Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Key)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}
	wg.Wait()
	return nil
}

}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to connect on port 9000: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}