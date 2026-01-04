package wl

import (
	"fmt"
	"net"
)

type Server struct {
	SocketPath  string
	ConnHandler func(Connection) error
}

func (s *Server) Listen() error {
	listener, err := net.Listen("unix", s.SocketPath)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		if s.ConnHandler != nil {
			go s.ConnHandler(Connection{
				conn: conn,
			})
		}
	}
}

type Connection struct {
	conn net.Conn
}

func (c Connection) Read(data []byte) (int, error) {
	return c.conn.Read(data)
}

func (c Connection) Write(data []byte) (int, error) {
	fmt.Printf("write %v\n", len(data))

	return c.conn.Write(data)
}

func (c Connection) Close() error {
	return c.conn.Close()
}
