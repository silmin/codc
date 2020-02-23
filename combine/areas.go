package combine

import (
    "strings"

    types "github.com/silmin/codc/typefile"
)

func Area(figure types.Figure, names []string) (types.Figure, error) {
    newArea := types.Areas {
        Name:   strings.Join(names, "-"),
        Source: names,
    }
    figure.Areas = append(figure.Areas, newArea)

    newSummary, err := combineSummary(figure.CounterSummary, names)
    if err != nil {
        return types.Figure{}, err
    }
    figure.CounterSummary = append(figure.CounterSummary, newSummary)

    return figure, nil
}
