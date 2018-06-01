package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/syhlion/httplog"
	"github.com/syhlion/restresp"
	"github.com/urfave/negroni"
)

func main() {

	apiListener, err := net.Listen("tcp", ":1031")
	if err != nil {
		log.Fatal(err)
	}

	publicAPIError := make(chan error)

	r := mux.NewRouter()
	sub := r.PathPrefix("/api").Subrouter()
	sub.HandleFunc("/hello", index()).Methods("GET")

	n := negroni.New()
	n.Use(httplog.NewLogger())
	n.UseHandler(r)

	go func() {
		publicAPIError <- http.Serve(apiListener, n)
	}()

	shutdownObserver := make(chan os.Signal, 1)

	select {
	case s := <-shutdownObserver:
		fmt.Println("Receive signal:", s)
		return
	case err := <-publicAPIError:
		fmt.Println("public_api:", err)
		return
	}
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Hello World")
		restresp.Write(w, "Hello World", http.StatusOK)
		return
	}

}
