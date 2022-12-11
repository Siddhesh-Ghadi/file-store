package server

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/Siddhesh-Ghadi/file-store/pkg/fileutil"
	"github.com/Siddhesh-Ghadi/file-store/pkg/model"
)

var serverDir string

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked / by ", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Welcome to File Store"))
}

func lsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /ls by ", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	
	fileNames, err := fileutil.GetAllFileNames(serverDir)
	if err != nil {
		log.Print("[ERR] ",err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		return
	}

	resp := model.LsResponse{}
	for _, fn := range fileNames{
		resp.Files = append(resp.Files, model.File{Name: fn})
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[ERR] Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}


func Start(port string, dir string){
	serverDir = dir

	// create data directory, ignore if exists
	fileutil.CreateDir(serverDir)

	// configure routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ls", lsHandler)

	// start server
	log.Print("[INFO] Server listening on ", port)
	http.ListenAndServe(":"+port, nil)
}
