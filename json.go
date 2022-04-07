package u2utils

import (
	"encoding/json"
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
