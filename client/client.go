package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/andromedha/golang-testgrpc/proto"

	"google.golang.org/grpc"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// input = strings.Replace(input, "\n", "", -1)
	// number, err := strconv.Atoi(input)

	// if err != nil {
	// 	log.Fatal("Input can not convert to Int.")
	// }
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000)

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failto connect Server")
	}
	client := pb.NewPushMetaDataClient(conn)

	registerclient(client, number)

	subscribe(client, number)

	for {
	}
}

func registerclient(client pb.PushMetaDataClient, number int) error {
	_, err := client.Register(context.Background(), &pb.RegistrationRequest{Clientnumber: int32(number)})
	if err != nil {
		log.Fatal("Register failed")
		return err
	}
	return nil
}

func subscribe(client pb.PushMetaDataClient, number int) {
	stream, err := client.Subscribe(context.Background())

	if err != nil {
		log.Fatal("Subscribe failed")
	}

	err = stream.Send(&pb.Clientdata{Clientnumber: int32(number), Name: randomString(10)})
	if err != nil {
		log.Fatal("Subscribe send failed")
	}
	recvMetadata(stream)

	stream.CloseSend()
}

func recvMetadata(stream pb.PushMetaData_SubscribeClient) {
	for {
		resp, rerr := stream.Recv()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			log.Fatalf("Failed to received %v", rerr)
		}
		log.Print(resp.Streamers)

	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
