package server

import (
	"log"
	"fmt"
	"strconv"
	"io"
	"os"
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
	w.Write([]byte("An error occurred. Please try again."))
}

func clientError(w http.ResponseWriter, err error) {
	log.Print("[ERR] ",err)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Bad Request. Please check your input."))
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

// 409 - if file exist
// 201 - upload successful
// 400 - missing/incorrect params in request.
// 500 - server error 
func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /add by ", r.RemoteAddr)
	if r.Method != "Post" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

    // left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		clientError(w, err)
		return 
	}
	name := r.Form.Get("name")
	// Retrieve the file from form data
	f, _, err := r.FormFile("file")
	if err != nil {
		clientError(w, err)
		return  
	}
	defer f.Close()

	// check if file exist based on file name
	fileNames, err := fileutil.GetAllFileNames(serverDir)
	if err != nil {
		serverError(w, err)
		return
	}
	for _, fn := range fileNames{
		if fn == name{
			w.WriteHeader(http.StatusConflict)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(name+" already exists."))
			return
		}
	}

	fullPath := serverDir + "/" + name
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		serverError(w, err)
		return  
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		serverError(w, err)
		return 
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(name+" uploaded successfully"))
}

// 200- update successful
// 400 - missing/incorrect params in request.
// 500 - server error 
func updateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] Invoked /add by ", r.RemoteAddr)
	if r.Method != "Put" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

    // left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		clientError(w, err)
		return 
	}
	name := r.Form.Get("name")
	// Retrieve the file from form data
	f, _, err := r.FormFile("file")
	if err != nil {
		clientError(w, err)
		return  
	}
	defer f.Close()

	fullPath := serverDir + "/" + name
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		serverError(w, err)
		return  
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		serverError(w, err)
		return 
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(name+" updated successfully"))
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
	http.HandleFunc("/add", addHandler)	// upload a file
	http.HandleFunc("/update", updateHandler)	// update a file

	// start server
	log.Print("[INFO] Server listening on ", port)
	http.ListenAndServe(":"+port, nil)
}
