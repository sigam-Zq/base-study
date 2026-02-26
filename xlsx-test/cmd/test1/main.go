package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx/v3"
)

// 字典类型定义
type DictType string

const (
	DictIndustry DictType = "industry" // 行业字典
	DictScale    DictType = "scale"    // 企业规模
	DictStatus   DictType = "status"   // 企业状态
	DictProvince DictType = "province" // 省份
	DictCity     DictType = "city"     // 城市
)

// 字典数据
var DictData = map[DictType][]string{
	DictIndustry: {"互联网", "金融", "教育", "医疗", "制造", "零售", "房地产", "其他"},
	DictScale:    {"微型企业", "小型企业", "中型企业", "大型企业", "集团企业"},
	DictStatus:   {"正常", "异常", "注销", "停业"},
	DictProvince: {"北京市", "上海市", "广东省", "江苏省", "浙江省", "山东省"},
	DictCity:     {"北京市", "上海市", "广州市", "深圳市", "杭州市", "南京市", "青岛市"},
}

// 企业数据结构体
type Company struct {
	ID          int64   `excel:"企业ID" dict:"-" note:"系统自动生成的企业唯一标识"`
	Name        string  `excel:"企业名称" dict:"-" note:"请输入企业全称" validate:"required"`
	Industry    string  `excel:"所属行业" dict:"industry" note:"请从下拉列表选择" validate:"required"`
	Scale       string  `excel:"企业规模" dict:"scale" note:"请从下拉列表选择" validate:"required"`
	Status      string  `excel:"企业状态" dict:"status" note:"请从下拉列表选择" validate:"required"`
	LegalPerson string  `excel:"法定代表人" dict:"-" note:"请输入法定代表人姓名" validate:"required"`
	Phone       string  `excel:"联系电话" dict:"-" note:"请输入企业联系电话" validate:"phone"`
	Email       string  `excel:"邮箱地址" dict:"-" note:"请输入企业邮箱" validate:"email"`
	Province    string  `excel:"所在省份" dict:"province" note:"请从下拉列表选择" validate:"required"`
	City        string  `excel:"所在城市" dict:"city" note:"请从下拉列表选择" validate:"required"`
	Address     string  `excel:"详细地址" dict:"-" note:"请输入详细地址" validate:"required"`
	Revenue     float64 `excel:"年营收(万元)" dict:"-" note:"请输入年营收金额，单位：万元" validate:"min=0"`
	EmployeeNum int     `excel:"员工人数" dict:"-" note:"请输入员工总数" validate:"min=0"`
	CreateTime  string  `excel:"创建时间" dict:"-" note:"系统自动记录"`
}

// Excel导入导出管理器
type ExcelManager struct {
	file           *xlsx.File
	headerStyle    *xlsx.Style
	dataStyle      *xlsx.Style
	dictValidators map[string]*xlsx.DataValidation
}

// 新建Excel管理器
func NewExcelManager() *ExcelManager {
	em := &ExcelManager{
		file:           xlsx.NewFile(),
		dictValidators: make(map[string]*xlsx.DataValidation),
	}
	em.initStyles()
	em.initValidators()
	return em
}

// 初始化样式
func (em *ExcelManager) initStyles() {
	// 头部样式：加粗、居中、背景色
	em.headerStyle = xlsx.NewStyle()
	em.headerStyle.Font.Bold = true
	em.headerStyle.Alignment.Horizontal = "center"
	em.headerStyle.Alignment.Vertical = "center"
	em.headerStyle.Fill.PatternType = "solid"
	em.headerStyle.Fill.BgColor = "00CCCCCC"
	em.headerStyle.Fill.FgColor = "00CCCCCC"
	em.headerStyle.Border.Left = "thin"
	em.headerStyle.Border.Right = "thin"
	em.headerStyle.Border.Top = "thin"
	em.headerStyle.Border.Bottom = "thin"

	// 数据样式
	em.dataStyle = xlsx.NewStyle()
	em.dataStyle.Border.Left = "thin"
	em.dataStyle.Border.Right = "thin"
	em.dataStyle.Border.Top = "thin"
	em.dataStyle.Border.Bottom = "thin"
}

// 初始化字典验证器
func (em *ExcelManager) initValidators() {
	for dictType, options := range DictData {
		dv := xlsx.NewDataValidation(1, 1000, 0, 0, true)
		dv.SetDropList(options)
		dv.SetInputTitle("请选择")
		dv.SetInputMessage("请从下拉列表中选择合适选项")
		dv.SetErrorStyle(xlsx.DataValidationErrorStyleStop)
		em.dictValidators[string(dictType)] = dv
	}
}

// 生成导入模板
func (em *ExcelManager) GenerateTemplate(filename string) error {
	sheet, err := em.file.AddSheet("企业数据导入模板")
	if err != nil {
		return err
	}

	// 添加标题行
	titleRow := sheet.AddRow()
	titleCell := titleRow.AddCell()
	titleCell.SetString("企业数据导入模板")
	titleCell.Merge(0, len(getExcelTags(Company{}))-1)

	titleStyle := xlsx.NewStyle()
	titleStyle.Font.Bold = true
	titleStyle.Font.Size = 16
	titleStyle.Alignment.Horizontal = "center"
	titleCell.SetStyle(titleStyle)

	// 添加说明行
	noteRow := sheet.AddRow()
	noteCell := noteRow.AddCell()
	noteCell.SetString("填写说明：红色星号(*)为必填项，灰色字段为字典选择项")
	noteCell.Merge(0, len(getExcelTags(Company{}))-1)

	noteStyle := xlsx.NewStyle()
	noteStyle.Font.Color = "FF0000"
	noteStyle.Font.Italic = true
	noteCell.SetStyle(noteStyle)

	sheet.AddRow() // 空行

	// 添加表头
	headerRow := sheet.AddRow()
	fieldInfos := em.getFieldInfos()

	for _, info := range fieldInfos {
		cell := headerRow.AddCell()
		cell.SetString(info.ExcelTag)

		// 必填项添加红色星号
		if info.Required {
			cell.SetString(info.ExcelTag + " *")
			requiredStyle := *em.headerStyle
			requiredStyle.Font.Color = "FFFF0000"
			cell.SetStyle(&requiredStyle)
		} else {
			cell.SetStyle(em.headerStyle)
		}
	}

	// 添加数据验证和备注
	em.addDataValidations(sheet, fieldInfos)

	// 设置列宽和冻结窗格
	em.adjustColumnWidth(sheet, fieldInfos)
	sheet.SetPanes(&xlsx.Panes{
		Freeze:      true,
		Split:       false,
		X:           0,
		Y:           4, // 冻结前4行（标题、说明、空行、表头）
		TopLeftCell: "A5",
		ActivePane:  "bottomLeft",
	})

	return em.file.Save(filename)
}

// 获取字段信息
func (em *ExcelManager) getFieldInfos() []FieldInfo {
	var infos []FieldInfo
	t := reflect.TypeOf(Company{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		excelTag := field.Tag.Get("excel")
		dictTag := field.Tag.Get("dict")
		noteTag := field.Tag.Get("note")
		validateTag := field.Tag.Get("validate")

		if excelTag != "" && excelTag != "-" {
			info := FieldInfo{
				FieldName: field.Name,
				ExcelTag:  excelTag,
				DictType:  dictTag,
				Note:      noteTag,
				Required:  strings.Contains(validateTag, "required"),
			}
			infos = append(infos, info)
		}
	}
	return infos
}

// 添加数据验证
func (em *ExcelManager) addDataValidations(sheet *xlsx.Sheet, fieldInfos []FieldInfo) {
	for col, info := range fieldInfos {
		if info.DictType != "" && info.DictType != "-" {
			if dv, exists := em.dictValidators[info.DictType]; exists {
				// 设置验证范围从第5行开始（跳过标题行）
				dv.SetRange(col, col, 4, 1000)
				sheet.AddDataValidation(dv)
			}
		}
	}
}

// 调整列宽
func (em *ExcelManager) adjustColumnWidth(sheet *xlsx.Sheet, fieldInfos []FieldInfo) {
	for col, info := range fieldInfos {
		width := 15.0 // 默认宽度
		switch {
		case len(info.ExcelTag) > 8:
			width = 20.0
		case info.DictType != "" && info.DictType != "-":
			width = 12.0 // 字典字段稍窄
		}
		if err := sheet.SetColWidth(col, col, width); err != nil {
			fmt.Printf("设置列宽失败: %v\n", err)
		}
	}
}

// 导入企业数据
func (em *ExcelManager) ImportCompanies(filename string) ([]Company, error) {
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	var companies []Company
	fieldMapping := em.createFieldMapping()

	for _, sheet := range file.Sheets {
		// 跳过空行和标题行，从第5行开始读取数据（基于模板结构）
		for i := 4; i <= sheet.MaxRow; i++ {
			row, err := sheet.Row(i)
			if err != nil {
				continue
			}

			// 检查是否为空行
			if em.isEmptyRow(row) {
				continue
			}

			company, err := em.parseRowToCompany(row, fieldMapping)
			if err != nil {
				fmt.Printf("解析第%d行数据失败: %v\n", i+1, err)
				continue
			}

			if company.Name != "" { // 基本验证：企业名称不为空
				companies = append(companies, company)
			}
		}
	}

	return companies, nil
}

// 创建字段映射
func (em *ExcelManager) createFieldMapping() map[string]int {
	mapping := make(map[string]int)
	fieldInfos := em.getFieldInfos()

	for i, info := range fieldInfos {
		mapping[info.ExcelTag] = i
	}
	return mapping
}

// 解析行数据到Company结构体
func (em *ExcelManager) parseRowToCompany(row *xlsx.Row, fieldMapping map[string]int) (Company, error) {
	var company Company
	v := reflect.ValueOf(&company).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		excelTag := field.Tag.Get("excel")

		if excelTag == "" || excelTag == "-" {
			continue
		}

		if colIndex, exists := fieldMapping[excelTag]; exists && colIndex < row.Sheet.MaxCol {
			cell := row.GetCell(colIndex)
			if cell != nil {
				cellValue := strings.TrimSpace(cell.Value)
				fieldValue := v.Field(i)

				if fieldValue.CanSet() {
					em.setFieldValue(fieldValue, cellValue)
				}
			}
		}
	}

	return company, nil
}

// 设置字段值
func (em *ExcelManager) setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int64:
		var intVal int64
		fmt.Sscanf(value, "%d", &intVal)
		field.SetInt(intVal)
	case reflect.Float64:
		var floatVal float64
		fmt.Sscanf(value, "%f", &floatVal)
		field.SetFloat(floatVal)
	}
}

// 检查是否为空行
func (em *ExcelManager) isEmptyRow(row *xlsx.Row) bool {
	if row == nil {
		return true
	}

	for i := 0; i < row.Sheet.MaxCol; i++ {
		cell := row.GetCell(i)
		if cell != nil && strings.TrimSpace(cell.Value) != "" {
			return false
		}
	}
	return true
}

// 导出企业数据
func (em *ExcelManager) ExportCompanies(companies []Company, filename string) error {
	sheet, err := em.file.AddSheet("企业数据")
	if err != nil {
		return err
	}

	// 添加表头
	headerRow := sheet.AddRow()
	fieldInfos := em.getFieldInfos()

	for _, info := range fieldInfos {
		cell := headerRow.AddCell()
		cell.SetString(info.ExcelTag)
		cell.SetStyle(em.headerStyle)
	}

	// 添加数据行
	for _, company := range companies {
		dataRow := sheet.AddRow()
		v := reflect.ValueOf(company)

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			excelTag := v.Type().Field(i).Tag.Get("excel")

			if excelTag != "" && excelTag != "-" {
				cell := dataRow.AddCell()
				em.setCellValue(cell, field)
				cell.SetStyle(em.dataStyle)
			}
		}
	}

	// 设置冻结窗格
	sheet.SetPanes(&xlsx.Panes{
		Freeze:      true,
		Split:       false,
		X:           0,
		Y:           1, // 冻结表头
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	return em.file.Save(filename)
}

// 设置单元格值
func (em *ExcelManager) setCellValue(cell *xlsx.Cell, field reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		cell.SetString(field.String())
	case reflect.Int, reflect.Int64:
		cell.SetInt64(field.Int())
	case reflect.Float64:
		cell.SetFloat(field.Float())
	default:
		cell.SetString(fmt.Sprintf("%v", field.Interface()))
	}
}

// 字段信息结构
type FieldInfo struct {
	FieldName string
	ExcelTag  string
	DictType  string
	Note      string
	Required  bool
}

// 获取Excel标签
func getExcelTags(s interface{}) []string {
	var tags []string
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get("excel"); tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
	}
	return tags
}

// 使用示例
func main() {
	em := NewExcelManager()

	// 生成导入模板
	err := em.GenerateTemplate("企业数据导入模板.xlsx")
	if err != nil {
		fmt.Printf("生成模板失败: %v\n", err)
	} else {
		fmt.Println("模板生成成功: 企业数据导入模板.xlsx")
	}

	// 导入数据示例
	companies, err := em.ImportCompanies("企业数据导入模板.xlsx")
	if err != nil {
		fmt.Printf("导入数据失败: %v\n", err)
	} else {
		fmt.Printf("成功导入 %d 条企业数据\n", len(companies))
	}

	// 导出数据示例
	err = em.ExportCompanies(companies, "企业数据导出.xlsx")
	if err != nil {
		fmt.Printf("导出数据失败: %v\n", err)
	} else {
		fmt.Println("数据导出成功: 企业数据导出.xlsx")
	}
}
