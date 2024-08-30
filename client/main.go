package main

import (
    "context"
    "crypto/tls"
    "flag"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"

    "github.com/aaronjoju07/grpc/pb" // Ensure this path matches your protobuf package import
)

func main() {
    serverAddr := flag.String(
        "server", "localhost:8080",
        "the server address in the format of host:port",
    )
    flag.Parse()

    // Set up TLS credentials
    creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(creds),
    }

    // Set up a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create a connection to the server
    conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // Create a new Calculator client
    client := pb.NewCalculatorClient(conn)

    // Make the Sum RPC call
    res, err := client.Sum(ctx, &pb.NumberRequest{
        Number: []int64{10, 10, 20, 90, 70},
    })
    if err != nil {
        log.Fatalf("could not calculate sum: %v", err)
    }

    // Print the result
    log.Printf("Sum result: %v", res.GetResult())
}
