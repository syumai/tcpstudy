package main

import (
	"encoding/binary"
	"io"
	"net"
	"os"
)

// https://play.golang.org/p/p7TwQleOGA3

/*
    0                   1                   2                   3
    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |          Source Port          |       Destination Port        |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                        Sequence Number                        |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                    Acknowledgment Number                      |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |  Data |           |U|A|P|R|S|F|                               |
   | Offset| Reserved  |R|C|S|S|Y|I|            Window             |
   |       |           |G|K|H|T|N|N|                               |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |           Checksum            |         Urgent Pointer        |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                    Options                    |    Padding    |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                             data                              |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/

type Header struct {
	SourcePort            uint16
	DestinationPort       uint16
	SequenceNumber        uint32
	Acknowledgment        uint32
	DataOffsetControlBits uint16
	Window                uint16
	Checksum              uint16
	UrgentPointer         uint16
	Options               uint32
}

func (h *Header) DataOffset() uint16 {
	return h.DataOffsetControlBits >> 12
}

func (h *Header) SetDataOffset() {
	h.DataOffsetControlBits = h.DataOffsetControlBits&0x0fff | 0x8000
}

func (h *Header) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, h)
}

const (
	FIN uint16 = 1 << iota
	SYN
	RST
	PSH
	ACK
	URG
)

func main() {
	err := Ping("127.0.0.1")
	if err != nil {
		panic(err)
	}
}

func Ping(addr string) error {
	conn, err := net.Dial("ip4:6", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	h := &Header{
		SourcePort        : 49443,
	DestinationPort       : 8080,
	SequenceNumber        : 0,
	Acknowledgment        : ,
	DataOffsetControlBits : ,
	Window                : ,
	Checksum              : ,
	UrgentPointer         : ,
	Options               : ,
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
