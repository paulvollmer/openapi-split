package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GetYamlFilenames(dirname string) ([]string, error) {
	files := make([]string, 0)
	d, err := ioutil.ReadDir(dirname)
	if err != nil {
		return files, err
	}
	for _, v := range d {
		if !v.IsDir() {
			filename := v.Name()
			fileext := strings.ToLower(path.Ext(filename))
			if fileext == ".yml" || fileext == ".yaml" {
				files = append(files, path.Join(dirname, filename))
			}
		}
	}
	return files, nil
}

func concatYamlFiles(files []string) ([]string, error) {
	d := make([]string, 0)
	for _, v := range files {
		content, err := readLines(v, "  ")
		if err != nil {
			return d, err
		}
		d = append(d, content...)
	}
	return d, nil
}

func concatYamlFilesFromDir(dirname string) []string {
	data := make([]string, 0)
	files, err := GetYamlFilenames(dirname)
	if err == nil {
		data, err = concatYamlFiles(files)
		if err != nil {
			printError(err)
		}
	}
	return data
}

func printError(err error) {
	fmt.Println("Error", err)
	os.Exit(1)
}

func readLines(path, leftPad string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, leftPad+scanner.Text())
	}
	return lines, scanner.Err()
}
