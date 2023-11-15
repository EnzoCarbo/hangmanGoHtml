package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type PageInit struct {
	Username string
}

var logs PageInit

func main() {

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "intro", nil)
	})

	http.HandleFunc("/username", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "username", nil)

	})

	http.HandleFunc("/treatment", func(w http.ResponseWriter, r *http.Request) {
		logs = PageInit{
			r.FormValue("pseudo")}
		fmt.Println(logs)
		http.Redirect(w, r, "/debut", http.StatusMovedPermanently)
	})

	http.HandleFunc("/debut", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "debut", logs)
	})

	http.HandleFunc("/level1", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "facile", logs)
	})

	http.HandleFunc("/level2", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level2", nil)
	})

	http.HandleFunc("/level3", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level3", nil)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Init serv
	http.ListenAndServe("localhost:8080", nil)
}
