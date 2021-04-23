package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	//createreverseindex("files")
	sfi("output.txt", "vestibulum")
}

func createreverseindex(dir string) {
	ss := make(map[string][]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := file.Name()
		content, err := ioutil.ReadFile(dir + string(os.PathSeparator) + fn)
		if err != nil {
			log.Fatal(err)
		}

		wl := strings.Split(cleantext(string(content)), " ")

		for i := range wl {
			mapitem := ss[wl[i]]
			mapitem = append(mapitem, fn)
			ss[wl[i]] = mapitem
		}
	}
	wf("output.txt", toJson(ss))
}

func cleantext(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	pS := reg.ReplaceAllString(s, " ")

	return pS
}

func toJson(searchlist map[string][]string) string {
	j, err := json.Marshal(searchlist)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	output := strings.ReplaceAll(string(j), "],", "],\n")
	return output
}

func wf(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func sfi(indf string, st string) {
	content, err := ioutil.ReadFile(indf)
	if err != nil {
		log.Fatal(err)
	}
	var terms map[string][]string
	json.Unmarshal([]byte(content), &terms)

	words := strings.Split(cleantext(string(st)), " ")
	var ffound [][]string
	for _, word := range words {
		ffound = append(ffound, terms[word])
	}

	i := 0
	var final []string
	if len(ffound) > 1 {
		for ; i <= (len(ffound) - 2); i++ {
			final = union(ffound[i], ffound[i+1])
		}
	} else {
		final = ffound[0]
	}

	fmt.Println(removeDuplicates(final))
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
