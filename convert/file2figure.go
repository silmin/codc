package convert

import (
	"encoding/json"
	"io/ioutil"

	types "github.com/silmin/codc/typefile"
)

func File2Figure(filename string) (types.Figure, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return types.Figure{}, err
	}

	var figure types.Figure
	if err := json.Unmarshal(bytes, &figure); err != nil {
		return types.Figure{}, err
	}

	return figure, nil
}

func Figure2File(filename string, figure types.Figure) error {
	bytes, err := json.MarshalIndent(figure, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0664)
}
