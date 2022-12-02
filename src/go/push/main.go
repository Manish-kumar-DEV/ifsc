package main

import (
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	file, error := excelize.OpenFile("src/68774.xlsx")
	if error != nil {
		log.Fatal(error)
	}

	m := make(map[string][]interface{})
	cols, error := file.GetCols("Sheet1")
	if error != nil {
		log.Fatal(error)
	}
	populateMap(cols, m)

	cols1, error1 := file.GetCols("Sheet2")
	if error1 != nil {
		log.Fatal(error1)
	}
	populateMap(cols1, m)

	marshal, _ := json.Marshal(m)
	ioutil.WriteFile("IFSC.json", marshal, os.ModePerm)
}

func populateMap(cols [][]string, m map[string][]interface{}) {
	for i, value := range cols[1] {
		if i == 0 {
			continue
		}
		bank := value[0:4]
		branch := value[5:]
		if v, err := strconv.Atoi(branch); err == nil {
			m[bank] = append(m[bank], v)
		} else {
			m[bank] = append(m[bank], branch)
		}
	}
}
