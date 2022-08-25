// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//data
	csvData := []byte(`,
33,33,2
[kanri]
012,0301429102,JR中央線（中央総武線）およびJR山手線（埼京線他）　武蔵境駅〜大崎駅
haha,`)

	err := os.WriteFile("data.csv", csvData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fCsv, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fCsv.Close()
	//Read roun 1
	csvReader := csv.NewReader(fCsv)
	csvReader.FieldsPerRecord = -1
	dataCsv, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(dataCsv)
	//Read roun 2
	fCsv.Seek(0, io.SeekStart)
	csvReader2 := csv.NewReader(fCsv)
	csvReader2.FieldsPerRecord = -1
	dataCsv2, err := csvReader2.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(dataCsv2)
	//template
	template := []byte(`company__corporateNumber,{{.col2_2}}-{{.col2_3}}
em_street,{{.col4_2}}
em_building_kana,{{.col4_3}}
company__street,{{get_value .col5_1}}
con_cat,{{con_cat .col5_1 .col4_3 .col4_2}}`)

	err = os.WriteFile("template.csv", template, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	os.Remove("template.csv")
	os.Remove("data.csv")
	//t, err := template.New(path.Base(pathTemplate)).Option("missingkey=zero").Funcs(template.FuncMap(templateFuncs())).ParseFiles(pathTemplate)
	//fmt.Println(result)
}
