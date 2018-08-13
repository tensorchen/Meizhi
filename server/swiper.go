package main

import (
	"io/ioutil"
	"net/http"
)

func ApiSwiper(rw http.ResponseWriter, req *http.Request) {

	rsp, err := http.Get("https://gank.io/api/random/data/福利/4")
	if err != nil {
		rw.Write([]byte("Http Get Fail"))
		return
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		rw.Write([]byte("ioutil ReadAll Fail"))
		return
	}

	rw.Write(data)
}
