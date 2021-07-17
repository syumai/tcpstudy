package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"net"
	"os"
	"syscall"
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

const ProtocolNumberTCP = 6

const ISN uint32 = 123456

func Ping(addr string) error {
	sockFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, ProtocolNumberTCP)
	if err != nil {
		return err
	}
	defer syscall.Close(sockFD)

	sockAddr := syscall.SockaddrInet4{
		Port: 56789,
		Addr: [4]byte{}, // zero value
	}
	if err := syscall.Bind(sockFD, sockAddr); err != nil {
		return err
	}

	ph := &tcp.PseudoHeader{
		PTCL:      ProtocolNumberTCP, // tcp
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
