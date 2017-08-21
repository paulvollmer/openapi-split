package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/paulvollmer/yaml2json/src"
)

const (
	version = "0.1.2"
	repoURL = "https://github.com/paulvollmer/openapi-split"
)

var (
	indexFile      string
	definitionsDir string
	pathsDir       string
	responsesDir   string
	format         string
)

func usage() {
	fmt.Print("Usage: openapi-split [flags]\n\n")
	fmt.Print("Flags:\n")
	flag.PrintDefaults()
	fmt.Println("\nTHIS IS AN ALPHA RELEASE. KEEP IN MIND THERE CAN BE MAJOR CHANGES IN FUTURE")
	fmt.Println("PLEASE REPORT ISSUES AT THE GITHUB REPOSITORY")
	fmt.Println(repoURL + "/issues")
}

func main() {
	flag.StringVar(&indexFile, "i", "index.yaml", "path to the main openapi specification file")
	flag.StringVar(&definitionsDir, "d", "./definitions", "filepath to the definitions yaml files")
	flag.StringVar(&pathsDir, "p", "./paths", "filepath to the paths yaml files")
	flag.StringVar(&responsesDir, "r", "./responses", "filepath to the responses yaml files")
	flag.StringVar(&format, "f", "json", "the file format to print to stdout. can be yaml or json")
	flagVersion := flag.Bool("version", false, "print out the version and exit")
	flag.Usage = usage
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	indexData, err := ioutil.ReadFile(indexFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("please check if the -i flag point to the correct specification yaml file.")
		os.Exit(1)
	}

	responsesData := concatYamlFilesFromDir(responsesDir)
	pathsData := concatYamlFilesFromDir(pathsDir)
	definitionsData := concatYamlFilesFromDir(definitionsDir)

	// print out the specification
	result := make([]string, 0)
	result = append(result, fmt.Sprintf("# Generated by openapi-split v%s\n\n", version))
	result = append(result, string(indexData)+"\n")
	if len(responsesData) != 0 {
		result = append(result, fmt.Sprintf("responses:\n%s\n", strings.Join(responsesData, "\n")))
	}
	if len(pathsData) != 0 {
		result = append(result, fmt.Sprintf("paths:\n%s\n", strings.Join(pathsData, "\n")))
	}
	if len(definitionsData) != 0 {
		result = append(result, fmt.Sprintf("definitions:\n%s\n", strings.Join(definitionsData, "\n")))
	}

	resultStr := strings.Join(result, "\n")

	switch strings.ToLower(format) {
	case "yaml":
		fmt.Println(resultStr)
		break

	case "json":
		yamlParsed, err := yaml2json.BytesToYAMLDoc([]byte(resultStr))
		if err != nil {
			printError(err)
		}
		jsonRaw, err := yaml2json.YAMLToJSON(yamlParsed)
		if err != nil {
			fmt.Println(yamlParsed)
			printError(err)
		}
		json, err := jsonRaw.MarshalJSON()
		if err != nil {
			printError(err)
		}
		fmt.Println(string(json))
		break

	default:
		fmt.Println("format not supported")
		os.Exit(1)
	}
}
