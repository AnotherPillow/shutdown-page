package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "<a href=\"/shutdown\" target=\"_blank\">shutdown</a>")
}
func getShutdown(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Attempting to shut down...\n")
	c := exec.Command("shutdown")
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/shutdown", getShutdown)

	http.ListenAndServe(":80", nil)
}
