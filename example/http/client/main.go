package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	conn, err := net.DialTCP("127.0.0.1:8080", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := httputil.NewClientConn(conn, nil)

	const msg = "Hello, world!"

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/echo", strings.NewReader(msg))
	if err != nil {
		panic(err)
	}

	err = cc.Write(req)
	if err != nil {
		panic(err)
	}

	res, err := cc.Read(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if string(b) != msg {
		fmt.Printf("want: %s, got: %s", msg, string(b))
		os.Exit(1)
	}
}
