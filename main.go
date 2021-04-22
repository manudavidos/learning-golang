package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

		wl := strings.Split(string(content), " ")

		for i := range wl {
			arr := ss[wl[i]]
			arr = append(arr, fn)
			ss[wl[i]] = arr
		}
	}
	fmt.Println(ss)
}
