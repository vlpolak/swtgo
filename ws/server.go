package ws

import (
	"flag"
	"github.com/vlpolak/swtgo/logger"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func StartServer() {
	flag.Parse()
	chat := newChat()
	go chat.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(chat, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logger.ErrorLogger("Failed listen and serve processing", err).Log()
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "C:\\Users\\Vladimir_Polyakov\\go\\src\\github.com\\vlpolak\\swtgo\\ws\\home.html")
}
