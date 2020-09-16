package rImpl

import (
	"fmt"
	"log"
	"net"
	"runx/server/rIterface"
	"runx/util"
)

type Server struct {
	//sever name
	Name          string `yaml:"name"`
	IPVersion     string `yaml:"ip-version"`
	IPAddress     string `yaml:"ip-address"`
	Port          int    `yaml:"port"`
	RouterHandler rIterface.IMsgHandler
	ConnMan       rIterface.IConnManger
	OnStartFunc   func(connection rIterface.IConnection)
	OnStopFunc    func(connection rIterface.IConnection)
}

func NewServer() *Server {
	log.Println(util.Config.GetString("server.name"))

	ser := &Server{
		Name:          util.Config.GetString("server.name"),
		IPVersion:     util.Config.GetString("server.ip-version"),
		IPAddress:     util.Config.GetString("sever.ip-address"),
		Port:          util.Config.GetInt("server.port"),
		RouterHandler: NewMsgHandler(),
		ConnMan:       NewConnManger(),
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
			//判断链接是否达到 最大链接数
			if s.ConnMan.Len() > util.Config.GetInt("server.max-connection") {
				accept.Close()
				continue
			}
			nc := NewConnection(s, accept, connID, s.RouterHandler)
			connID++
			go nc.Start()

		}
	}()

}

func (s *Server) Stop() {
	s.ConnMan.Clear()
	fmt.Println("[Stop] the server stop")
}

func (s *Server) AddRouter(Id uint64, router rIterface.IRouter) {
	s.RouterHandler.AddRouter(Id, router)
}

func (s *Server) GetConnMan() rIterface.IConnManger {
	return s.ConnMan
}

func (s *Server) SetStartFunc(start func(conn rIterface.IConnection)) {
	s.OnStartFunc = start
}

func (s *Server) SetStopFunc(stop func(conn rIterface.IConnection)) {
	s.OnStopFunc = stop
}

func (s *Server) CallStartFunc(conn rIterface.IConnection) {
	if s.OnStartFunc != nil {
		s.OnStartFunc(conn)
	} else {
		fmt.Println("OnStartFunc is nil")
	}
}

func (s *Server) CallStopFunc(conn rIterface.IConnection) {
	if s.OnStopFunc != nil {
		s.OnStartFunc(conn)
	} else {
		fmt.Println("OnStopFunc is nil")
	}
}
