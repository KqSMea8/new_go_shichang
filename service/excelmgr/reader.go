package excelmgr

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
	"xinxin/service/model/dto"
)

type FileImporter struct {
	xlFile *xlsx.File
	colIndexMap map[string]int
	lines []dto.InputInfo
}

func (c *FileImporter) initWorkBook(path string) error {
	if xlFile, err := xlsx.OpenFile(path); err == nil {
		c.xlFile = xlFile
	}else{
		return err
	}

	return nil
}

func (c *FileImporter)LoadData(path string) ([]dto.InputInfo, error) {
	if err := c.initWorkBook(path); err != nil {
		return c.lines,err
	}

	if len(c.xlFile.Sheets) == 0{
		return c.lines,errors.New("没有发现Excel File")
	}

	c.setTitle()

	sheet1 := c.xlFile.Sheets[0]
	for idx, row := range sheet1.Rows {
		if idx == 0 {
			continue
		}
		c.fillDataByRow(row,idx)
	}

	err := c.validate()

	return c.lines,err
}

func (c *FileImporter)setTitle(){
	colMap := make(map[string]string)
	//colMap = viper.GetStringMapString("ColumnMap")
	colMap["NumberFlag"]="磅单编号"
	colMap["DateStr"]="首次称重日期"
	colMap["TimeStr"]="毛重时间"
	colMap["CarNumber"]="车号"
	colMap["ShopNumber"]="货物名称"
	colMap["GrossWeight"]="毛重"
	colMap["TareWeight"]="皮重"
	colMap["NetWeight"]="净重"
	colMap["Vendor"]="收货单位"
	colMap["SendOutDept"]="发货单位"

	c.colIndexMap = make(map[string]int)

	captionRow := c.xlFile.Sheets[0].Rows[0]
	for idx, cellValue := range captionRow.Cells {
		for prop, value := range colMap {
			if value == cellValue.String() {
				c.colIndexMap[prop] = idx
			}
		}
	}
}

func (c *FileImporter)fillDataByRow(row *xlsx.Row, index int) {
	var line dto.InputInfo
	pt := reflect.ValueOf(&line).Elem()

	for colN, colIdx := range c.colIndexMap {
		if colIdx < len(row.Cells) {
			val := row.Cells[colIdx]
			value := val.String()
			pt.FieldByName(colN).SetString(value)
		}

	}

	pt.FieldByName("UnitPrice").SetString("0")
	pt.FieldByName("RowIndex").SetString(strconv.Itoa(index))

	c.lines = append(c.lines,line)
}

func (c *FileImporter)validate() error {
	for _,item := range c.lines {
		fmt.Println(item)
	}
	return nil
}