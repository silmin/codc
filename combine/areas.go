package combine

import (
    "strings"

    types "github.com/silmin/odc-combiner/typefile"
)

func Area(figure types.Figure, names []string) (types.Figure, error) {
    newArea := types.Areas {
        Name:   strings.Join(names, "-"),
        Source: names,
    }

    figure.Areas = append(figure.Areas, newArea)

    return figure, nil
}
