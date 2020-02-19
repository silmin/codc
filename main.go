package main

import (
    "os"
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"

    "github.com/urfave/cli"

    types "github.com/silmin/odc-combiner/typefile"
)

func isExistFile(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func file2Figure(filename string) (types.Figure, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
        return types.Figure{}, err
    }

    var figure types.Figure
    if err := json.Unmarshal(bytes, &figure); err != nil {
        log.Fatal(err)
        return types.Figure{}, err
    }

    return figure, nil
}

func main() {
    app := &cli.App {

        Name: "ODC-combiner",
        Usage: "combine json of opendatacam",
        Version: "0.0.1",

        Commands: []*cli.Command {
            {
                Name: "areas",
                Usage: "combine any areas",

                Flags: []cli.Flag {
                    &cli.StringFlag {
                        Name: "file",
                        Aliases: []string{"f"},
                        Usage: "input json file",
                        Value: "input.json",
                    },
                },
                Action: func (context *cli.Context) error {
                    filename := context.String("file")
                    if !isExistFile(filename) {
                        fmt.Println(filename, "not exist.")
                        return nil
                    }
                    fmt.Println("filename:", filename)
                    fmt.Println("args:", context.Args())

                    file2Figure(filename)

                    //CombineAreas(filename, context.Args().Slice())

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

