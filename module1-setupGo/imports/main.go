// from the urfave/cli/v2 package example
// https://cli.urfave.org/v2/examples/arguments/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Hello %q", cCtx.Args().Get(0))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
