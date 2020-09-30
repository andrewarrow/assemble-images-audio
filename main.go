package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("assemble-images-audio <directory>")
	if len(os.Args) == 1 {
		return
	}
	files, _ := ioutil.ReadDir(os.Args[1])
	prevNum := "000"
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".png") {
			three := name[0:3]
			if three != prevNum {
				fmt.Println(three, name)
			}
			prevNum = three
		}
	}
	exec.Command("cat_images.sh").Run()
}
