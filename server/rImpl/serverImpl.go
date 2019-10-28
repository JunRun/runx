package rImpl

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	//sever name
	name      string
	IPVersion string
	IPAddress string
	Port      int
}

func NewServer(Name string) *Server {
	log.Println(Name)
	ser := &Server{
		name:      Name,
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
		for {
			accept, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept TcpConnection error: ", err)
				continue
			}
			go func() {
				for {
					bytes := make([]byte, 512)
					cnt, err := accept.Read(bytes)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}

					fmt.Printf("recv message %s", bytes)

					if _, err := accept.Write(bytes[:cnt]); err != nil {
						fmt.Println("write buf back err:", err)
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {

}
