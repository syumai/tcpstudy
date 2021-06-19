package main

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	listener, err := net.ListenTCP("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		sc := httputil.NewServerConn(conn, nil)

		req, err := sc.Read()
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		b, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		rd := bytes.NewReader(b)

		res := &http.Response{
			Status:        "200 OK",
			StatusCode:    http.StatusOK,
			Proto:         "HTTP/1.0",
			ProtoMajor:    1,
			ProtoMinor:    0,
			Header:        http.Header{},
			Body:          io.NopCloser(rd),
			ContentLength: int64(len(b)),
			Close:         true,
			Request:       req,
		}

		err = sc.Write(req, res)
		if err != nil {
			panic(err)
		}
	}
}
