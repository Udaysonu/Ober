package main
import(
	"context"
	"log"
	"time"
 	pb "github.com/udaysonu/ober/grpc_proto"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Uday"},
		{FirstName: "Kiran"},
		{FirstName: "Bakka"},
	}

	stream, err:= c.LongGreet(context.Background())

	if err!=nil{
		log.Fatal("Error while calling GreetManyTimes: %v\n",reqs)
	}

	for _,req:=range reqs {
		log.Printf("Sending req: %v\n",req)
		stream.Send(req)
		time.Sleep(1*time.Second)
	}
	res,err:=stream.CloseAndRecv()

	if err!=nil{
		log.Fatalf("Error while receiving response fro LongGreet %v",err)
	}

	log.Printf("LongGreet: %s\n",res.Result)
}