package main

import (
	"log"
	"net/http"

	"github.com/couchbaselabs/gocb"
	"github.com/couchbaselabs/gocb/gocbcore"
	"github.com/daishisystems/payment/router"
)

func main() {

	gocbcore.SetLogger(gocbcore.DefaultStdOutLogger())

	cluster, _ := gocb.Connect("couchbase://52.17.36.153")
	bucket, _ := cluster.OpenBucket("ryanair", "")

	router := router.NewRouter(bucket)
	log.Fatal(http.ListenAndServe(":8080", router))
}
