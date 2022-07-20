package csv_reder

import (
	"bytes"
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
