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
	Name          string `yaml:"name"`
	IPVersion     string `yaml:"ip_version"`
	IPAddress     string `yaml:"ip_address"`
	Port          int    `yaml:"port"`
	RouterHandler rIterface.IMsgHandler
}

func NewServer() *Server {
	log.Println(util.Config.GetString("server.name"))

	ser := &Server{
		Name:          util.Config.GetString("server.name"),
		IPVersion:     util.Config.GetString("server.ip-version"),
		IPAddress:     util.Config.GetString("sever.ip-address"),
		Port:          util.Config.GetInt("server.port"),
		RouterHandler: NewMsgHandler(),
	}
	return ser
}

func (s *Server) Sever() {
	s.Start()

	select {}
}
func (s *Server) Start() {
	go func() {
		//开启工作池
		s.RouterHandler.StartWorkPool()
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

			nc := NewConnection(accept, connID, s.RouterHandler)
			connID++
			go nc.Start()

		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) AddRouter(Id uint64, router rIterface.IRouter) {
	s.RouterHandler.AddRouter(Id, router)
}
