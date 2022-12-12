package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
	"github.com/Siddhesh-Ghadi/file-store/pkg/model"
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
)

var serverAddr = "http://localhost:8080"

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func Help() {
	fmt.Println("Help")
}

func Ls() {
	resp, err := http.Get(serverAddr + "/ls")
	if err != nil {
		handleError(err)
	}
	//We Read the response body on the line below.
   	body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
		handleError(err)
   	}
	//Convert the body to type string
	var b model.LsResponse
	e := json.Unmarshal(body, &b)
	if e != nil {
		handleError(err)
	}
	for _, v := range b.Files{
		fmt.Println(v.Name)
	}
}

func Add(path string) {
	_, name := filepath.Split(path)
	extraParams := map[string]string{
		"name": name,
	}
	request, err := newfileUploadRequest("POST", serverAddr + "/add", extraParams, "file", path)
	if err != nil {
		handleError(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		handleError(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
    if err != nil {
		handleError(err)
	}
    resp.Body.Close()
		fmt.Println(body)
	}
}

func Update(path string) {
	_, name := filepath.Split(path)
	extraParams := map[string]string{
		"name": name,
	}
	request, err := newfileUploadRequest("PUT", serverAddr + "/update", extraParams, "file", path)
	if err != nil {
		handleError(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		handleError(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			handleError(err)
		}
		resp.Body.Close()
		fmt.Println(body)
	}
}

func Rm(path string){
	_, name := filepath.Split(path)
	extraParams := map[string]string{
		"name": name,
	}
	request, err := newfileUploadRequest("DELETE", serverAddr + "/rm", extraParams, "file", "")

	if err != nil {
		handleError(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		handleError(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			handleError(err)
		}
		resp.Body.Close()
		fmt.Println(body)
	}
}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-file-upload-example/
func newfileUploadRequest(method string,uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}