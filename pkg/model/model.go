package model

type File struct {
	Name string `json:"name"`
}

type LsResponse struct {
	Files []File `json:"files"`
}

type Freq struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type FreqResponse struct {
	Freqs []Freq `json:"freqs"`
}