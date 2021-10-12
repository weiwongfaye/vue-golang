package main

import (
	"backend/endpoints"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type corsRouterDecorator struct {
	R *mux.Router
}

func (c *corsRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requeted-With")
	}
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}

func main() {
	r := mux.NewRouter()
	r = endpoints.AddRouterEndpoints(r)

	fs := http.FileServer(http.Dir("../frontend/dist"))
	r.PathPrefix("/").Handler(fs)

	http.Handle("/", &corsRouterDecorator{r})
	fmt.Println("Listening")
	log.Panic(http.ListenAndServe(":3000", nil))
}
