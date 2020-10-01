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
	exec.Command("rm", "video/*.png").Run()
	dir := os.Args[1] + "/images"
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".png") {
			three := name[0:3]
			order := name[4:6]
			seconds := name[7:9]

			fmt.Println(three, order, seconds)
			secondsi, _ := strconv.Atoi(seconds)

			for i := 0; i < secondsi; i++ {
				doCopy(dir, three, order, seconds, i)
			}
		}
	}
}

func doCopy(dir, key, order, seconds string, vstart int) {
	exec.Command("cp", fmt.Sprintf("%s/%s_%s-%s.png", dir, key, order, seconds),
		fmt.Sprintf("video/%s_%s-%02d.png", key, order, vstart)).Run()
}
