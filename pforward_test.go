package pforward

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func TestForward(t *testing.T) {
	go localServer()

	Forward(ForwardCnf{serverPort: 9001,
		destHost: "localhost",
		destPort: 9000})

	conn, err := net.DialTimeout(
		"tcp",
		"127.0.0.1:4000",
		time.Second*2)
	if err != nil {
		t.Fatalf("dial error: %v", err)
	}

	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buff := bytes.Buffer{}
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(&buff)

	io.Copy(writer, reader)

	if !strings.Contains(buff.String(), "local server") {
		t.Errorf("connection was not forwarded")
	}
}

func localServer() net.Listener {
	listener, _ := net.Listen("tcp", "localhost:9000")

	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go func(conn net.Conn) {
			buff := make([]byte, 1024)
			conn.Read(buff)

			conn.Write([]byte("local server"))
			conn.Close()
		}(conn)
	}
}
