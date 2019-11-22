package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

func main() {

	const path = "/workspace/tmp/result.xlsx"
	var (
		ProductList []string

		file   *xlsx.File
		sheet1 *xlsx.Sheet
		row1   *xlsx.Row
		cell1  *xlsx.Cell
		cell2  *xlsx.Cell
	)

	excelFileName := "/Users/subway/Desktop/test_bak.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open excel file failed.")
		// create file
		f, err := os.Create("C:/test/log.txt")
		if err != nil {
			log.Error("log error.")
		}
		defer f.Close()
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
								for m, vv := range arr {
									fmt.Println(arr)
									fmt.Println(le)
									fmt.Println(arr[le-3])
									fmt.Println(m)
									fmt.Println(vv)
									fmt.Println("------------")
								}
							}
						}
					}
				}
			}
		}
	}

	err = file.Save(path)
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
