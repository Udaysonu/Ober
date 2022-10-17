package main
import(
	"context"
	"log"
	"io"
	pb "github.com/udaysonu/ober/grpc_proto"
)

func doGeetManyTimes(c pb.GreetServiceClient){
	log.Println("doGreetManytimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	stream, err:= c.GreetManyTimes(context.Background(),req)

	if err!=nil{
		log.Fatal("Error while calling GreetManyTimes: %v\n",req)
	}

	for {
		msg,err:=stream.Recv()

		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatal("Error while reading the stream: %v\n",err)
		}

		log.Printf("GreetManyTimes: %s\n",msg.Result)
	}
}