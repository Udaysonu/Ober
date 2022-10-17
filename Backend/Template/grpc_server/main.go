package main

import (
	"log"
	pb "github.com/udaysonu/ober/grpc_proto"
	"google.golang.org/grpc"
	"net"
)

var addr string= "localhost:5051"

type Server struct{
	pb.GreetServiceServer
}

func main(){
	log.Printf("Listening on %s\n",addr)

	lis,err:=net.Listen("tcp",addr)

	if err!=nil{
		log.Fatalf("Failed to listen on: %v\n",err)
	}

	log.Printf("Listening on %s\n",addr)

	s:=grpc.NewServer()
	pb.RegisterGreetServiceServer(s,&Server{})
	if err=s.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve: %v\n",err)
	}
}