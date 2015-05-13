package main

import (
	"log"
	"net/http"

	"github.com/couchbaselabs/gocb/gocbcore"
	"github.com/daishisystems/payment/router"
)

func main() {

	gocbcore.SetLogger(gocbcore.DefaultStdOutLogger())
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
