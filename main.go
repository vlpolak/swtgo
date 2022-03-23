package main

import "github.com/vlpolak/swtgo/server"

func main() {
	//ws.StartServer()
	s := server.CreateServer()
	s.Serve()
}
