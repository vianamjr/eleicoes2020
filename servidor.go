package main

import (
	"html/template"
	"net/http"
)
var temp = template.Must(template.New("").ParseGlob("./templates/*.html"))
func StartService() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":80", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	dados := pesquisas()
	temp.ExecuteTemplate(w, "Index", dados)
}

