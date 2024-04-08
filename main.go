package main

import (
	"fmt"
	"io"
	"os"
)

func writeToFile() (string, error) {
	dirname, err := os.MkdirTemp("", "testingdb")
	if err != nil {
		return "", err
	}

	// defer os.RemoveAll(dirname)
	qname := "sampledb.txt"

	absPath := dirname + "/" + qname

	f, _ := os.Create(absPath)
	defer f.Close()

	for i := 0; i < 10; i++ {

		offset, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			return "", err
		}

		fmt.Println(offset)

		if i%2 == 0 {
			_, _ = f.Write([]byte("hello"))
		} else {
			_, _ = f.Write([]byte("file seek test"))
		}

	}

	return absPath, nil
}

func readFileAtOffset(path string) (string, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}

	defer f.Close()

	// give any whence in seek
	offset, err := f.Seek(0, 10)
	if err != nil {
		return "", err
	}

	var b = make([]byte, 5)
	if _, err := f.ReadAt(b, offset); err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered error: ", r)
		}
	}()

	fmt.Println("\n---------Seek test from Go---------")

	path, err := writeToFile()
	if err != nil {
		panic(err)
	}

	{
		// this is for the C program to read the file, to get the path created by the golang
		fd, err := os.Create("path.txt")
		if err != nil {
			panic(err)
		}

		_, err = fd.Write([]byte(path))
		if err != nil {
			panic(err)
		}
	}

	data, err := readFileAtOffset(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("data read:", data)
}
