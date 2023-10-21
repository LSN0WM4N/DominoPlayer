package main

import (
	"app/dominoplayer"
	"net/http"
)

func main() {
	http.HandleFunc("/start", dominoplayer.Start)
	http.HandleFunc("/reset", dominoplayer.Reset)
	http.HandleFunc("/step", dominoplayer.Step)

	http.ListenAndServe("localhost:8000", nil)
}
