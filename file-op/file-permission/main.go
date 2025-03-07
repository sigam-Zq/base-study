package main

import (
	"fmt"
	"os"
)

func isPermission() {

	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666) //直接给O_WRONLY标志，可以判断是否有写权限

	defer file.Close()

	if err != nil {

		if os.IsPermission(err) {

			fmt.Println("没有读取权限：", err)

		}

	}

	fmt.Println("可读")

	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666) //直接给O_WRONLY标志，可以判断是否有写权限

	defer file.Close()

	if err != nil {

		if os.IsPermission(err) {

			fmt.Println("没有写权限", err)

		}

	}

	fmt.Println("可写")

}
