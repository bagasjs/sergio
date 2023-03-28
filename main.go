package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"encoding/json"
)

func ReadXMLContent(text string) string {
	var pattern = regexp.MustCompile("<.*?>")
	var content = string(pattern.ReplaceAll([]byte(text), []byte(" ")))
	return content
}

func ReadEntireXMLInDir(dirPath string) map[string]string {
	files := Unwrap(os.ReadDir(dirPath)).([]os.DirEntry)
	var contents = map[string]string {}
	for _, file := range files {
		if !file.IsDir() {
			var filePath = dirPath + "/" + file.Name()
			var content = ReadXMLContent(string(Unwrap(os.ReadFile(filePath)).([]byte)))
			contents[file.Name()] = content
		}
	}
	return contents
}

type TermFreq map[string]int
type TermFreqIndex map[string]TermFreq

func GetSortedTFKeys(tf TermFreq) []string {
	var keys = make([]string, 0, len(tf))
	for key := range tf {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i int, j int) bool {
		return tf[keys[i]] > tf[keys[j]]
	})
	return keys
}

func Index(text string) TermFreq {
	var result = map[string]int{}
	var lexer = &Lexer{ content : text, }
	for token, ok := lexer.Next(); ok; token, ok = lexer.Next() {
		if _, ok := result[token]; ok {
			result[token] += 1
 		} else {
 			result[token] = 1
 		}
	}
	return result
}

func main() {
	var result any

	var dirPath = "./docs.gl/gl4"
	result = Unwrap(os.ReadDir(dirPath))

	var tfIndex = TermFreqIndex{}

	for _, file := range result.([]os.DirEntry) {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".xhtml") {
			var filePath = dirPath + "/" + file.Name()
			var content = ReadXMLContent(string(Unwrap(os.ReadFile(filePath)).([]byte)))
			fmt.Println("Indexing", filePath)
			var tf = Index(content)
			fmt.Println(filePath, "has", len(tf), "unique tokens")
			tfIndex[filePath] = tf
		}
	}

	result = Unwrap(json.MarshalIndent(tfIndex, "", "\t"))

	Check(os.WriteFile("output.json", result.([]byte), 0666))
}