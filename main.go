package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OK\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func PostMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/healthcheck", HealthCheck)
	router.GET("/hello/:name", Hello)
	router.POST("/post_message", PostMessage)

	log.Fatal(http.ListenAndServe(":8080", router))
}

