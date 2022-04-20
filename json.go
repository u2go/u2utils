package u2utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func JsonLoadFromFile(file string, value any) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, value)
}

func JsonTypeConvert(in any, out any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func PrintJson(data any) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("PrintJson Error: ", err.Error())
		return
	}
	fmt.Println(string(b))
}

func PrintJsonPretty(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("PrintJsonPretty Error: ", err.Error())
		return
	}
	fmt.Println(string(b))
}
