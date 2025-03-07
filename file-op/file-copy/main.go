package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	fileCopy("a.txt", "b.txt")

}

func fileCopy(src, dst string) error {

	srcFile, err := os.Open(src)

	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	var dstFile *os.File

	if _, err := os.Stat(dst); os.IsNotExist(err) {

		dstFile, err = os.Create(dst)
		if err != nil {
			panic(err)
		}
	} else {

		dstFile, err = os.OpenFile(dst, os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

	}

	defer srcFile.Close()
	fileLen, err := io.Copy(dstFile, srcFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("fileLen:", fileLen)

	time.Sleep(10 * time.Second)

	dstFile.Sync()
	return nil
}
