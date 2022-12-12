package main

import (
  "os"
  "github.com/Siddhesh-Ghadi/file-store/pkg/client"
)



func main() {

	if len(os.Args) < 2{
		client.Help()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "ls":
		client.Ls()
	case "add":
		files := os.Args[2:]
		for _, f := range files{
			client.Add(f)
		}
	case "update":
		files := os.Args[2:]
		for _, f := range files{
			client.Update(f)
		}	
	case "rm":
		files := os.Args[2:]
		for _, f := range files{
			client.Rm(f)
		}	
	case "freq-words":
		limit := "10" 
		order := "dsc"
		args := os.Args[2:]
		for i, a  := range args{
			if a == "--limit" || a == "-n"{
				limit = args[i+1]
			} else if a == "--order" {
				order = args[i+1]
			}
		}
		client.FreqWords(limit, order)
	default:
		client.Help()
	}

}