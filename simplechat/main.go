package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[*websocket.Conn]bool
}

func CreateServer() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (server *Server) handleWebSocket(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())
	server.connections[ws] = true
	server.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Read Error:", err)
			continue
		}

		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (server *Server) broadcast(data []byte) {
	for ws := range server.connections {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(data); err != nil {
				fmt.Println("Error: ", err)
			}
		}(ws)
	}
}

func main() {
	port := flag.String("port", "3000", "Enter port number")
	flag.Parse()

	server := CreateServer()
	http.Handle("/ws", websocket.Handler(server.handleWebSocket))
	http.ListenAndServe(":"+*port, nil)
}
