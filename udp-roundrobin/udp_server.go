package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ip := "localhost"
	port := 7000
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: port}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s", conn.LocalAddr())

	var requestCount int = 0
	buf := make([]byte, 1024)
	for {
		n, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("read from udp failed: %v", err)
		} else {
			requestCount++
			log.Printf("count %d received %d bytes from %s", requestCount, n, remote)
			log.Printf("%s", buf[:n])
		}

		// echo to client with requestCount
		host := os.Getenv("HOST")
		msg := fmt.Sprintf("%s requestCount: %d", host, requestCount)
		_, err = conn.WriteToUDP([]byte(msg), remote)
		if err != nil {
			log.Printf("write to udp failed: %v", err)
		} else {
			log.Printf("echo to %s", remote)
		}
	}
	log.Println("finish request count", requestCount)
}
