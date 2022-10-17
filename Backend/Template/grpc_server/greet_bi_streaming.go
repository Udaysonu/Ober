package main
import(
	"log"
	"io"
 	pb "github.com/udaysonu/ober/grpc_proto"
)


func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error{
	log.Println("GreetEveryone was invoked")

	for{
		req,err:=stream.Recv()
		log.Printf("%v %v\n",req,err)
		if err==io.EOF{
			return nil
		}

		if err!=nil{
			log.Fatalf("Error while reading client stream: %v\n",err);
		}

		res:="Hello "+ req.FirstName + "!"
		err = stream.SendMsg(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v",err);
		}


	}
}