package main

import 	(
			"html/template"
			"log"
			"net/http"
		)

var tpl *template.Template

type pageData struct {
	Title string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	pd := pageData {
		Title: "Index!",

	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
		log.Println("Error logged.", err)
		http.Error(w, "Fehler.", http.StatusInternalServerError)
		return
	}
}