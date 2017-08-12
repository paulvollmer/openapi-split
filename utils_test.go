package main

import "testing"

func Test_GetYamlFilenames(t *testing.T) {
	files, err := GetYamlFilenames("./fixtures/definitions")
	if err != nil {
		t.Error(err)
	}
	if len(files) != 3 {
		t.Error("GetYamlFilenames not equal")
	}
}
