package wl

import (
	"fmt"
	"net"
)

type Client struct {
	Connection
}

func NewClient(socketPath string) (*Client, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	return &Client{
		Connection: Connection{
			conn: conn,
		},
	}, nil
}
