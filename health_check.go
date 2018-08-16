package main

import "net/http"

func HealthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
}
