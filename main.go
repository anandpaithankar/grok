package main

import (
	"context"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func start(ctx context.Context) error {
	l, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv())
	if err != nil {
		return err
	}

	log.Println("ngrok ingress : ", l.URL())
	return http.Serve(l, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	contentType := r.Header.Get("Content-Type")
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
