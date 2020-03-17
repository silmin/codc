package convert

import (
	"encoding/json"
	"io/ioutil"
)

func Interface2File(filename string, outObj interface{}) error {
	bytes, err := json.MarshalIndent(outObj, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}
