package main

import (
    "os"
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"

    "github.com/urfave/cli"

    types "github.com/silmin/odc-combiner/typefile"
    "github.com/silmin/odc-combiner/combine"
)

func isExistFile(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func file2Figure(filename string) (types.Figure, error) {
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

func figure2file(filename string, figure types.Figure) error {
    bytes, err := json.MarshalIndent(figure, "", "    ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filename, bytes, 0664)
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
    app := &cli.App {

        Name: "ODC-combiner",
        Usage: "combine json of opendatacam",
        Version: "0.1.0",

        Commands: []*cli.Command {
            {
                Name: "areas",
                Usage: "combine any areas",

                Flags: []cli.Flag {
                    &cli.StringFlag {
                        Name: "src",
                        Aliases: []string{"s"},
                        Usage: "input json file",
                        Value: "input.json",
                    },
                    &cli.StringFlag {
                        Name: "dst",
                        Aliases: []string{"d"},
                        Usage: "output json file",
                        Value: "output.json",
                    },
                },
                Action: func (context *cli.Context) error {
                    inFile := context.String("src")
                    outFile := context.String("dst")
                    if !isExistFile(inFile) {
                        fmt.Println(inFile, "not exist.")
                        return nil
                    }
                    args := context.Args().Slice()
                    fmt.Println("input:", inFile)
                    fmt.Println("args:", args)

                    figure, err := file2Figure(inFile)
                    if err != nil {
                        log.Fatal(err)
                        return err
                    }

                    if !isExistAreas(figure, args) {
                        fmt.Println("contains something that doesn't exist in", args)
                        return nil
                    }

                    figure, err = combine.Area(figure, args)
                    if err != nil {
                        log.Fatal(err)
                        return err
                    }

                    if err := figure2file(outFile, figure); err != nil {
                        log.Fatal(err)
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

