package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"net"
	"os"
	"unsafe"

	"github.com/syumai/tcpstudy/tcp"
)

// https://play.golang.org/p/p7TwQleOGA3

func main() {
	err := Ping("127.0.0.1")
	if err != nil {
		panic(err)
	}
}

const ISN uint32 = 123456

func Ping(addr string) error {
	conn, sourcePort, err := tcp.Dial()
	ph := &tcp.PseudoHeader{
		PTCL:      tcp.ProtocolNumberTCP, // tcp
		TCPLength: uint16(unsafe.Sizeof(tcp.Header{})),
	}

	b := bytes.NewReader([]byte(net.ParseIP("127.0.0.1")))
	if err := binary.Read(b, binary.BigEndian, &ph.SourceAddress); err != nil {
		return err
	}
	ph.DestinationAddress = ph.SourceAddress

	h := tcp.Header{
		SourcePort:            49443,
		DestinationPort:       8080,
		SequenceNumber:        ISN + 0,
		Acknowledgment:        0,
		DataOffsetControlBits: tcp.SYN,
		Window:                math.MaxUint16, // temporary value
		Checksum:              0,
		UrgentPointer:         0,
		Options:               0,
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
