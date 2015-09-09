package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var inputChan chan bool = make(chan bool)
var outputChan chan map[string]anything = make(chan map[string]anything)

type anything interface{}

// return value represents directory hierarchy.
// each level of directory is a map of string to anything, if anything is string, then it is a regular file,
// if anything is a map of string to anything, then it is a sub folder.
func scanDir(rootdir string) map[string]anything {
	dirStructure := make(map[string]anything)
	var walkFunc = func(path string, info os.FileInfo, _ error) error {
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		fname, err := filepath.Rel(rootdir, path)
		if err != nil {
			log.Printf("scanDir err is %s", err)
			return nil
		}

		segments := strings.Split(fname, "/")
		segmentsSize := len(segments)
		currentLevel := dirStructure
		for index, seg := range segments {
			//skip dot started folders
			if strings.HasPrefix(seg, ".") {
				return nil
			}

			if index != (segmentsSize - 1) {
				currentLevel = currentLevel[seg].(map[string]anything)
			}
		}

		if info.Mode().IsDir() {
			currentLevel[info.Name()] = make(map[string]anything)
		} else if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".md") {
			currentLevel[info.Name()] = ""
		}

		return nil
	}

	filepath.Walk(rootdir, walkFunc)
	return dirStructure
}

// called from main()
func startMonDir(rootdir string) {
	go func() {
		dirStructure := scanDir(rootdir)
		ticker := time.NewTicker(1 * time.Minute)

		for {
			select {
			case <-ticker.C:
				dirStructure = scanDir(rootdir)
			case <-inputChan:
				outputChan <- dirStructure
			}
		}
	}()
}

func getCurrentDirStruct() map[string]anything {
	inputChan <- true
	return <-outputChan
}

// this is used for https://leaverou.github.io/awesomplete/

func formatToAutoComplete(structure map[string]anything) string {
	result := formatToAutoCompleteMap(structure)
	arr := make([]string, len(result))
	index := 0
	for key, _ := range result {
		arr[index] = key
		index++
	}

	if jsStr, err := json.Marshal(arr); err == nil {
		return string(jsStr)
	} else {
		return "[]"
	}
}

func formatToAutoCompleteMap(structure map[string]anything) map[string]bool {
	result := make(map[string]bool)
	for key, value := range structure {
		result[key] = true
		switch value.(type) {
		case string:
			//do nothing if this is just a regular file.
		default:
			subResult := formatToAutoCompleteMap(value.(map[string]anything))
			for subKey, _ := range subResult {
				result[subKey] = true
			}
		}
	}
	return result
}

// this is used for http://responsivemultimenu.com/
func formatToNavTree(structure map[string]anything) string {
	return formatToNavTreeRecursion(structure, "/")
}

// be careful with HTML and URL encoding.
func formatToNavTreeRecursion(structure map[string]anything, parentPath string) string {
	if len(structure) == 0 {
		return ""
	}
	resultHTML := "<ul>\n"

	for key, value := range structure {
		switch value.(type) {
		case string:
			href := url.URL{
				Path: fmt.Sprintf("%s%s", parentPath, key),
			}
			resultHTML += fmt.Sprintf(`<li><a href="%s">%s</a></li>`, href.String(), html.EscapeString(key))
			resultHTML += "\n"
		default:
			href := url.URL{
				Path: fmt.Sprintf("%s%s/", parentPath, key),
			}
			resultHTML += fmt.Sprintf(`<li><a href="%s">%s</a>`, href.String(), html.EscapeString(key))
			resultHTML += "\n"

			resultHTML += formatToNavTreeRecursion(value.(map[string]anything), parentPath+key+"/")
			resultHTML += "</li>\n"
		}
	}
	resultHTML += "</ul>\n"
	return resultHTML
}
