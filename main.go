package main

import (
    "os"
    "fmt"
    "log"

    "github.com/urfave/cli"
)


func main() {
    app := &cli.App {

        Name: "test",
        Usage: "test CLI tool",
        Version: "0.0.1",

        Commands: []*cli.Command {
            {
                Name: "hoge",
                Usage: "say hoge",

                Flags: []cli.Flag {
                    &cli.BoolFlag {
                        Name: "cat",
                        Aliases: []string{"c"},
                        Usage: "Echo with cat",
                    },
                },

                Action: func (context *cli.Context) error {
                    fmt.Println("hoge")
                    if context.Bool("cat") { fmt.Println("nyan !") }
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

