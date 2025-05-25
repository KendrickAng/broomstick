package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		fmt.Println("Received interrupt signal, shutting down...")
		cancel()
		time.Sleep(3 * time.Second) // allow other goroutines to finish
		os.Exit(0)
	}()

	for {
		select {
		case <-ctx.Done():
			// If shutdown initiated, don't accept new connections
			continue
		default:
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleConnection(ctx, conn)
		}
	}
}

func handleConnection(_ context.Context, conn net.Conn) {
	defer conn.Close()

	// Do some authentication here...

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	fmt.Println("Received data:", string(buf[:n]))

	conn.Write([]byte("Hello from BFF!"))

	fmt.Println("Handled connection from:", conn.RemoteAddr())
}
