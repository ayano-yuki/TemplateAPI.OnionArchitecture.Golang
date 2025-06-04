package main

import "templateapi/onionarchitecture/golang/server"

func main() {
	s := server.NewServer()
	s.Start()
}
