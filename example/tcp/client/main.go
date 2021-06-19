package main

import (
	"io"
	"net"
	"os"
)

func main() {
	err := Ping(":8080")
	if err != nil {
		panic(err)
	}
}

func Ping(addr string) error {
	conn, err := net.Dial("tcp", addr)
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
