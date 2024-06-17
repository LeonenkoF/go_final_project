package main

import (
	"fmt"
	"main/config"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

var webDir = "D:/Dev/go_final_project/web"

func main() {
	cfg := config.New()

	r := chi.NewRouter()
	FileServer(r, "/", http.Dir(webDir))

	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})

}
