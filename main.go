package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Reader interface {
	Read() ([]byte, error)
}

type StdinReader struct {
	reader *bufio.Reader
}

func NewStdinReader() *StdinReader {
	r := bufio.NewReader(os.Stdin)

	return &StdinReader{reader: r}
}
func (r *StdinReader) Read() ([]byte, error) {
	bytes, _, err := r.reader.ReadLine()

	return bytes, err
}

type FileReader struct {
	reader *bufio.Reader
}

func NewFileReader(filename string) *FileReader {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	r := bufio.NewReader(file)

	return &FileReader{reader: r}
}
func (r *FileReader) Read() ([]byte, error) {
	bytes, _, err := r.reader.ReadLine()

	return bytes, err
}

func main() {
	filename := flag.String("f", "", "Target file")
	prefix := flag.String("p", "", "Prefix the lines with content")
	suffix := flag.String("s", "", "Suffix the lines with content")
	flag.Parse()

	reader := initReader(*filename)

	newFile, err := os.Create("output")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer newFile.Close()

	w := bufio.NewWriter(newFile)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		var newLine string
		if *prefix != "" {
			newLine = *prefix + string(line)
		}
		if *suffix != "" {
			newLine = string(line) + *suffix
		}
		if *suffix != "" && *prefix != "" {
			newLine = *prefix + string(line) + *suffix
		}

		_, err = w.Write([]byte(newLine + "\n"))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	}

	err = w.Flush()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func initReader(filename string) Reader {
	switch {
	case "" != filename:
		return NewFileReader(filename)
	default:
		return NewStdinReader()
	}
}
