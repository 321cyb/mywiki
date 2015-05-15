package main

import "testing"

func TestGetTemplate(t *testing.T) {
	_, err := getTemplate("web/md.tpl")
	if err != nil {
		t.Errorf("getTemplate for web/md.tpl error: %s", err)
	}
}
