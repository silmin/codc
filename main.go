package main

import (
    "os"
    "fmt"
    "log"

    "github.com/urfave/cli"
)


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
                        Value: "",
                    },
                },
                Action: func (context *cli.Context) error {
                    filename := context.String("file")
                    if filename == "" { return nil }
                    fmt.Println("filename:", filename)
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

