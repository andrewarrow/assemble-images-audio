package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("assemble-images-audio <directory>")
	if len(os.Args) == 1 {
		return
	}
	files, _ := ioutil.ReadDir(os.Args[1] + "/images")
	m := map[string][]int{}
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".png") {
			three := name[0:3]
			min := name[4:6]

			mini, _ := strconv.Atoi(min)
			m[three] = append(m[three], mini)
		}
	}
	fmt.Println(m)
}
