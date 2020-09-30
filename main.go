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
	/*
		for i := 0; i < len(m); i++ {
			key := fmt.Sprintf("%03d", i+1)
			size := len(m[key])
			mapOfInts := map[int]bool{}
			for _, v := range m[key] {
				mapOfInts[v] = true
			}
			for i, v := range m[key] {
				if i == size-1 {
					for j := 0; j < 5; j++ {
						doCopy(dir, key, v, v+j)
					}
				} else {
					vstart := v
					for {
						doCopy(dir, key, v, vstart)
						vstart++
						if mapOfInts[vstart] == false {
							doCopy(dir, key, v, vstart)
						} else {
							break
						}
					}
				}
			}
		}
	*/
}

func doCopy(dir, key, order, seconds string, vstart int) {
	exec.Command("cp", fmt.Sprintf("%s/%s_%s-%s.png", dir, key, order, seconds),
		fmt.Sprintf("video/%s_%02d.png", key, vstart)).Run()
}
