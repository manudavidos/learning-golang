package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createreverseindex("files")
}

func createreverseindex(dir string) {
	//ss := make(map[string][]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s contents: %s\n", file.Name(), content)
	}
}
