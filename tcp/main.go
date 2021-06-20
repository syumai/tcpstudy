package main

import (
	"bytes"
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

const ISN uint32 = 123456

//+--------+--------+--------+--------+
//|           Source Address          |
//+--------+--------+--------+--------+
//|         Destination Address       |
//+--------+--------+--------+--------+
//|  zero  |  PTCL  |    TCP Length   |
//+--------+--------+--------+--------+

type PseudoHeader struct {
	SourceAddress      uint32
	DestinationAddress uint32
	_                  uint8
	PTCL               uint8
	TCPLength          uint16
}

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

	ph := &PseudoHeader{
		PTCL:      6, // tcp
		TCPLength: 8, // TODO: temporary value
	}

	b := bytes.NewReader([]byte(net.ParseIP("127.0.0.1")))
	if err := binary.Read(b, binary.BigEndian, &ph.SourceAddress); err != nil {
		return err
	}
	ph.DestinationAddress = ph.SourceAddress

	h := &Header{
		SourcePort:            49443,
		DestinationPort:       8080,
		SequenceNumber:        ISN + 0,
		Acknowledgment:        0,
		DataOffsetControlBits: SYN,
		Window:                65535, // temporary value
		//Checksum              : ,
		//UrgentPointer         : ,
		//Options               : ,
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
