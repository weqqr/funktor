package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/weqqr/funktor/pkg/wl"
)

func main() {
	xdgRuntimeDir := os.Getenv("XDG_RUNTIME_DIR")

	waylandDisplay := os.Getenv("WAYLAND_DISPLAY")
	socketPath := filepath.Join(xdgRuntimeDir, waylandDisplay)

	fakeWaylandDisplay := "wlspy-0"
	fakeSocketPath := filepath.Join(xdgRuntimeDir, fakeWaylandDisplay)

	err := os.RemoveAll(fakeSocketPath)
	if err != nil {
		panic(err)
	}

	server := wl.Server{
		SocketPath:  fakeSocketPath,
		ConnHandler: makeConnHandler(socketPath),
	}

	go func() {
		err := server.Listen()
		if err != nil {
			panic(err)
		}
	}()

	os.Setenv("WAYLAND_DISPLAY", fakeWaylandDisplay)

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func makeConnHandler(waylandDisplay string) func(wl.Connection) error {
	return func(serverConn wl.Connection) error {
		client, err := wl.NewClient(waylandDisplay)
		if err != nil {
			panic(err)
		}

		go func() {
			io.Copy(client, serverConn)
		}()

		go func() {
			io.Copy(serverConn, client)
		}()

		return nil
	}
}
