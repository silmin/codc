package convert

import (
	"encoding/json"
	"io/ioutil"

	types "github.com/silmin/codc/typefile"
)

func Figure2File(filename string, figure types.Figure) error {
	bytes, err := json.MarshalIndent(figure, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}
