package main

import "github.com/runx/server/rImpl"

func main() {
	s := rImpl.NewServer("LCL")
	s.Sever()

}
