package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ApiToday(rw http.ResponseWriter, req *http.Request) {

	log.Println("Meizhi ApiToday Enter")

	rsp, err := http.Get("http://gank.io/api/today")
	if err != nil {
		rw.Write([]byte("Http Get Fail"))
		log.Println("Http Get Fail ", err)
		return
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		rw.Write([]byte("ioutil ReadAll Fail"))
		log.Println("ioutil ReadAll Fail", err)
		return
	}

	today := &GankApiToday{}
	err = json.Unmarshal(data, today)
	if err != nil {
		rw.Write([]byte("Json Unmarshal Fail"))
		log.Println("Json Unmarshal Fail", err)
		return
	}

	if len(today.Results.Fuli) == 0 {
		rw.Write([]byte("Fuli NULL"))
		log.Println("Fuli NULL", err)
		return
	}

	r := &Result{
		Code:    0,
		Message: "Ok",
		Data: &DataEntity{
			Image: today.Results.Fuli[0].Url,
		},
	}

	rd, _ := json.Marshal(r)

	rw.Write([]byte(rd))
}

func Today(rw http.ResponseWriter, req *http.Request) {

	log.Println("Meizhi ApiToday Enter")

	rsp, err := http.Get("http://gank.io/api/today")
	if err != nil {
		rw.Write([]byte("Http Get Fail"))
		log.Println("Http Get Fail ", err)
		return
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		rw.Write([]byte("ioutil ReadAll Fail"))
		log.Println("ioutil ReadAll Fail", err)
		return
	}

	today := &GankApiToday{}
	err = json.Unmarshal(data, today)
	if err != nil {
		rw.Write([]byte("Json Unmarshal Fail"))
		log.Println("Json Unmarshal Fail", err)
		return
	}

	if len(today.Results.Fuli) == 0 {
		rw.Write([]byte("Fuli NULL"))
		log.Println("Fuli NULL", err)
		return
	}

	rw.Header().Set("Content-Type", "text/html; charset=UTF-8")
	rd := `<img src="` + today.Results.Fuli[0].Url + `"/>`
	rw.Write([]byte(rd))
}
