package main

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"strings"
	"log"
)

func serverCall() (string){

	url := "http://www.broadbandmap.gov/broadbandmap/broadband/jun2014/wireless?latitude=40.7195898&longitude=-73.9998334&format=json"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	type Message struct {
		Status string
		Results interface{}
	}
	
	type InnerMessage struct {
		Frm string
	}
	
	n := string(body[:])
	
	dec := json.NewDecoder(strings.NewReader(n))
	
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if (m.Status == "OK") {
			parseWired(m.Results)
		}
	}
	
	return string(body)

}