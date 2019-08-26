// examp project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Masstrct struct {
	Values []Somestrct3 `json:"values,omitempty"`
}
type Somestrct struct {
	Id     int          `json:"id,omitempty"`
	Title  string       `json:"title,omitempty"`
	Value  interface{}  `json:"value,omitempty"`
	Values []Somestrct2 `json:"values,omitempty"`
}

type Somestrct2 struct {
	Id     int         `json:"id,omitempty"`
	Title  string      `json:"title,omitempty"`
	Params []Somestrct `json:"params,omitempty"`
}

type Somestrct3 struct {
	Id    int         `json:"id,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type MessageError struct {
	Message string `json:"message"`
}
type Err struct {
	Error MessageError `json:"error"`
}

func main() {

	plan, err := ioutil.ReadFile("TestcaseStructure.json")

	if err != nil {
		Errors()
		return
	}
	var data Somestrct2
	err = json.Unmarshal(plan, &data)
	if err != nil {
		Errors()
		return
	}
	plan, err = ioutil.ReadFile("Values.json")
	if err != nil {
		Errors()
		return
	}
	var data2 Masstrct
	err = json.Unmarshal(plan, &data2)
	if err != nil {
		Errors()
		return
	}

	var data3 []Somestrct3
	data3 = data2.Values
	rec(&data, &data3)

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error json")
		return
	}
	err = ioutil.WriteFile("StructureWithValues.json", file, 0644)
	if err != nil {
		fmt.Println("Error writefile")
	}

}

func rec(str *Somestrct2, str2 *[]Somestrct3) {
	for i, val := range str.Params {
		c(&str.Params[i], str2)
		for _, val2 := range val.Values {
			rec(&val2, str2)
		}
	}

}

func c(s *Somestrct, str2 *[]Somestrct3) {
	x := s
	for _, el := range *str2 {
		if el.Id == x.Id {
			x.Value = el.Value
			break
		}
	}
}

func Errors() {
	m := MessageError{"Входные файлы некорректны"}
	mr := Err{m}
	erfile, err := json.MarshalIndent(mr, "", " ")
	if err != nil {
		fmt.Println("Error json")
		return
	}
	err = ioutil.WriteFile("error.json", erfile, 0644)
	if err != nil {
		fmt.Println("Error write file")
		return
	}
}
