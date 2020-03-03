package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/silmin/codc/combine"
	"github.com/silmin/codc/convert"
	types "github.com/silmin/codc/typefile"
)

func isExistFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func isExistAreas(figure types.Figure, names []string) bool {
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

func main() {
	app := &cli.App{

		Name:    "codc | Combine OpenDataCam's data",
		Usage:   "combine json of opendatacam",
		Version: "0.3.0",

		Commands: []*cli.Command{
			{
				Name:  "form",
				Usage: "form default json to original formmat",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "src",
						Aliases: []string{"s"},
						Usage:   "input json file",
						Value:   "input.json",
					},
					&cli.StringFlag{
						Name:    "dst",
						Aliases: []string{"d"},
						Usage:   "output json file",
						Value:   "formd.json",
					},
				},
				Action: func(context *cli.Context) error {
					inFile := context.String("src")
					outFile := context.String("dst")
					if !isExistFile(inFile) {
						fmt.Println(inFile, "not exist.")
						return nil
					}
					fmt.Println("input:", inFile)

					fmt.Println("output:", outFile)

					return nil
				},
			},
			{
				Name:  "areas",
				Usage: "combine any areas",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "src",
						Aliases: []string{"s"},
						Usage:   "input json file",
						Value:   "formd.json",
					},
					&cli.StringFlag{
						Name:    "dst",
						Aliases: []string{"d"},
						Usage:   "output json file",
						Value:   "output.json",
					},
				},
				Action: func(context *cli.Context) error {
					inFile := context.String("src")
					outFile := context.String("dst")
					if !isExistFile(inFile) {
						fmt.Println(inFile, "not exist.")
						return nil
					}
					args := context.Args().Slice()
					fmt.Println("input:", inFile)
					fmt.Println("args:", args)

					figure, err := convert.File2Figure(inFile)
					if err != nil {
						return err
					}

					if !isExistAreas(figure, args) {
						fmt.Println("contains something that doesn't exist in", args)
						return nil
					}

					figure, err = combine.Area(figure, args)
					if err != nil {
						return err
					}

					if err := convert.Figure2File(outFile, figure); err != nil {
						return err
					}
					fmt.Println("output:", outFile)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
