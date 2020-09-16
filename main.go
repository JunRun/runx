package main

import (
	"runx/router"
	"runx/server/rImpl"
	_ "runx/util"
)

func main() {
	s := rImpl.NewServer()
	s.AddRouter(0, &router.PingRouter{})
	s.AddRouter(1, &router.Hell{})
	s.Sever()
}
