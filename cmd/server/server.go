package main

import (
	"os/user"
	"log"
	"github.com/Siddhesh-Ghadi/file-store/pkg/server"
)

func main() {
	usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
	server.Start("8080", usr.HomeDir+"/.file-store")
}