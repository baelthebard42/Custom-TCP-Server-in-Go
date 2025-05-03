package main

import (
	"fmt"

	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{} //value doesnt matter, only the event(signal) matters. takes 0 bytes of memory
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.acceptLoop()
	<-s.quitch
	return nil
}

func (s *Server) acceptLoop() {

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading data into buffer", err)
			continue
		}

		msg := buf[:n]
		fmt.Println(string(msg))
	}

}

func main() {

	server := NewServer(":8000")
	fmt.Println(server.Start())

}
