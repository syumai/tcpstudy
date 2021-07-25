package tcp

import (
	"net"
	"syscall"
	"time"
)

type tcpConn struct {
	fd int
}

var _ net.Conn = &tcpConn{}

func (c *tcpConn) Read(b []byte) (int, error) {
	return syscall.Read(c.fd, b)
}

func (c *tcpConn) Write(b []byte) (int, error) {
	return syscall.Write(c.fd, b)
}

func (c *tcpConn) Close() error {
	return syscall.Close(c.fd)
}

func (c *tcpConn) LocalAddr() net.Addr {
	// TODO: implement
	panic("unimplemented")
}

func (c *tcpConn) RemoteAddr() net.Addr {
	// TODO: implement
	panic("unimplemented")
}

func (c *tcpConn) SetDeadline(t time.Time) error {
	// TODO: implement
	panic("unimplemented")
}

func (c *tcpConn) SetReadDeadline(t time.Time) error {
	// TODO: implement
	panic("unimplemented")
}

func (c *tcpConn) SetWriteDeadline(t time.Time) error {
	// TODO: implement
	panic("unimplemented")
}
