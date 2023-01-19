package web

import (
	"log"
	"net/http"
)

func init() {
	// The index page handler. This is a bit special because it handles the index
	// page ("/") and any pages that don't match a registered route (serves as the
	// catch all handler).
	routes.register(route{
		Path: "/",
		Handler: logger(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				if r.URL.Path != "/" {
					// NOTE: This is the catch all code path. We could do things here like
					// redirect old broken links, render a 404 page, etc.

					http.NotFound(w, r)
				} else {
					data := map[string]interface{}{
						"Tracks":       MusicCollection,
						"DefaultTrack": MusicCollection[0],
					}

					if err := templates.ExecuteTemplate(w, "pages/index.tmpl", data); err != nil {
						log.Printf("web: error rendering index page; %s", err)
					}
				}
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
		}),
	})
}
