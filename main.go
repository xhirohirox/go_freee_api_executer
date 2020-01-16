package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/auth_freee", authFreee)
	http.HandleFunc("/auth/freee/callback", authFreeeCallback)
	http.HandleFunc("/register_time_sheet", registerTimeSheet)
	http.ListenAndServe(":8080", nil)
}

func authFreee(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("authFreee.tpl"))
	url := map[string]string{
		"ClientId":    "",
		"RedirectUri": "",
	}
	tpl.Execute(w, url)
}

func authFreeeCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return
	}
	tpl := template.Must(template.ParseFiles("authFreee.tpl"))
	url := map[string]string{
		"ClientId":    "",
		"RedirectUri": "",
	}
	tpl.Execute(w, url)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from Docker container!")
}

func registerTimeSheet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Request!")
}
