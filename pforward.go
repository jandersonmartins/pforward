package pforward

import (
	"fmt"
	"io"
	"log"
	"net"
)

type ForwardCnf struct {
	serverPort int32
	destHost   string
	destPort   int32
}

func Forward(cnf ForwardCnf) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cnf.serverPort))
	if err != nil {
		log.Fatalf("listener error: %v", err)
	}
	destAddress := fmt.Sprintf("%s:%d", cnf.destHost, cnf.destPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connection error: %v", err)
		} else {
			go handleConnection(conn, destAddress)
		}
	}
}

func handleConnection(conn net.Conn, destAddress string) {
	dest, err := net.Dial("tcp", destAddress)
	defer conn.Close()
	if err != nil {
		log.Printf("dest (%s) conn error: %v", destAddress, err)
	} else {
		defer dest.Close()
		go func() {
			if _, err := io.Copy(dest, conn); err != nil {
				log.Printf("can't copy to dest: %v", err)
			}
		}()
		if _, err := io.Copy(conn, dest); err != nil {
			log.Printf("can't copy data: %v", err)
		}
	}
}
