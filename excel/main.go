package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

func main() {

	var (
		ProductList []string

		file   *xlsx.File
		sheet1 *xlsx.Sheet
		row1   *xlsx.Row
		cell1  *xlsx.Cell
		cell2  *xlsx.Cell
	)

	excelFileName := "C:/test/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open excel file failed.")
		// create file
		f, err := os.create("C:/test/log.txt")
		check(err)
		defer f.close()
		// write file
		n3, err := f.WriteString("open excel file failed\n")
		fmt.Printf("wrote %d bytes\n", n3)
		f.Sync()
		// exit
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
						arr := strings.Split(value, " ")
						if arr != nil {
							row1 = sheet1.AddRow()
							cell1 = row1.AddCell()
							cell1.Value = arr[0]
							cell2 = row1.AddCell()
							if le := len(arr); le > 3 {
								cell2.Value = arr[le-3]
							}
						}
					}
				}
			}
		}
	}

	err = file.Save("C:/test/result.xlsx")
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
