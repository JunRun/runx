package rImpl

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type SeverConfig struct {
	Server map[string]Server
}
type Server struct {
	//sever name
	Name      string `yaml:"name"`
	IPVersion string `yaml:"ip_version"`
	IPAddress string `yaml:"ip_address"`
	Port      int    `yaml:"port"`
}

func NewServer(Name string) *Server {
	log.Println(Name)

	ser := &Server{
		Name:      Name,
		IPVersion: "tcp4",
		IPAddress: "0.0.0.0",
		Port:      9090,
	}
	return ser
}

func (s *Server) Sever() {
	s.Start()

	select {}
}
func (s *Server) Start() {
	go func() {
		fmt.Printf("[Start] Server Listen ip %s:%d", s.IPAddress, s.Port)

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))
		if err != nil {
			fmt.Println("Resolve TcpAddr Error")
		}

		listen, err := net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			fmt.Printf("Listen TcpAddr:%s:%d  Error", s.IPAddress, s.Port)
		}
		var connID uint64
		connID = 0
		for {
			accept, err := listen.AcceptTCP()

			if err != nil {
				fmt.Println("Accept TcpConnection error: ", err)
				continue
			}

			nc := NewConnection(accept, connID, CallBack)
			connID++
			go nc.Start()

		}
	}()

}

func (s *Server) Stop() {

}

func CallBack(Conn *net.TCPConn, bytes []byte, cnt int) error {

	fmt.Printf("recv message %s", bytes)
	fmt.Println("")
	if _, err := Conn.Write(bytes[:cnt]); err != nil {
		fmt.Println("write buf back err:", err)
		return errors.New("CallBack error")

	}
	return nil

}
