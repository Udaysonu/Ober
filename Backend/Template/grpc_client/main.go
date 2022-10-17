package main
import(
	_ "context"
	"log"
	pb "github.com/udaysonu/ober/grpc_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

var addr string = "localhost:5051"

func main(){
	conn, err:=grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("Failed to connect: %v\n",err)
	}
	defer conn.Close()
	c:=pb.NewGreetServiceClient(conn)

	// doGeetManyTimes(c)
	doGreetEveryone(c)
}