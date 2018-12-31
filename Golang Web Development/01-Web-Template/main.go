package main

import (
	"html/template"
	// "io"
	"net/http"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
	// http.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000",nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Index Page.")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
// func about(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "This is About Page.")
// }
// func contact(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "This is Contact Page.")
// }