package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/xuri/excelize/v2"
)

func main() {

	mapDataExecl1 := readExecl("dstn.xlsx")
	mapDataExecl2 := readExecl("dstn2.xlsx")

	if checkBool(mapDataExecl1, mapDataExecl2) == true {
		fmt.Println("The content of execl file are the same")
	}

	if checkBool(mapDataExecl1, mapDataExecl2) == false {
		fmt.Println("The content of execl file aren't the same")
		compareMap(mapDataExecl1, mapDataExecl2)
	}
}

func readExecl(nameFile string) []map[string]interface{} {
	fileExecl, err := excelize.OpenFile(nameFile)
	if err != nil {
		log.Printf("Read File Err %v", err)
	}
	// Sheet 1 == 0
	nameSheet := fileExecl.GetSheetName(1)
	if nameSheet == "" {
		fmt.Println("Sheet index invaild")
	}

	rowsFileExecl, err := fileExecl.GetRows(nameSheet)
	if err != nil {
		log.Printf("Read File Err %v", err)
	}

	// var result []map[string]interface{}
	// headers := rowsFileExecl[0]
	// fmt.Println("TEST", headers)
	// for i := 1; i < len(rowsFileExecl); i++ {
	// 	row := rowsFileExecl[i]
	// 	rowMap := make(map[string]interface{})
	// 	for j := 0; j < len(headers); j++ {
	// 		rowMap[headers[j]] = row[j]
	// 	}
	// 	result = append(result, rowMap)
	// }

	dataExecl := make([]map[string]interface{}, 0)
	for rowIndex, row := range rowsFileExecl {
		if rowIndex == 0 {
			continue
		}
		rowData := make(map[string]interface{})
		// ColumnNumberToName -> Input int and output string name col, err
		for colIndex, colCell := range row {
			colName, err := excelize.ColumnNumberToName(colIndex + 1) // A = 1
			if err != nil {
				fmt.Println(err)
				// return
			}
			rowData[colName] = colCell
		}
		dataExecl = append(dataExecl, rowData)
	}

	return dataExecl
}

// Output true or false
func checkBool(check1, check2 []map[string]interface{}) bool {
	if reflect.DeepEqual(check1, check2) {
		return true
	}
	return false
}

// If false -> run to function compare between two map input
func compareMap(compare1, compare2 []map[string]interface{}) {
	for key, value := range compare1 {
		if !reflect.DeepEqual(value, compare2[key]) {
			fmt.Printf("Row %v is different: %v vs %v\n", key, value, compare2[key])
		}
	}
}

// // Output map[key:value]
// func convertSliceToMap(slice []interface{}) map[string]interface{} {
// 	result := make(map[string]interface{})
// 	for _, item := range slice {
// 		pair := item.(map[string]interface{})
// 		for key, value := range pair {
// 			result[key] = value
// 		}
// 		fmt.Println(result)
// 	}
// 	return result
// }

// // Output map[map[key:value],map[key:value]]
// func convertSliceToArrayMap(slice []interface{}) []map[string]interface{} {
// 	var maps []map[string]interface{}
// 	for _, item := range slice {
// 		if m, ok := item.(map[string]interface{}); ok {
// 			maps = append(maps, m)
// 		}
// 	}
// 	return maps
// }

// // Switch Value between two array map
// func switchMaps(map1, map2 []map[string]interface{}) {
// 	for keyMap := range map1 {
// 		for valueMap := range map1[keyMap] {
// 			temp := map1[keyMap][valueMap]
// 			map1[keyMap][valueMap] = map2[keyMap][valueMap]
// 			map2[keyMap][valueMap] = temp
// 		}
// 	}
// 	fmt.Println("Switch value", map1)
// 	fmt.Println("Switch value", map2)
// }

// func convertArrayMapToJson(arrayMap []map[string]interface{}) {
// 	jsonConvert, err := json.Marshal(arrayMap)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Type %T\n", jsonConvert)
// 	fmt.Println(string(jsonConvert))
// }

// func compareSlices(slice1, slice2 []interface{}) {
// 	map1 := make(map[string]interface{})
// 	map2 := make(map[string]interface{})

// 	for i, val := range slice1 {
// 		map1[strconv.Itoa(i)] = val
// 	}

// 	for i, val := range slice2 {
// 		map2[strconv.Itoa(i)] = val
// 	}

// 	for key, value := range map1 {
// 		if !reflect.DeepEqual(value, map2[key]) {
// 			fmt.Printf("Element %v is different: %v vs %v\n", key, value, map2[key])
// 		}
// 	}
// }
