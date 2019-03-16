package ExcelWorker

import (
	Model "../models"
	"fmt"
	"github.com/tealeg/xlsx"
	"mime/multipart"
)

func ExcelAsSlice(r multipart.File, size int64) ([][][]string, error) {
	file, err := xlsx.OpenReaderAt(r,size)
	if err != nil {
		return nil, err
	}

	slice, err := file.ToSlice()
	if err != nil {
		return nil, err
	}

	return slice, err
}

func StudentsToExcel(students []Model.Student) *xlsx.File {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	xlsx.SetDefaultFont(11, "Calibri")

	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	sheet.SetColWidth(0, 3, 24)
	row = sheet.AddRow()
	row.SetHeight(12)
	row.WriteSlice(&[]string{"ФИО", "Email", "Пароль"}, 3)

	for _, v := range students {
		row = sheet.AddRow()
		row.SetHeight(60)
		row.WriteSlice(&[]string{v.Name, v.Email, v.Password}, 3)
	}

	return file

}

func ScoresToExcel(res []Model.Result) *xlsx.File {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	xlsx.SetDefaultFont(11, "Calibri")

	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	sheet.SetColWidth(0, 8, 24)
	row = sheet.AddRow()
	row.SetHeight(12)
	row.WriteSlice(&[]string{"ФИО", "Base", "Reading", "Writing", "Listening", "Sum", "Grade", "Recommended Level"}, 8)

	for _, res := range res {
		row = sheet.AddRow()
		row.SetHeight(60)
		row.WriteSlice(GetScoreSlice(res), 8)
	}

	return file
}

func GetScoreSlice(res Model.Result) *[]string {

	name := res.Name

	if(res.Score.Sum==0){
		slice := []string{name, "-", "-", "-", "-", "-", "", "-"}
		return &slice
	}

	base := fmt.Sprintf("%v\\%v", res.Score.Base, res.Score.BaseAmount)
	reading := fmt.Sprintf("%v\\%v", res.Score.Reading, res.Score.ReadingAmount)
	writing := fmt.Sprintf("%v\\%v", res.Score.Writing, res.Score.WritingAmount)
	listening := fmt.Sprintf("%v\\%v", res.Score.Listening, res.Score.ListeningAmount)
	sumReal := res.Score.Sum
	sumAmount := res.Score.SumAmount
	recLvl:=res.Score.RecommendedLevel

	var grade string
	if(sumReal==0) {
		grade = fmt.Sprintf("%v", 0)

	} else {
		grade = fmt.Sprintf("%.2f",float64(sumReal) / float64(sumAmount) *10)
	}
	sum := fmt.Sprintf("%v\\%v", sumReal, sumAmount)
	slice := []string{name, base, reading, writing, listening, sum, grade, recLvl}
	return &slice
}
