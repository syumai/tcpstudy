package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	err := Pong(":8080")
	if err != nil {
		panic(err)
	}
}

func Pong(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	br := bufio.NewReader(conn)
	line, _, err := br.ReadLine()
	if err != nil {
		return err
	}

	if string(line) != "ping" {
		return fmt.Errorf("msg must be ping, got: %s", string(line))
	}
	fmt.Println(line)

	const msg = "pong\n"

	_, err = io.Copy(conn, strings.NewReader(msg))
	if err != nil {
		return err
	}
	return nil
}
