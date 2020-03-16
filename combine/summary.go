package combine

import (
    "strings"
    "errors"

    types "github.com/silmin/codc/typefile"
)

func getSummary(summaries []types.CounterSummary, name string) (types.CounterSummary, error) {
    for _, summary := range summaries {
        if name == summary.AreaName {
            return summary, nil
        }
    }
    return types.CounterSummary{}, errors.New(name + "does not exist.")
}

func combineSummary(summaries []types.CounterSummary, names []string) (types.CounterSummary, error) {
    newSummary := types.CounterSummary {}
    for _, name := range names {
        summary, err := getSummary(summaries, name)
        if err != nil {
            return types.CounterSummary{}, err
        }
        newSummary.Total += summary.Total
        newSummary.Car += summary.Car
        newSummary.Bus += summary.Bus
        newSummary.Truck += summary.Truck
        newSummary.Person += summary.Person
    }
    newSummary.AreaName = strings.Join(names, "-")

    return newSummary, nil
}

