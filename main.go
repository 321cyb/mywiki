package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
)

var port int
var rootDir string

func init() {
	flag.IntVar(&port, "port", 1025, "port for wiki server")
	flag.StringVar(&rootDir, "root", ".", "root directory of wiki server")

	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	flag.Parse()
	if absRootDir, err := filepath.Abs(rootDir); err != nil {
		fmt.Fprintln(os.Stderr, "root directory error!")
		os.Exit(1)
	} else {
		rootDir = absRootDir
	}

	startMonDir(rootDir)

	box := rice.MustFindBox("web")
	staticFileServer := http.StripPrefix("/_static/", http.FileServer(box.HTTPBox()))
	http.Handle("/_static/", staticFileServer)
	http.HandleFunc("/_search", handleSearch)
	http.HandleFunc("/", handleReq)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
