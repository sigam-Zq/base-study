package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/*

create table tmp_menu
(
    `code` varchar(32)  DEFAULT NULL COMMENT '菜单标识（同等级唯一）',
    `parent_code` varchar(128) DEFAULT NULL COMMENT 'code树（父子关系）',
    `name` varchar(128) DEFAULT NULL COMMENT '展示菜单名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单';

*/

type Route struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Routes []*Route `json:"routes"`
}

func main() {
	dataR, err := os.OpenFile("./data/menu.json", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	// data := make([]byte, 0)

	// var buff = make([]byte, 1<<10)
	// b := bytes.NewBuffer(buff)

	// for n, err := b.ReadFrom(dataR); n != 0 && err != io.EOF; {
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	data = append(data, b.Bytes()...)
	// }

	/*

				The Go code you provided is intended to read data from a file (`menu.json`) into a byte slice. However, there are a few issues that could lead to inefficiencies or even a deadlock:

		1. **Buffer Initialization**: The buffer `b` is initialized with a pre-allocated slice `buff`, but `bytes.NewBuffer(buff)` creates a buffer with `buff` as its initial content, not as its capacity. This means the buffer is already "full" and `ReadFrom` will not append new data to it.

		2. **Loop Condition**: The loop condition `n != 0 && err != io.EOF` is problematic. If `b.ReadFrom(dataR)` reads no data (`n == 0`) but does not encounter an EOF, the loop will continue indefinitely, potentially causing a deadlock.

		3. **Data Reading**: The `ReadFrom` method reads from `dataR` into the buffer `b`, but since `b` is initialized with a full buffer, it won't actually read any new data.

		Here's a revised version of the code that addresses these issues:

		```go
		package main

		import (
			"bytes"
			"fmt"
			"io"
			"os"
		)

		func main() {
			dataR, err := os.OpenFile("./data/menu.json", os.O_RDONLY, 0)
			if err != nil {
				panic(err)
			}
			defer dataR.Close()

			var data []byte
			buff := make([]byte, 1<<10) // 1 KB buffer

			for {
				n, err := dataR.Read(buff)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				data = append(data, buff[:n]...)
			}

			fmt.Println("Data read:", string(data))
		}
		```

		### Key Changes:
		- **Buffer Usage**: The buffer `buff` is used directly in the `Read` method, which reads data into `buff` and returns the number of bytes read.
		- **Loop Condition**: The loop continues until `n == 0`, which indicates that no more data is available to read.
		- **Data Appending**: Only the bytes actually read (`buff[:n]`) are appended to `data`.

		This should resolve the memory inefficiency and potential deadlock issues.

	*/

	var data []byte
	buff := make([]byte, 1<<10) // 1 KB buffer

	for {
		n, err := dataR.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		data = append(data, buff[:n]...)
	}

	var listRoute []*Route

	err = json.Unmarshal(data, &listRoute)
	if err != nil {
		panic(err)
	}

	for _, v := range listRoute {
		fmt.Println(v.Name)

	}

}
