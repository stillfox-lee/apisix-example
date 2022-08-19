package main

import (
	"log"
	"net"
)

func main() {

	ip := "localhost"
	port := 9200
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: port}

	buf := make([]byte, 1024)

	msg := "Hello, World!"
	for i := 0; i < 10; i++ {
		conn, err := net.DialUDP("udp", nil, &addr)
		if err != nil {
			log.Printf("connection to %s %d failed %s\n", ip, port, err)
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Printf("write to %s failed %s\n", ip, err)
			continue
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read from %s failed %s\n", ip, err)
			continue
		}
		log.Printf("received %d bytes: %s\n", n, buf[:n])
		conn.Close()
	}
}
