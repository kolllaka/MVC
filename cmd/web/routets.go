package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet", app.showSnippetHandler)
	mux.HandleFunc("/snippet/create", app.createSnippetHandler)
	mux.HandleFunc("/snippet/update", app.updateSnippetHandler)
	mux.HandleFunc("/snippet/delete", app.deleteSnippetHandler)

	return mux
}
