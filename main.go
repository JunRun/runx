package main

import (
	"github.com/runx/router"
	"github.com/runx/server/rImpl"
	_ "github.com/runx/util"
)

func main() {
	s := rImpl.NewServer()
	s.AddRouter(&router.PingRouter{})
	s.Sever()
}
