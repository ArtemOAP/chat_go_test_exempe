package main

import (
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/test" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {

	hub := newHub()
	go hub.run()
	http.HandleFunc("/test", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(":8081", nil)
	//err := http.ListenAndServeTLS("host:8081", "./fullchain1.pem", "./privkey1.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
