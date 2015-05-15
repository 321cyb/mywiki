package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
	"github.com/russross/blackfriday"

	"gopkg.in/flosch/pongo2.v3"
)

func renderMD(uri string, md []byte, w http.ResponseWriter) {
	extensions := blackfriday.EXTENSION_NO_INTRA_EMPHASIS | blackfriday.EXTENSION_TABLES | blackfriday.EXTENSION_FENCED_CODE | blackfriday.EXTENSION_AUTOLINK | blackfriday.EXTENSION_STRIKETHROUGH | blackfriday.EXTENSION_SPACE_HEADERS
	content := blackfriday.Markdown(md, blackfriday.HtmlRenderer(blackfriday.HTML_USE_XHTML|blackfriday.HTML_USE_SMARTYPANTS|blackfriday.HTML_SMARTYPANTS_FRACTIONS|blackfriday.HTML_TOC, "", ""), extensions)

	t, err := getTemplate("web/md.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteWriter(pongo2.Context{"content": string(content)}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderSearchPage(w http.ResponseWriter, p *SearchResult) {
	var test []byte
	t, err := getTemplate("web/search.tpl")
	if err != nil {
		goto internalError
	}
	test, err = json.Marshal(*p)
	if err == nil {

		fmt.Println(string(test))
	}
	err = t.ExecuteWriter(pongo2.Context{"search": p}, w)
	if err != nil {
		goto internalError
	}

	return

internalError:
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func getTemplate(tpl string) (*pongo2.Template, error) {
	templateBox, err := rice.FindBox(filepath.Dir(tpl))
	if err != nil {
		fmt.Println("cannot find box")
		return nil, err
	}
	// get file contents as string
	templateString, err := templateBox.String(filepath.Base(tpl))
	if err != nil {
		fmt.Println("cannot find file")
		return nil, err
	}

	t, err := pongo2.FromString(templateString)
	if err != nil {
		return nil, err
	}

	return t, nil
}
