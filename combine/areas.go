package combine

import (
    "strings"

    types "github.com/silmin/odc-combiner/typefile"
)

func Area(figure types.Figure, params []string) (types.Figure, error) {

    newArea := types.Areas{
        Name:   strings.Join(params, "-"),
        Source: params,
    }

    figure.Areas = append(figure.Areas, newArea)

    return figure, nil
}
