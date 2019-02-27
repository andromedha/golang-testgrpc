package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/andromedha/golang-testgrpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	name        string
	substream   map[string]pb.PushMetaData_SubscribeServer
	clientcount int
	streamers   map[int32]string
}

func newServer() *server {
	return &server{
		name:        "BackendServer",
		substream:   make(map[string]pb.PushMetaData_SubscribeServer),
		clientcount: 0,
		streamers:   make(map[int32]string)}
}

func (s *server) Register(in *pb.RegistrationRequest, stream pb.PushMetaData_RegisterServer) error {
	log.Printf("New Client received %v \n", in.GetClientnumber())
	s.clientcount = s.clientcount + 1
	return nil
}

func (s *server) Subscribe(stream pb.PushMetaData_SubscribeServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received Subscription from %d \n", in.Clientnumber)
		_, ok := s.substream[strconv.Itoa(int(in.GetClientnumber()))]
		if !ok {
			s.substream[strconv.Itoa(int(in.GetClientnumber()))] = stream
			s.streamers[in.GetClientnumber()] = in.GetName()
			s.PushMetaDataUpdates()
		}
		if ok {
			log.Printf("Client %d already subscribed \n", in.GetClientnumber())
		}
	}
}

func (s *server) PushMetaDataUpdates() {
	time.Sleep(1 * time.Second)
	for k, v := range s.substream {
		log.Printf("Push Notification to Client %v", k)
		err := v.Send(&pb.Notification{Streamers: s.streamers})
		if err != nil {
			log.Fatalf("Send failed %v", err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server started")
	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterPushMetaDataServer(s, myServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
