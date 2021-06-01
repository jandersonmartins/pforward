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

const (
	TEST_SERVER_ADDRESS  = "127.0.0.1:9000"
	TEST_SERVER_RESPONSE = "local server response"
)

func TestForward(t *testing.T) {
	go localServer()

	go Forward(ForwardCnf{serverPort: 9001,
		destHost: "127.0.0.1",
		destPort: 9000})

	conn, err := net.DialTimeout(
		"tcp",
		TEST_SERVER_ADDRESS,
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

	if !strings.Contains(buff.String(), TEST_SERVER_RESPONSE) {
		t.Errorf("connection was not forwarded")
	}
}

func localServer() net.Listener {
	listener, _ := net.Listen("tcp", TEST_SERVER_ADDRESS)

	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go func(conn net.Conn) {
			buff := make([]byte, 1024)
			conn.Read(buff)

			conn.Write([]byte(TEST_SERVER_RESPONSE))
			conn.Close()
		}(conn)
	}
}
