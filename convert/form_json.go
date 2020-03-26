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
	counterSummaries, err := dproxy.New(figure).M("counterSummary").Map()
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
	var newSummaries []interface{} = make([]interface{}, len(keys))
	for idx, key := range keys {
		newArea := make(map[string]interface{}, len(areas[key].(map[string]interface{})))
		for k, v := range areas[key].(map[string]interface{}) {
			newArea[k] = v
		}
		newSummary := make(map[string]interface{}, len(counterSummaries[key].(map[string]interface{})))
		for k, v := range counterSummaries[key].(map[string]interface{}) {
			newSummary[k] = v
		}
		newArea["id"] = key
		newAreas[idx] = newArea
		newSummary["areaId"] = key
		newSummary["areaName"] = areas[key].(map[string]interface{})["name"]
		newSummaries[idx] = newSummary
	}

	formd["areas"] = newAreas
	formd["counterSummary"] = newSummaries

	return Interface2File(outFile, formd)
}
