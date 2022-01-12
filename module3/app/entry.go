package main

import "github.com/vlpolak/swtgo/module3/server"

func main() {
	s := server.CreateServer()
	s.Serve()
}
