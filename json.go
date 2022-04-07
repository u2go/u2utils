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
