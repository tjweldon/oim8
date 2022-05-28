package peer

import (
	"bufio"
	"fmt"
	"net"
)

// Server manages the receipt of messages
type Server struct {
	writer *bufio.Writer
	conn   net.PacketConn
}

// NewServer initialises the server with a PacketConn and the bufio.Writer provided
func NewServer(writer *bufio.Writer) (*Server, error) {
	pc, err := net.ListenPacket("udp", ":42069")
	if err != nil {
		return nil, err
	}
	s := &Server{
		writer: writer,
		conn:   pc,
	}

	return s, nil
}

// Close closes the underlying net.PacketConn
func (s *Server) Close() {
	s.conn.Close()
}

func (s *Server) Serve() error {
	// main loop
	for {
		// read incoming messages
		buf := make([]byte, 1024)
		n, addr, err := s.conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		// this gorutine is the worker that serves incoming messages.
		go s.handleMsg(addr, buf[:n])
	}
}

// handleMsg handles an incoming message
func (s *Server) handleMsg(addr net.Addr, buf []byte) {
	s.conn.WriteTo(buf, addr)

	_, err := s.writer.Write(buf)
	if err != nil {
		fmt.Printf("Encountered error [%s] wile handling message: %s\n", err, buf)
		return
	}

	s.writer.Flush()
}
