package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	fmt.Println("Hello world")

	fileName := "E:\\temp\\3.xlsx"

	xlFile, err := xlsx.OpenFile(fileName)
	if nil != err {
		fmt.Println("error");
		return
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				str, _ := cell.String()
				fmt.Printf("%s\t", str)
			}
			fmt.Println()
		}
	}
}
