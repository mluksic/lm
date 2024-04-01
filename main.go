package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("f", "", "Target file")
	prefix := flag.String("p", "", "Prefix the lines")
	suffix := flag.String("s", "", "Suffix the lines")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	newFile, err := os.Create(*fileName + "_output")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer newFile.Close()

	r := bufio.NewReader(file)
	w := bufio.NewWriter(newFile)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		var newLine string
		if *prefix != "" {
			newLine = *prefix + string(line)
		}
		if *suffix != "" {
			newLine = string(line) + *suffix
		}

		w.Write([]byte(newLine + "\n"))
	}

	err = w.Flush()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
