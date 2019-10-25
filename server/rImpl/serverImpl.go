package rImpl

import "log"

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
		IPVersion: "ip-v01",
		IPAddress: "0.0.0.0",
		Port:      9090,
	}
	return ser
}

func (s *Server) Sever() {
	s.Start()
}
func (s *Server) Start() {

}

func (s *Server) Stop() {

}
