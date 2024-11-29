package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	ReadCsv("data/三盛2号楼.csv")

}

func ReadCsv(file_path string) (res [][]string) {
	file, err := os.Open(file_path)
	if err != nil {
		log.Println("open_err:", err)
		return
	}
	defer file.Close()
	// 初始化csv-reader

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if isGBK(bytes) {
		fmt.Println("isGBK")
	}
	if isUtf8(bytes) {
		fmt.Println("isUtf8")
	}

	file2, err := os.Open(file_path)
	if err != nil {
		log.Println("open_err:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file2)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 允许懒引号（忘记遇到哪个问题才加的这行）
	reader.LazyQuotes = true
	// 返回csv中的所有内容
	record, read_err := reader.ReadAll()
	if read_err != nil {
		log.Println("read_err:", read_err)
		return
	}

	for i, v := range record {
		log.Printf(" i %d  record-line %v \n", i, v)
	}

	// for i, value := range record {
	// 	record_utf := value
	// 	if !ValidUTF8([]byte(record_utf)) {
	// 		record_utf, _, _, _ = gogb2312.ConvertGB2312String(record_utf)
	// 	}
	// 	record_utf = strings.TrimSpace(record_utf)
	// 	record[i] = record_utf
	// }
	return record
}

func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
