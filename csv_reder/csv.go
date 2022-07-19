package csv_reder

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func ToUpper(in string, params ...string) (out string, err error) {
	return strings.ToUpper(in), nil
}

var mapFunc = map[string]func(in string, params ...string) (out string, err error){
	"to-upper": ToUpper,
}

func RenderData(template, out string, jsonData string) error {
	var mapData map[string]string
	err := json.Unmarshal([]byte(jsonData), &mapData)
	if err != nil {
		log.Fatalln("failed unmarshal json data", mapData)
	}
	read, err := os.Open(template)
	defer read.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	reader := csv.NewReader(read)
	reader.FieldsPerRecord = -1
	settings, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("fail to read file", err)
	}
	var results [][]string
	for _, elements := range settings {
		rows := make([]string, len(elements))
		for i, e := range elements {
			rows[i] = getValue(e, mapData)
		}
		results = append(results, rows)
	}
	write, err := os.Create(out)
	defer write.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(write)
	defer w.Flush()
	err = w.WriteAll(results)
	if err != nil {
		log.Fatalln("failed write data", err)
	}
	return nil
}
func getValue(e string, mapData map[string]string) string {
	e = strings.TrimLeft(e, " ")
	if len(e) == 0 {
		return ""
	}
	arr := strings.Split(e, " ")
	if len(arr) == 1 {
		if e[0] != '.' {
			return e
		}
		return mapData[e[1:]]
	}
	//FuncCall
	f := mapFunc[arr[0]]
	if f == nil {
		log.Printf("csv_reder: please implement %s function", arr[0])
		return ""
	}
	var v string
	var err error
	if len(arr) >= 3 {
		v, err = f(mapData[arr[1][1:]], arr[2:]...)
	} else {
		v, err = f(mapData[arr[1][1:]])
	}
	if err != nil {
		log.Println(err)
		return ""
	}
	return v
}
