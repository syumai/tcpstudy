package main

import (
	"io"
	"net"
	"os"
)

func main() {
	err := Ping("127.0.0.1")
	if err != nil {
		panic(err)
	}
}

func Ping(addr string) error {
	conn, err := net.Dial("ip4:1", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	const msg = "ping\n"

	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
