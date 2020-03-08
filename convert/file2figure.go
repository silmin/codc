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
