package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"

	"gopkg.in/flosch/pongo2.v3"
)

func renderMD(uri string, md []byte, w http.ResponseWriter) {
	extensions := blackfriday.EXTENSION_NO_INTRA_EMPHASIS | blackfriday.EXTENSION_TABLES | blackfriday.EXTENSION_FENCED_CODE | blackfriday.EXTENSION_AUTOLINK | blackfriday.EXTENSION_STRIKETHROUGH | blackfriday.EXTENSION_SPACE_HEADERS
	all := blackfriday.Markdown(md, blackfriday.HtmlRenderer(blackfriday.HTML_USE_SMARTYPANTS|blackfriday.HTML_SMARTYPANTS_FRACTIONS|blackfriday.HTML_TOC, "", ""), extensions)

	t, err := getTemplate("web/md.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dirStructure := getCurrentDirStruct()
	fmt.Printf("%v\n", dirStructure)
	autoComplete := formatToAutoComplete(dirStructure)
	navTree := formatToNavTree(dirStructure)
	fmt.Printf("%v\n", navTree)

	toc, content := splitTOCAndContent(string(all))

	err = t.ExecuteWriter(pongo2.Context{"content": content, "toc": toc, "autoComplete": autoComplete, "navTree": navTree}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func splitTOCAndContent(whole string) (string, string) {
	rdr := strings.NewReader(whole)
	doc, err := goquery.NewDocumentFromReader(rdr)
	if err != nil {
		return "", ""
	}

	body := doc.Find("body")
	nav := body.Find("nav").First()
	toc, _ := nav.Html()

	nav.Remove()
	body = doc.Find("body")
	content, _ := body.Html()
	return toc, content
}

func renderSearchPage(w http.ResponseWriter, p *SearchResult) {
	dirStructure := getCurrentDirStruct()
	fmt.Printf("%v\n", dirStructure)
	autoComplete := formatToAutoComplete(dirStructure)
	navTree := formatToNavTree(dirStructure)
	fmt.Printf("%v\n", navTree)

	var test []byte
	t, err := getTemplate("web/search.tpl")
	if err != nil {
		goto internalError
	}
	test, err = json.Marshal(*p)
	if err == nil {
		fmt.Println(string(test))
	}

	err = t.ExecuteWriter(pongo2.Context{"search": p, "autoComplete": autoComplete, "navTree": navTree}, w)
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
