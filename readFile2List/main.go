package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var dictList []string

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dictFilePath := path + string(filepath.Separator) + "dict.txt"
	if !fileExist(dictFilePath) {
		panic(errors.New("is not read File"))
	}

	f, err := os.OpenFile(dictFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()
	var buf []byte
	tmpStr := ""
	// fmt.Println(1 << 10)
	for {
		if size > (1 << 10) {
			buf = make([]byte, 1024)
		} else {
			buf = make([]byte, size)
		}
		n, err := f.Read(buf)
		if err != nil {
			panic(err)
		}
		tmpStr += string(buf)
		size -= int64(n)
		if size == 0 {
			break
		}
	}
	// log.Println(tmpStr)
	dictList = strings.Split(tmpStr, "\n")

	fmt.Println("=")
	for i := len(dictList) - 1; i >= 0; i-- {
		dictList[i] = strings.TrimSpace(dictList[i])
		if dictList[i] == "" {
			dictList = append(dictList[:i], dictList[i+1:]...)
		}
	}
	fmt.Println("=")
	log.Printf("\n dictList --  %v  -- size %d cap %d\n", dictList, len(dictList), cap(dictList))
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
