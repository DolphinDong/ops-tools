package excel_tools

import (
	"github.com/DolphinDong/ops-tools/tools"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

var (
	result = map[string][]ExcelRow{}
)

type ExcelRow []string

func SheetSort() (err error) {
	excel, err := excelize.OpenFile(Options.InputExcelPath)
	if err != nil {
		return errors.WithStack(err)
	}
	rows, err := excel.GetRows(Options.SheetSort.SheetName)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(rows) == 0 {
		Logger.Infof("excel %v(%v) is empty", Options.InputExcelPath, Options.SheetSort.SheetName)
		return
	}
	targetFieldIndex := -1
	targetFieldInSheet := false
	if targetFieldIndex, targetFieldInSheet = tools.StringInSlice(Options.SheetSort.SortField, rows[0]); !targetFieldInSheet {
		return errors.Errorf("field: %v not fount in %v(%v)", Options.SheetSort.SortField, Options.InputExcelPath, Options.SheetSort.SheetName)
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		targetFieldValue := row[targetFieldIndex]
		if _, exist := result[targetFieldValue]; exist {
			result[targetFieldValue] = append(result[targetFieldValue], row)
		} else {
			result[targetFieldValue] = []ExcelRow{row}
		}
	}
	for targetFieldValue, data := range result {
		baseDir := Options.SheetSort.OutputPath
		if Options.SheetSort.NeedCreateSubDir {
			baseDir = filepath.Join(Options.SheetSort.OutputPath, targetFieldValue)
		}
		err = SaveData2Excel(rows[0], data, filepath.Join(baseDir, targetFieldValue+".xlsx"))
		if err != nil {
			return err
		}
	}
	return
}

func SaveData2Excel(header []string, data []ExcelRow, filePath string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	// 创建一个工作表
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return errors.WithStack(err)
	}
	for colIndex, value := range header {
		cellName, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		f.SetCellValue(sheetName, cellName, value)
	}
	// 当前excel没有数据时不创建
	hasData := false
	for rowIndex, rowData := range data {
		if NeedExcludeRow(Options.SheetSort.ExcludeRule, header, rowData) {
			continue
		}
		hasData = true
		for colIndex, value := range rowData {
			cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cellName, value)
		}
	}
	if hasData {
		_ = os.MkdirAll(filepath.Dir(filePath), os.ModeDir)

		if err := f.SaveAs(filePath); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func NeedExcludeRow(excludeRule map[string]string, header, rowData []string) bool {
	for fieldName, fieldValue := range excludeRule {
		if fieldIndex, exist := tools.StringInSlice(fieldName, header); exist {
			if rowData[fieldIndex] == fieldValue {
				return true
			}
		}
	}
	return false
}
