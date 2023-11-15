package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "intro", nil)
	})

	http.HandleFunc("/debut", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "debut", nil)
	})

	http.HandleFunc("/level1", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level1", nil)
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
