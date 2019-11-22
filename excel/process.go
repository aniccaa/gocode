package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

func main() {
	const path = "/Users/subway/Desktop/result.xlsx"
	var (
		ProductList []string

		file   *xlsx.File
		sheet1 *xlsx.Sheet
		row1   *xlsx.Row // sheet的title
		row2   *xlsx.Row
		cell1  *xlsx.Cell // 买家姓名列
		cell2  *xlsx.Cell // 商品详情列
		cell3  *xlsx.Cell // 数量列
	)

	excelFileName := "/Users/subway/Desktop/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open excel file failed.")
		return
	}

	// 生成新的excel保存结果集
	file = xlsx.NewFile()
	sheet1, err = file.AddSheet("Sheet1")
	if err != nil {
		log.WithFields(log.Fields{"sss": sheet1}).Info("xxx")
		fmt.Printf(err.Error())
	}

	var customColumn, productColumn int
	var customer string
	for num, sheet := range xlFile.Sheets {
		if num == 0 {
			// 打印sheet的title
			// 增加买家姓名title
			row1 = sheet1.AddRow()
			cell1 = row1.AddCell()
			cell1.Value = "买家姓名"
			// 增加商品详情title
			cell2 = row1.AddCell()
			cell2.Value = "商品详情"
			// 增加数量title
			cell3 = row1.AddCell()
			cell3.Value = "数量"
		}

		for i, row := range sheet.Rows {
			// sheet的第一行
			if i == 0 {
				for j, cell := range row.Cells {
					fmt.Println(cell.String())
					// 判断title
					switch cell.String() {
					case "买家姓名":
						customColumn = j
						//fmt.Println(j)
					case "商品详情":
						productColumn = j
						//fmt.Println(j)
					default:
						continue
					}

				}
			}

			if i > 0 {
				for j, cell := range row.Cells {
					// 每一行的j列
					if j == customColumn {
						// 对买家姓名列进行处理
						customer = cell.Value
					}
					if j == productColumn {
						// 对商品详情列进行处理
						ProductList = strings.Split(cell.String(), "||")
						for _, value := range ProductList {
							// example: 深海精萃保湿香氛身体乳 350ml * 3 盒 ||深海精萃保湿香氛沐浴露 350ml * 2 盒 ||
							// arr: [深海精萃保湿香氛身体乳 350ml * 3 盒 ]
							// arr[0] = "深海精萃保湿香氛身体乳", arr[1] = "350ml", arr[2] = "*", arr[3] = "3"
							arr := strings.Split(value, " ")
							if arr != nil {
								row2 = sheet1.AddRow()

								cell1 = row2.AddCell()
								cell1.Value = customer

								cell2 = row2.AddCell()
								cell2.Value = arr[0] + arr[1]

								cell3 = row2.AddCell()
								cell3.Value = arr[3]
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
