package exists

import (
	"os"

	types "github.com/silmin/codc/typefile"
)

func File(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Area(figure types.Figure, names []string) bool {
	for _, name := range names {
		flg := false
		for _, area := range figure.Areas {
			if name == area.Name {
				flg = true
				break
			}
		}
		if !flg {
			return false
		}
	}
	return true
}
