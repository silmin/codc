package main

import (
	"github.com/koron/go-dproxy"

	"github.com/silmin/codc/convert"
)

func IsFormd(filename string) (bool, error) {
	figure, err := convert.File2Interface(filename)
	if err != nil {
		return false, err
	}

	_, err = dproxy.New(figure).M("areas").Array()
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}
}
