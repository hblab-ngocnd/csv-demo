package csv_reder

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"get_value":        tplFuncGetValue(),
		"split_value_dash": splitValueByDash(),
	}
}
func ParserData(in, out string, jsonData string) error {
	var mapData map[string]string
	err := json.Unmarshal([]byte(jsonData), &mapData)
	if err != nil {
		log.Fatalln("failed unmarshal json data", mapData)
	}
	t, err := template.New(path.Base(in)).Funcs(templateFuncs()).ParseFiles(in)
	if err != nil {
		fmt.Println(err)
		return err
	}
	buffers := new(bytes.Buffer)
	err = t.Execute(buffers, mapData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.Create(out)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = file.Write(buffers.Bytes())
	return err
}
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
func tplFuncGetValue() interface{} {
	return func(value interface{}) string {
		return getValue(value)
	}
}
func getValue(value interface{}) string {
	valueStr := ToString(value)
	return valueStr
}
func splitValueByDash() interface{} {
	return func(value interface{}, location int) string {
		return splitByDash(value, location)
	}
}
func splitByDash(value interface{}, location int) string {
	valueStr := ToString(value)
	input := strings.ReplaceAll(valueStr, "–", "-")
	output := strings.Split(input, "-")
	if len(output) < location {
		return ""
	}

	return output[location-1]
}
func ToString(raw interface{}) string {
	switch raw.(type) {
	case int64:
		return strconv.Itoa(int(raw.(int64)))
	case int:
		return strconv.Itoa(raw.(int))
	case float32:
		return strconv.Itoa(ToInt(raw, 0))
	case float64:
		return strconv.Itoa(ToInt(raw, 0))
	case string:
		return raw.(string)
	case []string:
		if len(raw.([]string)) > 0 {
			return (raw.([]string))[0]
		}
		return ""
	default:
		return ""
	}
}
func ToInt(value interface{}, defaultValue int) int {
	switch value.(type) {
	case int:
		return value.(int)
	case int64:
		return int(value.(int64))
	case float32:
		return int(value.(float32))
	case float64:
		return int(value.(float64))
	case string:
		if result, err := strconv.Atoi(value.(string)); err == nil {
			return result
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
