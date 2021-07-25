package tcp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"syscall"
)

// parseAddr parses given `ipv4:port` style address into IPv4 address bytes and port.
// ex) `127.0.0.1:56789` => `[4]byte{127, 0, 0, 1}, 56789`
func parseAddr(addr string) (ipV4Addr [4]byte, port int, err error) {
	hostPort := strings.Split(addr, ":")
	hostStr := hostPort[0]
	portStr := hostPort[1]

	ipV4AddrNums := strings.Split(hostStr, ".")
	if len(ipV4AddrNums) != 4 {
		return [4]byte{}, 0, fmt.Errorf("invalid IPv4Addr given: %v", hostStr)
	}
	for i, v := range ipV4AddrNums {
		n, err := strconv.Atoi(v)
		if err != nil {
			return [4]byte{}, 0, fmt.Errorf("invalid IPv4Addr number given: %w", err)
		}
		ipV4Addr[i] = byte(n)
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return [4]byte{}, 0, fmt.Errorf("invalid port number given: %w", err)
	}
	return
}

func Dial(addr string) (conn net.Conn, sourcePort int, err error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, ProtocolNumberTCP)
	if err != nil {
		return nil, 0, err
	}

	// bind source port
	sourcePort = 56789
	for {
		sockAddr := syscall.SockaddrInet4{
			Port: sourcePort,
			Addr: [4]byte{}, // zero value
		}
		err := syscall.Bind(fd, sockAddr)
		if err != nil {
			sourcePort++
			continue
		}
		break
	}
	return &tcpConn{fd: fd}, sourcePort, nil
}
