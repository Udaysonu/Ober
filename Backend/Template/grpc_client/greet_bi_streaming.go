package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/udaysonu/ober/grpc_proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Println("doGreetEveryone was invoked")

	stream, err:= c.GreetEveryone(context.Background())

	if err!=nil{
		log.Fatalf("Error while creating stream: %v\n",err);
	}

	reqs:=[]*pb.GreetRequest{
		{FirstName: "Bakka"},
		{FirstName: "Uday"},
		{FirstName: "Test"},
	}

	waitc:= make(chan struct{})

	go func() {
		for _,req:=range reqs{
			log.Printf("Send request: %v\n",req)
			stream.Send(req)
			time.Sleep(1*time.Second)
			log.Printf("-------sending %v\n",req)
		}

		stream.CloseSend()
	}()

	go func(){
		for {
			res,err := stream.CloseAndRecv()

			if err == io.EOF{
				break
			}

			if err!=nil{
				log.Printf("Error while receiving: %v\n",err)
				break
			}

			log.Printf("Received: %v\n",res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
