package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {

	object1 := readExecl("dstn.xlsx")
	object2 := readExecl("dstn2.xlsx")

	// // Same the row and column
	// if len(object1) == len(object2) {
	// 	res := compareSlices(object1, object2)
	// 	convertArrayArrayToExecl(res, "result.xlsx")
	// 	fmt.Println("Result", res)
	// }

	// // Not same the row or column
	// if len(object1) == len(object2) {
	// 	res := compareSlicesUpgrade(object1, object2)
	// 	// convertArrayArrayToExecl(res, "result.xlsx")
	// 	fmt.Println("Result", res)
	// }

	if len(object1) >= len(object2) {
		res := compareSlicesV2(object1, object2)
		convertArrayArrayToExecl(res, "result.xlsx")
		fmt.Println("Result", res)
	} else {
		res := compareSlicesV2(object2, object1)
		convertArrayArrayToExecl(res, "result.xlsx")
		fmt.Println("Result", res)
	}
}

func readExecl(nameFile string) [][]string {
	fileExecl, err := excelize.OpenFile(nameFile)
	if err != nil {
		fmt.Println("Open and read file execl error", err)
	}

	// Sheet1 == 0
	// Sheet2 == 1
	// Sheet3 == 2
	// Sheet4 == 3
	nameSheet := fileExecl.GetSheetName(3)
	if nameSheet == "" {
		fmt.Println("Sheet index invaild")
	}

	rowsFileExecl, err := fileExecl.GetRows(nameSheet)
	if err != nil {
		fmt.Println("Get data row from execl file error", err)
	}
	return rowsFileExecl
}

func compareSlices(slice1, slice2 [][]string) [][]string {
	for i := range slice1 {
		for j := range slice1[i] {
			// check if and insert DIFF into [][]string
			if slice1[i][j] != slice2[i][j] {
				fmt.Printf("Value is different: %v and %v \n", slice1[i][j], slice2[i][j])
				slice1[i][j] = "DIFF"
				// slice1[i] = append(slice1[i], "DIFF")
				// break
			}
		}
	}
	return slice1
}

func convertArrayArrayToExecl(slice [][]string, nameFile string) {
	newFileExecl := excelize.NewFile()
	newSheetName := "Sheet1"

	for rowIndex, row := range slice {
		for colIndex, colCell := range row {
			colName, err := excelize.ColumnNumberToName(colIndex + 1)
			if err != nil {
				fmt.Println("Get column name error", err)
			}
			cellName := colName + strconv.Itoa(rowIndex+1)
			newFileExecl.SetCellValue(newSheetName, cellName, colCell)
		}
	}

	err := newFileExecl.SaveAs(nameFile)
	if err != nil {
		fmt.Println("Error save file execl", err)
	}
}

func compareSlicesUpgrade(sliceA, sliceB [][]string) [][]string {
	// Create new slice [][]string with len(sliceA)
	fmt.Println(len(sliceA))
	result := make([][]string, len(sliceA))
	for i := range sliceA {
		result[i] = make([]string, len(sliceA[i]))
	}
	// Loop compare for each element in sliceA to sliceB
	for rowA := range sliceA {
		for valA := range sliceA[rowA] {
			count := 0
			for rowB := range sliceB {
				for valB := range sliceB[rowB] {
					if sliceA[rowA][valA] == sliceB[rowB][valB] {
						result[rowA][valA] = sliceB[rowB][valB]
					}
					if sliceA[rowA][valA] != sliceB[rowB][valB] {
						count++
						if count == len(sliceA)*len(sliceA[rowA]) {
							fmt.Printf("Have value is different in execl is %v\n", sliceA[rowA][valA])
						}
					}
				}
			}
		}
	}
	updateSlice(result)

	return result
}

func updateSlice(slice [][]string) [][]string {
	for row := range slice {
		for val := range slice[row] {
			if slice[row][val] == "" {
				slice[row][val] = "DIFF"
			}
		}
	}
	return slice
}

func compareSlicesV2(sliceA, sliceB [][]string) [][]string {
	// MSSV IN COLUMN 1
	onlyKey := 1
	// Create new slice [][]string with len(sliceA)
	result1 := make([][]string, len(sliceA))
	for i := range sliceA {
		result1[i] = make([]string, len(sliceA[i]))
	}
	result2 := make([][]string, len(sliceB))
	for i := range sliceA {
		result2[i] = make([]string, len(sliceB[i]))
	}

	for rowA := range sliceA {
		countRow := 0
		for rowB := range sliceB {
			// SAME MSSV
			if sliceA[rowA][onlyKey] == sliceB[rowB][onlyKey] {
				for valA := range sliceA[rowA] {
					countVal := 0
					for valB := range sliceB[rowB] {
						// ALL VALUE IN ROW MSSV IS SAME
						if sliceA[rowA][valA] == sliceB[rowB][valB] {
							result1[rowA][valA] = sliceB[rowB][valB]
							result2[rowA][valA] = sliceB[rowB][valB]
						}
						// HAVE A VALUE IN ROW MSSV IS DIFF
						if sliceA[rowA][valA] != sliceB[rowB][valB] {
							countVal++
							if countVal == len(sliceA[rowA]) {
								fmt.Printf("Value is diffrent is: %v \n", sliceA[rowA][valA])
								result1[rowA] = append(result1[rowA], "DIFF")
								result2[rowB] = append(result2[rowB], "DIFF")
							}
						}
					}
				}
			}
			// DIFF MSSV -> NEW ROW
			if sliceA[rowA][onlyKey] != sliceB[rowB][onlyKey] {
				countRow++
				if countRow == len(sliceB) {
					result1[rowA] = sliceA[rowA]
					result1[rowA] = append(result1[rowA], "NEW")

					result2[rowA] = sliceB[rowA]
					result2[rowA] = append(result2[rowA], "MISS")
				}
			}
		}

	}

	fmt.Println("TEST", result1)
	fmt.Println("TEST2", result2)

	return mergeResult(result1, result2)
}

func mergeResult(result1, result2 [][]string) [][]string {
	finalLen := len(result1) + len(result2)
	finalResult := make([][]string, finalLen)
	for i := 0; i < len(result1); i++ {
		finalResult[i] = make([]string, len(result1[i]))
		copy(finalResult[i], result1[i])
	}

	countMiss := 0
	for row := range result2 {
		if result2[row][len(result2[row])-1] == "MISS" {
			countMiss++
			finalResult[len(result1)+countMiss] = append(finalResult[len(result1)+countMiss], result2[row]...)
		}
	}

	return finalResult
}
