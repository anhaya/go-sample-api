package server

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func InitServer(router http.Handler) {
	http.Handle("/", router)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8081),
	}

	log.Printf("server started: %s", strconv.Itoa(8081))
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
