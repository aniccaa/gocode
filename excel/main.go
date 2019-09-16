package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)



func main()  {

	var (
		ProductList []string

		file *xlsx.File
		sheet1 *xlsx.Sheet
		row1 *xlsx.Row
		cell1 *xlsx.Cell
	)

	excelFileName := "/Users/subway/Desktop/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open xls failed.")
		os.Exit(1)
	}

	// 生成新的excel保存结果集
	file = xlsx.NewFile()
	sheet1, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}


	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					// fmt.Printf("%s\n", text)
					ProductList = strings.Split(text, "||")
					for _, value := range ProductList {
						fmt.Printf("%s//", value)
						// write(value)
						row1 = sheet1.AddRow()
						cell1 = row1.AddCell()
						cell1.Value = value
					}
				}
			}
		}
	}

	err = file.Save("/Users/subway/Desktop/result.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

}


//func write(value string) {
//
//
//	row = sheet.AddRow()
//	cell = row.AddCell()
//	cell.Value = value
//
//}