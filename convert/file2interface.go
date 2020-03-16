package convert

import (
	"encoding/json"
	"io/ioutil"
)

func File2Interface(filename string) (interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return bytes, err
	}

	var figure interface{}
	err = json.Unmarshal([]byte(bytes), &figure)
	if err != nil {
		return figure, err
	}

	return figure, nil
}
