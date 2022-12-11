package server

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked / by ", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Welcome to File Store"))
}

func Start(port string){
	// configure routes
	http.HandleFunc("/", rootHandler)

	// start server
	log.Print("[INFO] Server listening on ", port)
	http.ListenAndServe(":"+port, nil)
}
