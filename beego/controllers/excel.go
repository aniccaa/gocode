package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

type ExcelController struct {
	beego.Controller
}

func (this *ExcelController) Get() {
	this.Layout = "excel.tpl"
	this.TplName = "excel.tpl"
}

func (this *ExcelController) Post() {
	f, h, e := this.GetFile("name")
	if e != nil {
		log.Fatal("get file error", e)
	}
	defer f.Close()

	err := this.SaveToFile("name", "static/upload/"+h.Filename) // 保存在／static/upload，没有文件夹先创建
	if err != nil {
		log.Fatal("save to file failed", err)
		return
	}

	// 对excel表进行处理
	if err := processExcel1(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("process excel file error.")
		return
	}

	this.Ctx.Redirect(302, "/success")
	// 下载处理后的文件
	// this.Ctx.Output.Download("static/download/result.xlsx", "result.xlsx")
}

func processExcel() error {

	const path = "static/download/result.xlsx"
	var (
		ProductList []string

		file   *xlsx.File
		sheet1 *xlsx.Sheet
		row1   *xlsx.Row
		cell1  *xlsx.Cell
		cell2  *xlsx.Cell
	)

	excelFileName := "static/upload/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open excel file failed.")
		return err
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

	err = file.Save(path)
	if err != nil {
		return err
	}
	return nil
}

func processExcel1() error {
	const path = "static/download/result.xlsx"
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

	excelFileName := "static/upload/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.WithFields(log.Fields{"xlFile": xlFile, "error": err}).Info("open excel file failed.")
		return err
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
		return err
	}
	return nil
}
