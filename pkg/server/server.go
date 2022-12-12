package server

import (
	"log"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/Siddhesh-Ghadi/file-store/pkg/fileutil"
	"github.com/Siddhesh-Ghadi/file-store/pkg/model"
)

var serverDir string

func serverError(w http.ResponseWriter, err error) {
	log.Print("[ERR] ",err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Something went wrong."))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked / by ", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome to File Store"))
}

func lsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /ls by ", r.RemoteAddr)
	
	fileNames, err := fileutil.GetAllFileNames(serverDir)
	if err != nil {
		serverError(w, err)
		return
	}

	resp := model.LsResponse{}
	for _, fn := range fileNames{
		resp.Files = append(resp.Files, model.File{Name: fn})
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func wcHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /wc by ", r.RemoteAddr)
	
	fileNames, err := fileutil.GetAllFileNames(serverDir)
	if err != nil {
		serverError(w, err)
		return
	}

	count := 0
	for _, file := range fileNames {
		c, err := fileutil.GetWordCount(serverDir+"/"+file)
		if err != nil {
			serverError(w, err)
			return
		}
		count = count + c
	}

	resp := make(map[string]string)
	resp["count"] = fmt.Sprint(count)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func freqWordsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /freq-words by ", r.RemoteAddr)
	
	// parse query params
	limit := 10
	order := "asc"
	for k, v := range r.URL.Query() {
        if k == "limit" && len(v) != 0  {
			i, err := strconv.Atoi(v[0])
			if err == nil {
				limit = i
			}
		} else {
			if k == "order" && len(v) != 0 {
				if v[0] == "asc" || v[0] == "dsc" {
					order = v[0]
				}
			}
		}
    }

	fileNames, err := fileutil.GetAllFileNames(serverDir)
	if err != nil {
		serverError(w, err)
		return
	}

	words := []string{}
	for _, file := range fileNames {
		ws, err := fileutil.GetWords(serverDir+"/"+file)
		if err != nil {
			serverError(w, err)
			return
		}
		words = append(words, ws...)
	}

	freq :=  fileutil.GetWordFreq(words, limit, order)

	resp := model.FreqResponse{} 
	for word, count := range freq {
		resp.Freqs = append(resp.Freqs, model.Freq{Word: word, Count: count})
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func Start(port string, dir string){
	serverDir = dir

	// create data directory, ignore if exists
	fileutil.CreateDir(serverDir)

	// configure routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ls", lsHandler)	// list all files
	http.HandleFunc("/wc", wcHandler)	// return number of words from all files
	http.HandleFunc("/freq-words", freqWordsHandler)	// return number of occurrences of words from all files

	// start server
	log.Print("[INFO] Server listening on ", port)
	http.ListenAndServe(":"+port, nil)
}
