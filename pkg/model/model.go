package model

type File struct {
	Name string `json:"name"`
}

type LsResponse struct {
	Files []File `json:"files"`
}