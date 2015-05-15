package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	// r.RequestURI looks like this: /path/to/dest/
	dstFile := filepath.Join(rootDir, r.URL.Path)
	if stat, err := os.Stat(dstFile); err == nil {
		if stat.Mode().IsDir() {
			// First check if index.md exist.
			indexMD := filepath.Join(dstFile, "index.md")
			if stat, err := os.Stat(indexMD); err == nil {
				if stat.Mode().IsRegular() {
					if md, err := ioutil.ReadFile(indexMD); err == nil {
						renderMD(r.URL.Path, md, w)
						return
					}
				}
			}

			// Then render default directory page.
			dirMD := "Directory Contents:\n\n"
			entries, err := ioutil.ReadDir(dstFile)
			if err == nil {
				for _, ent := range entries {
					//TODO: add hyper link
					if ent.Mode().IsDir() {
						dirMD += "* " + ent.Name() + "/\n"
					} else {
						dirMD += "* " + ent.Name() + "\n"
					}
				}
				renderMD(r.URL.Path, []byte(dirMD), w)
			} else {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			}
		} else if stat.Mode().IsRegular() {
			if md, err := ioutil.ReadFile(dstFile); err == nil {
				renderMD(r.URL.Path, md, w)
			} else {
				http.Error(w, "400 Bad Request", http.StatusBadRequest)
			}
		}
	} else {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q != "" {
		if res := searchFolder(q, rootDir); res != nil {
			renderSearchPage(w, res)
		} else {
		}
	} else {
	}
}
