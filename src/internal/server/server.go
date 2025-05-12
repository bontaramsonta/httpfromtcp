package server

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"
)

type Server struct {
	listener net.Listener
	closed   atomic.Bool
}

func Serve(port int) (*Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := &Server{listener: l}
	go s.listen()
	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true)
	return s.listener.Close()
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handle(conn)
	}
}

const response = `HTTP/1.1 200 OK
Content-Type: text/plain

Hello World!`

func (s *Server) handle(conn net.Conn) {
	serverClosed := s.closed.Load()
	if serverClosed {
		return
	}
	conn.Write([]byte(response))
}
