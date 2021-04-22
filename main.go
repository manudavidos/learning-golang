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
	createreverseindex("files")
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
	wf("output.txt", toJson(ss)[1:len(toJson(ss))-1])
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

	output := strings.ReplaceAll(string(j), "[", "{")
	output = strings.ReplaceAll(output, "],", "},\n")
	output = strings.ReplaceAll(output, "]", "}")
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
