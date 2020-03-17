package convert

import (
	"github.com/koron/go-dproxy"
)

func getKeys(m map[string]interface{}) []string {
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func FormJson(inFile string, outFile string) error {

	figure, err := File2Interface(inFile)

	areas, err := dproxy.New(figure).M("areas").Map()
	if err != nil {
		return err
	}
	keys := getKeys(areas)

	formd, err := dproxy.New(figure).Map()
	if err != nil {
		return err
	}

	// そのままだと型周りで動けないので新しく作った
	var newAreas []interface{} = make([]interface{}, len(keys))
	for idx, key := range keys {
		newArea := make(map[string]interface{}, len(areas[key].(map[string]interface{})))
		for k, v := range areas[key].(map[string]interface{}) {
			newArea[k] = v
		}
		newArea["id"] = key
		newAreas[idx] = newArea
	}

	formd["areas"] = newAreas

	return Interface2File(outFile, formd)
}
