package main

import (
	"fmt"
	"hangmanweb/hangman"
	"html/template"
	"net/http"
	"os"
)

type PageInit struct {
	Username string
}

var logs PageInit

func main() {
	hangman.New()
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "intro", nil)
	})

	http.HandleFunc("/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "whatsapp", nil)
	})

	http.HandleFunc("/username", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "username", nil)

	})

	http.HandleFunc("/treatment", func(w http.ResponseWriter, r *http.Request) {
		logs = PageInit{
			r.FormValue("pseudo")}
		fmt.Println(logs)
		http.Redirect(w, r, "/debut", http.StatusMovedPermanently)
		http.Redirect(w, r, "/level1", http.StatusMovedPermanently)
		http.Redirect(w, r, "/level2", http.StatusMovedPermanently)
		http.Redirect(w, r, "/level3", http.StatusMovedPermanently)
	})

	http.HandleFunc("/debut", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "debut", logs)
	})

	http.HandleFunc("/level1", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "easy", logs)
	})

	http.HandleFunc("/level2", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "medium", nil)
	})

	http.HandleFunc("/level3", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "hard", nil)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Init serv
	http.ListenAndServe("localhost:8080", nil)
}
