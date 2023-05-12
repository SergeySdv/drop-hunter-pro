package ui

import (
	"crypto_scripts/internal/server/config"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"net/url"
)

//go:embed build/*
var ui embed.FS

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// кешируем запросы на неизменные файлы с начала запуска приложения
	// todo: заменить на хеш дирректории со статическим контентом
	key := config.CFG.InstanceId + r.URL.Path
	w.Header().Set("Etag", key)

	if r.Header.Get("If-None-Match") == key {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	sub, _ := fs.Sub(ui, "build")
	fs := http.FS(sub)

	http.FileServer(fs).ServeHTTP(w, r)
}

func ListenStatic(staticServerHost, pathToStaticFiles string) error {
	staticServerHost = "http://localhost" + staticServerHost
	fs := spaHandler{staticPath: pathToStaticFiles, indexPath: "index.html"}
	http.Handle("/", fs)

	serverUrl, err := url.Parse(staticServerHost)
	if err != nil {
		return err
	}

	log.Println("Listening static server on :", serverUrl.Port())

	err = http.ListenAndServe(":"+serverUrl.Port(), fs)
	if err != nil {
		return err
	}
	return nil
}
