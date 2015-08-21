package main

import (
		//"io/ioutil"
		//"net/http"
		//"html/template"
		//"log"
		//"fmt"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
)

// type Data struct {
//     Name string
// }

// type Page struct {
// 		Title string
// 		Body  []byte
// }

// type loginPage struct {
// 		Title string
// }

// type savePage struct {
// 		Title string
// }

// type sliderPage struct {
// 		id string
// 		Title, Name, Link1, Link2, Link3, Img1, Img2, Img3, SubTitle1, SubTitle2, SubTitle3 string
// 		MainContent []byte
// }

// type Person struct {
// 		Name string
//         Phone string
// }

// func loadPage_Example(title string) (*Page, error){
//         session, err := mgo.Dial("LOCALHOST:27017")
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()

//         // Optional. Switch the session to a monotonic behavior.
//         session.SetMode(mgo.Monotonic, true)

//         c := session.DB("janus").C("users")

//         result := Person{}
//         err = c.Find(bson.M{"name": title}).One(&result)
//         if err != nil {
//                 log.Fatal(err)
//         }
		
//         return &Page{Title: title, Body: []byte(result.Phone)}, nil
// }

// func loadPage_Login(title string) (*loginPage, error){
//         return &loginPage{Title: title}, nil
// }

func mongo_i(session_id string, name string, value string){
        
        session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        
        c := session.DB("intatl").C(session_id)

        c.Insert(bson.M{"key": name, "value": value})
}

// func loadPage_Slider(id string) (*sliderPage, error){

// 	    session, err := mgo.Dial("LOCALHOST:27017")
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()

//         session.SetMode(mgo.Monotonic, true)

//         c := session.DB("test").C("siteContent")

//         result := sliderPage{}
		
// 		err = c.Find(bson.M{"id": id}).One(&result)
		
//         if err != nil {
//                 log.Fatal(err)
//         }
		
//         return &sliderPage{Title: result.Title, Name: result.Name, Link1: result.Link1, Link2: result.Link2, Link3: result.Link3, Img1: result.Img1, Img2: result.Img2, Img3: result.Img3, SubTitle1: result.SubTitle1, SubTitle2: result.SubTitle2, SubTitle3: result.SubTitle3, MainContent: []byte(result.MainContent)}, nil
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/view/"):]
//     p, _ := loadPage_Example(title)
//     t, _ := template.ParseFiles("view.html")
//     t.Execute(w, p)
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/login/"):]
//     name := r.FormValue("Name")
//     fmt.Println("name:", name)
//     p, _ := loadPage_Login(title)
//     t, _ := template.ParseFiles("login.html")
//     t.Execute(w, p)
// }

// func ajaxHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/ajax/"):]
//     p, _ := loadPage_Login(title)
//     t, _ := template.ParseFiles("ajax.html")
//     t.Execute(w, p)
// }

// func sliderHandler(w http.ResponseWriter, r *http.Request) {
//     id := r.URL.Path[len("/slider/"):]
//     p, _ := loadPage_Slider(id)
//     t, _ := template.ParseFiles("slider.html")
//     t.Execute(w, p)
// }

// func saveHandler(w http.ResponseWriter, r *http.Request) {
//     id := r.URL.Path[len("/save/"):]
//     p, _ := loadPage_Save(id, "Name")
//     t, _ := template.ParseFiles("save.html")
//     t.Execute(w, p)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//     // id := r.URL.Path[len("/mongo-i/"):]
//     // p, _ := mongo-i("test", "Name", "Value")
//     // t, _ := template.ParseFiles("index.html")
//     // t.Execute(w, p)
//     mongo_i("Test", "Name", "Value")
// }

// func main() {
// 	http.Handle("/insert/", http.StripPrefix("/insert/", mongo_i("Test", "Umm", "Fuck Yea")))
// // 	http.HandleFunc("/login/", loginHandler)
// // 	http.HandleFunc("/ajax/", ajaxHandler)
// // 	http.HandleFunc("/slider/", sliderHandler)
// // 	http.HandleFunc("/save/", saveHandler)
// // 	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
// // 	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
// // 	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
// // 	http.Handle("/txt/", http.StripPrefix("/txt/", http.FileServer(http.Dir("txt"))))
// 	http.ListenAndServe(":8080", nil)
// }
