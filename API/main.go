package main

import "API/server"

func main() {
	s := server.NewServer()
	s.Start()
}
