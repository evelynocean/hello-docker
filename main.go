package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

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
	// NSQMessages := make(chan string, 200)
	// config := nsq.NewConfig()

	// q, er := nsq.NewConsumer("hello", "test", config)
	// if er != nil {
	// 	log.Panic("-- ----------------- NewConsumer :", er)
	// }

	// q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

	// 	message.DisableAutoResponse()
	// 	defer message.Finish()
	// 	fmt.Println("string(message.Body):", string(message.Body))
	// 	NSQMessages <- string(message.Body)
	// 	return nil
	// }))

	// if err := q.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
	// 	log.Panic("------------------- ConnectToNSQLookupd: ", err)
	// }
	r := mux.NewRouter()
	sub := r.PathPrefix("/api").Subrouter()
	sub.HandleFunc("/hello", index()).Methods("GET")

	n := negroni.New()
	n.Use(httplog.NewLogger())
	n.UseHandler(r)
	fmt.Println("service start:")
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

//NSQMessages chan string
func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// select {
		// case s := <-NSQMessages:
		str := time.Now().Format("2006-01-02 15:04:05")
		restresp.Write(w, str, http.StatusOK)
		// }
		return
	}

}
