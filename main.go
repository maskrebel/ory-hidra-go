package main

import (
	"fmt"
	"net/http"
	"html/template"
	"path"
)
var data = map[string]interface{}{
	"title": "SDK",
	"name":  "Batman",
}




func index_handler(w http.ResponseWriter, r *http.Request)  {
	var filepath = path.Join("templates", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login_handler(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "http://www.google.com", 301)
}

func main()  {
	fmt.Println("Server Run In http://localhost:3000")
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/login", login_handler)
	http.ListenAndServe(":3000", nil)
}