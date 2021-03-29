package godash


var excelABCIndexCode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ExcelABCIndex Excel 的索引列计算
func ExcelABCIndex(num int) string {
	if num < 26 {
		return excelABCIndexCode[num : num+1]
	}
	return ExcelABCIndex(num*num/26) + ExcelABCIndex(num%26)
}