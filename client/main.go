package main

import (
    "context"
    "log"
    "os"
    "time"

    pb "golang-grpc/service"

    "google.golang.org/grpc"
)

const (
    address     = "localhost:50051"
    defaultName = "world"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    
    conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    }

    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}
