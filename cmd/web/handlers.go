package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	// will return a 404 not found error if the page does not exactly match
	// the "/" pattern
	if r.URL.Path != "/" {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)

	}

	// w.Write([]byte("Hello from Snippetbox"))
}

// ================================= showSnippetHandler Start ==========================================

func (app *application) showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

// ================================= showSnippetHandler End ==========================================

func (app *application) createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
