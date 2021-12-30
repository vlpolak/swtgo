package main

import (
	"github.com/vlpolak/swtgo/module3/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.MakeHandler(handler.RootHandler))
	http.HandleFunc("/materials/", handler.MakeHandler(handler.MaterialsHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
