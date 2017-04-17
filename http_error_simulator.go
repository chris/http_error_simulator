package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request for", r.URL)

  body := "<html><head><title>502 Bad Gateway</title></head><body bgcolor=\"white\"><center><h1>502 Bad Gateway</h1></center><hr><center>nginx</center></body></html>"

	http.Error(w, body, http.StatusBadGateway)
	return
}

func main() {
	http.HandleFunc("/", handler)

	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	log.Println("Listening on port", port, "...")
	if port == "443" {
		http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	} else {
		http.ListenAndServe(":"+port, nil)
	}
}
