package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("tmp/imgUp.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 读取第一个工作表的名称
	sheetName := f.GetSheetName(0)
	fmt.Println("Sheet Name:", sheetName)

	// 读取单元格数据

	// row, err := f.GetRows(sheetName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// cellValue, err := f.GetCellValue(sheetName, "A2")
	// fmt.Println("row[1][1]:", row[1][1])
	picCells, err := f.GetPictureCells(sheetName)
	fmt.Println("picCells:", picCells)

	for _, picCell := range picCells {

		pic, err := f.GetPictures(sheetName, picCell)
		if err != nil {
			fmt.Println(err)
			return
		}
		nameCell := strings.Replace(picCell, "B", "A", -1)
		name, err := f.GetCellValue(sheetName, nameCell)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile(fmt.Sprintf("tmp/data/%s.jpeg", name), pic[0].File, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// fmt.Println("pic:", pic)

}
