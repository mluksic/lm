package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("f", "", "Target file")
	prefix := flag.Bool("p", false, "Should prefix to lines")
	suffix := flag.Bool("a", false, "Should suffix to lines")
	content := flag.String("c", "", "Content to prefix/append")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(0)
	}
	defer file.Close()

	newFile, err := os.Create(*fileName + "_output")
	defer newFile.Close()

	r := bufio.NewReader(file)
	w := bufio.NewWriter(newFile)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		var newLine string
		if *prefix == true {
			newLine = *content + string(line)
		}
		if *suffix == true {
			newLine = string(line) + *content
		}

		w.Write([]byte(newLine + "\n"))
	}

	os.Exit(1)
}
