package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type SearchResult struct {
	SearchTerm string
	Results    []SearchResultPerFile
}

type SearchResultPerFile struct {
	Appearances [][]string
	FileName    string
}

const CONTEXTS = 3

func grepFile(path string, pattern *regexp.Regexp) *SearchResultPerFile {
	result := &SearchResultPerFile{
		Appearances: make([][]string, 0),
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Errorf("Open file %s error: %s", path, err)
		return result
	}

	lines := make(map[int]string)
	lineno := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines[lineno] = scanner.Text()
		lineno++
	}

	if err := scanner.Err(); err != nil {
		log.Errorf("reading file %s error: %s", path, err)
	}

	linesToShow := make([]int, 0)
	for i := 0; i < lineno; i++ {
		if pattern.FindString(lines[i]) != "" {
			linesToShow = addToSlice(i, linesToShow)
			for j := 1; j <= CONTEXTS && (i-j) >= 0; j++ {
				linesToShow = addToSlice(i-j, linesToShow)
			}
			for j := 1; j <= CONTEXTS && (i+j) < lineno; j++ {
				linesToShow = addToSlice(i+j, linesToShow)
			}
		}
	}

	sort.Ints(linesToShow)

	last := -1
	for i, number := range linesToShow {
		if i == 0 || (last+1) != number {
			result.Appearances = append(result.Appearances, []string{lines[number]})
		} else {
			numAppearences := len(result.Appearances)
			result.Appearances[numAppearences-1] = append(result.Appearances[numAppearences-1], lines[number])
		}

		last = number
	}

	if len(result.Appearances) > 0 {
		return result
	} else {
		return nil
	}
}

func addToSlice(v int, s []int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return s
		}
	}
	s = append(s, v)
	return s
}

func searchFolder(searchTerm string, root string) *SearchResult {
	result := SearchResult{
		SearchTerm: searchTerm,
		Results:    make([]SearchResultPerFile, 0),
	}

	exp, err := regexp.Compile(searchTerm)
	if err != nil {
		log.Errorf("searchRootFolder, regexp compile error: %s", err)
		return nil
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && (strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".txt")) {
			res := grepFile(path, exp)
			if res != nil {
				if relName, err := filepath.Rel(root, path); err == nil {
					res.FileName = relName
				}
				result.Results = append(result.Results, *res)
			}
		}
		return nil
	})

	if len(result.Results) > 0 {
		return &result
	} else {
		return nil
	}
}
