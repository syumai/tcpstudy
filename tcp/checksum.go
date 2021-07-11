package tcp

import (
	"bytes"
	"encoding/binary"
	"io"
)

/*
  +--------+--------+--------+--------+
  |           Source Address          |
  +--------+--------+--------+--------+
  |         Destination Address       |
  +--------+--------+--------+--------+
  |  zero  |  PTCL  |    TCP Length   |
  +--------+--------+--------+--------+
*/
type PseudoHeader struct {
	SourceAddress      uint32
	DestinationAddress uint32
	_                  uint8
	PTCL               uint8
	TCPLength          uint16
}

func CalculateChecksum(ph *PseudoHeader, h *Header, data []byte) (uint16, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, ph); err != nil {
		return 0, err
	}
	if err := binary.Write(&buf, binary.BigEndian, h); err != nil {
		return 0, err
	}
	if _, err := io.Copy(&buf, bytes.NewReader(data)); err != nil {
		return 0, err
	}
	var (
		sum uint16
		n   int
		err error
		b   []byte = make([]byte, 2)
	)
	for ; err != io.EOF; n, err = buf.Read(b) {
		if err != nil {
			return 0, err
		}
		if n == 1 {
			b[1] = 0 // pad zero
		}
		var i uint16
		err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
		if err != nil {
			return 0, err
		}
		sum += i
	}
}
