package main

import (
	"os"
	"strconv"

	"github.com/tealeg/xlsx/v3"
)

func main() {

	f := xlsx.NewFile()
	sheet, _ := f.AddSheet("Sheet1")
	sheet.AddRow().WriteSlice([]string{"序号", "上报人", "联系方式", "所属网格"}, -1)

	for i := 0; i < 10; i++ {
		sheet.AddRow().WriteSlice(&[]string{
			strconv.Itoa(i + 1),
			strconv.Itoa(i + 1),
			strconv.Itoa(i + 1),
			strconv.Itoa(i+1) + "----------------",
		}, -1)
	}
	style := xlsx.NewStyle()
	font := *xlsx.NewFont(12, "Verdana")
	font.Bold = true
	font.Italic = true
	font.Underline = true
	font.Strike = true
	style.Font = font
	fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
	style.Fill = fill
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = border
	style.ApplyBorder = true
	style.ApplyFill = true
	style.ApplyFont = true

	file, err := os.Create("data/a.xlsx")
	if err != nil {
		panic(err)
	}
	// f.
	err = f.Write(file)
	if err != nil {
		panic(err)
	}

}
