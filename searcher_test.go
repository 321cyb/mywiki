package main

import (
	"regexp"
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestSearcherBasic(t *testing.T) {
	exp, err := regexp.Compile(`BBB`)
	if err != nil {
		log.Fatalf("regexp Compile error: %s", err)
	}

	result := grepFile("testdata/match-basic.txt", exp)
	if result == nil {
		t.Error("grepFile should return a valid result, not nil.")
	}
}

func TestSearcherHead(t *testing.T) {
	exp, err := regexp.Compile(`BBB`)
	if err != nil {
		log.Fatalf("regexp Compile error: %s", err)
	}

	result := grepFile("testdata/match-head.txt", exp)
	if result == nil {
		t.Error("grepFile should return a valid result, not nil.")
	}
}

func TestSearcherEnd(t *testing.T) {
	exp, err := regexp.Compile(`BBB`)
	if err != nil {
		log.Fatalf("regexp Compile error: %s", err)
	}

	result := grepFile("testdata/match-end.txt", exp)
	if result == nil {
		t.Error("grepFile should return a valid result, not nil.")
	}
}

func TestSearcherNone(t *testing.T) {
	exp, err := regexp.Compile(`BBB`)
	if err != nil {
		log.Fatalf("regexp Compile error: %s", err)
	}

	result := grepFile("testdata/match-none.txt", exp)
	if result != nil {
		t.Error("grepFile should return nil.")
	}
}

func TestSearcherSeries(t *testing.T) {
	exp, err := regexp.Compile(`BBB`)
	if err != nil {
		log.Fatalf("regexp Compile error: %s", err)
	}

	result := grepFile("testdata/match-series.txt", exp)
	if result == nil {
		t.Error("grepFile should return nil.")
	}

	if len(result.Appearances) != 1 {
		t.Errorf("match-series.txt should have only one apperance while it has %d.", len(result.Appearances))
	}

	if len(result.Appearances[0]) != (2 + 2*CONTEXTS) {
		t.Errorf("match-series.txt should have %d lines to show.", 2+2*CONTEXTS)
	}

}

func TestSearcherFolder(t *testing.T) {
	res := searchFolder("BBB", "testdata")
	if res == nil {
		log.Fatal("search folder testdata failed: %s")
	}

	if len(res.Results) != 4 {
		t.Errorf("under testdata, there should be 4 matches, but it returns %d", len(res.Results))
	}
}
