package main

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"net/smtp"
	"fmt"
)


func sendMail(to []string, subject string, msg string) error {
	var (
		username string = "jon@rebirtharmitage.com"
		password string = "starLight7"
		host     string = "smtp.gmail.com"
		port     string = "587"
	)
	
	auth := smtp.PlainAuth(
		"",
		username,
		password,
		host,
	)

	address := fmt.Sprintf("%v:%v", host, port)

	//	build our message
	body := []byte("Subject: " + subject + "\r\n\r\n" + msg)

	err := smtp.SendMail(
		address,
		auth,
		username,
		to,
		body,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
TYPE : Page
struct for use with HTTP/TEMPLATE to display web pages.  Webpages internal data is stored here.
*/
type Page struct {
	Title string
	Body  string
}

func loadPage(title string) (*Page, error){
	return &Page{Title: title, Body: "blank..."}, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("Test")
  renderTemplate(w, "index", p)
}

func indexNew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	p, _ := loadPage("Test")
  renderTemplate(w, "index_old", p)
}


func email(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	if (name == ""){ name = "NO NAME"}
	email := request.FormValue("email")
	if (email == ""){ email = "NO EMAIL"}
	phone := request.FormValue("phone")
	if (phone == ""){ phone = "NO PHONE"}
	body := request.FormValue("body")
	if (body == ""){ body = "NO MESSAGE"}
	a := sendMail([]string{"jon@rebirtharmitage.com"}, email + " : " + phone, body)
	if (a == nil){
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}else{
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

/*
	Core function for delivering webpages to the clients
*/
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func RedirectIndex(response http.ResponseWriter, request *http.Request){
	redirectTarget := "http://www.google.com"
	http.Redirect(response, request, redirectTarget, 302)
}


/*
	Start of ROUTER Section
*/
var router = mux.NewRouter()
var s = router.Host("www.rebirtharmitage.com").Subrouter()

func main() {
	s.HandleFunc("/", index)
	s.HandleFunc("/new", indexNew)
	s.HandleFunc("/sendmail", email).Methods("POST")
	
	//This Handles all the static file calls such as css/file.css etc.
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("../rebirtharmitage/")))
	
	http.Handle("/", s)
	http.ListenAndServe(":8080", nil)
}
