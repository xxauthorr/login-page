package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	tpl *template.Template

	key   = []byte("First-session-integration")
	store = sessions.NewCookieStore(key)
)

func init() {
	var err error
	// Parsing every .html files in the folder ../views/
	tpl, err = tpl.ParseGlob("./static/*.html")

	if err != nil {
		panic(err)
	}
}
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	id := session.Values["emailId"].(string)
	article := Articles{
		Article{
			Email:   id,
			Name:    "Hari",
			Details: "onullka"},
	}
	json.NewEncoder(w).Encode(article)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	FileServer := http.FileServer(http.Dir("./static/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", FileServer))

	//Route Handle functions for each routes
	router.HandleFunc("/", LoginHandler)

	router.HandleFunc("/loginCheck", LoginCheckHandler)
	router.HandleFunc("/home", auth(HomeHandler)).Methods("GET")
	router.HandleFunc("/home", ArticleHandler).Methods("POST")
	router.HandleFunc("/logOut", LogOutHandler)
	router.NotFoundHandler = http.HandlerFunc(NoPageHandler)

	fmt.Println("Server running succesfully ")
	//server created at the port 3000
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		//prints the error if the port is unable to create
		log.Fatal(err)
	}
}
