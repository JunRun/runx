package rImpl

import (
	"fmt"
	"github.com/runx/server/rIterface"
	"github.com/runx/util"
	"log"
	"net"
)

type Server struct {
	//sever name
	Name      string `yaml:"name"`
	IPVersion string `yaml:"ip_version"`
	IPAddress string `yaml:"ip_address"`
	Port      int    `yaml:"port"`
	Router    rIterface.IRouter
}

func NewServer() *Server {
	log.Println(util.Config.GetString("server.name"))

	ser := &Server{
		Name:      util.Config.GetString("server.name"),
		IPVersion: util.Config.GetString("server.ip_version"),
		IPAddress: util.Config.GetString("sever.ip_address"),
		Port:      util.Config.GetInt("server.port"),
		Router:    nil,
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

			nc := NewConnection(accept, connID, s.Router)
			connID++
			go nc.Start()

		}
	}()

}
func (s *Server) SendMsg(Id uint64, data []byte) {

}

func (s *Server) Stop() {

}

func (s *Server) AddRouter(router rIterface.IRouter) {
	s.Router = router
}
