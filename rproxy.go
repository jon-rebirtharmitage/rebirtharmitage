package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "strings"
)

func HandleIt(w http.ResponseWriter, r *http.Request) {
  if strings.Contains(r.Host, "wiki.rebirtharmitage.com"){
    http.Redirect(w, r, "https://wiki.rebirtharmitage.com:8085", 302)
  }else if strings.Contains(r.Host, "internetatlas.co"){
    http.Redirect(w, r, "http://www.internetatlas.co:8081", 302)
  }else if strings.Contains(r.Host, "tattoos"){
    http.Redirect(w, r, "http://www.rebirtharmitage.com:8084", 302)
  }else if strings.Contains(r.Host, "rebirtharmitage.com"){
    http.Redirect(w, r, "http://www.rebirtharmitage.com:8080", 302)
  }
}

var router = mux.NewRouter()

func main(){
  http.HandleFunc("/", HandleIt)
	http.ListenAndServe(":80", nil)
}