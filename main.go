package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("assemble-images-audio <directory>")
	if len(os.Args) == 1 {
		return
	}
	dir := os.Args[1] + "/images"
	files, _ := ioutil.ReadDir(dir)
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
	for i := 0; i < len(m); i++ {
		key := fmt.Sprintf("%03d", i+1)
		for _, v := range m[key] {
			fmt.Println(key, v)
			exec.Command("cp", fmt.Sprintf("%s/%s_%02d.png", dir, key, v),
				fmt.Sprintf("video/%s.png", "hi")).Run()
		}
	}
}
