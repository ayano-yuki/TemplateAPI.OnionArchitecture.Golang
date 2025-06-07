package main

import "API/server"

// @title           TemplateAPI.OnionArchitecture.Golang
// @version         1.0
// @description     This is API Template for Onion Architecture.
// @host            localhost:8080
// @BasePath        /

func main() {
	s := server.NewServer()
	s.Start()
}
