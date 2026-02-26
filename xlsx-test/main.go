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
	// style := xlsx.NewStyle()
	// font := *xlsx.NewFont(12, "Verdana")
	// font.Bold = true
	// font.Italic = true
	// font.Underline = true
	// font.Strike = true
	// style.Font = font
	// fill := *xlsx.NewFill("none", "", "0000FF")
	// style.Fill = fill
	// border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	// style.Border = border
	// style.ApplyBorder = true
	// style.ApplyFill = true
	// style.ApplyFont = true

	style := xlsx.NewStyle()
	style.Fill = *xlsx.NewFill("solid", "a3a3f3", "a3a3f3")
	r, err := f.Sheet["Sheet1"].Row(0)
	if err != nil {
		panic(err)
	}

	r.GetCell(1).SetStyle(style)

	dv := xlsx.NewDataValidation(1, 3, 999, 3, true)
	dv.SetDropList([]string{"夜班", "白班"})

	f.Sheet["Sheet1"].AddDataValidation(dv)
	// f.Sheet["Sheet1"].Cols.Add(xlsx.NewColForRange(0, 1))
	// fmt.Println("Cols", f.Sheet["Sheet1"].Cols)
	// c := f.Sheet["Sheet1"].Col(0)
	// fmt.Println("c", c)
	// style2 := xlsx.NewStyle()
	// style2.Fill = *xlsx.NewFill("solid", "00FF0000", "0000ff")
	// c.SetStyle(style2)

	file, err := os.Create("data/a2.xlsx")
	if err != nil {
		panic(err)
	}
	// f.
	err = f.Write(file)
	if err != nil {
		panic(err)
	}

}
