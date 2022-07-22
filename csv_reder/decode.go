package csv_reder

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

func DecodeData(filePath, tmlPath string) (map[string]string, error) {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// remember to close the file at the end of the program
	defer f.Close()
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	mapCSVData := make(map[string]string)
	for i, cols := range data {
		for j, col := range cols {
			key := fmt.Sprintf("col%d_%d", i+1, j+1)
			mapCSVData[key] = col
		}
	}

	t, err := template.New(path.Base(tmlPath)).Option("missingkey=zero").Funcs(templateDecodeFuncs()).ParseFiles(tmlPath)
	if err != nil {
		return nil, err
	}
	buffers := new(bytes.Buffer)
	err = t.Execute(buffers, mapCSVData)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(buffers.Bytes())
	csvReader = csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1
	data, err = csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, len(data))
	for _, cols := range data {
		if len(cols) >= 2 {
			key := cols[0]
			result[key] = strings.Join(cols[1:], "")
		}
	}
	return result, nil
}
func templateDecodeFuncs() template.FuncMap {
	return template.FuncMap{
		"get_value": tplDecodeFuncGetValue(),
		"con_cat":   tplDecodeFuncConcat(),
	}
}
func tplDecodeFuncGetValue() interface{} {
	return func(value interface{}) string {
		if value == nil {
			return ""
		}
		return strings.Repeat(value.(string), 2)
	}
}
func tplDecodeFuncConcat() interface{} {
	return func(value ...interface{}) string {
		arr := make([]string, len(value))
		for i, v := range value {
			arr[i] = v.(string)
		}
		return strings.Join(arr, "-")
	}
}
