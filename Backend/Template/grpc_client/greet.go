package main
import(
	"context"
	"log"
	pb "github.com/udaysonu/ober/grpc_proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res,err:=c.Greet(context.Background(),&pb.GreetRequest{
		FirstName:"Clement",
	})

	if err!=nil{
		log.Fatalf("could not greet: %v\n",err)
	}

	log.Printf("Greeting: %s\n",res.Result)
}