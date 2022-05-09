package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	// will return a 404 not found error if the page does not exactly match 
	// the "/" pattern
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Use the template.ParseFiles to load the template into a template set
	// if there is an error, we log the error to the console and return
	// an http Internal Server Error to the response writer.

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// you can pass the files as a variadic parameter to the template.ParseFiles function.
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// w.Write([]byte("Hello from Snippetbox"))
}



// ================================= showSnippetHandler Start ==========================================

func showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

// ================================= showSnippetHandler End ==========================================




func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
