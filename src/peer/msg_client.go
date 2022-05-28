package peer

import (
	"log"
	"net"
)

// Client is the service that sends outgoing message to a single server.
type Client struct {
	msgs   <-chan []byte
	server *net.UDPAddr
}

// NewClient initialises a client with the resolved net.UDPAddr of the
// host and port supplied. Messages are sent via the msgs channel and the
// Client will stop when the msgs channel is closed.
func NewClient(host, port string, msgs <-chan []byte) *Client {
	// resolve the supplied host and port to a udp addr
	server, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	// use that to initialse the Client
	c := Client{
		msgs:   msgs,
		server: server,
	}

	return &c
}

// Run is the main goroutine for the client. It ranges over messages and
// sends them via UDP to the server that is supplied on instantiation.
func (c *Client) Run() {
	// main loop
	for msg := range c.msgs {
		// create a per-message connection
		conn, err := net.DialUDP("udp", nil, c.server)
		if err != nil {
			log.Fatal(err)
		}

		// write the message to the connection
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}

		// read the response if any
		buffer := make([]byte, 1024)
		_, _, err = conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		// close the connection for this message.
		_ = conn.Close()
	}
}
