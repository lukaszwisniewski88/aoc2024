package main

import (
	"context"
	"fmt"
	"log"
	"lukaszwisniewski88/aoc2024/first"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "first",
				Aliases: []string{"1"},
				Usage:   "First day of Advent of Code, supply the input file path",
				Action: func(ctx context.Context, c *cli.Command) error {
					sum, err := first.ProcessDayOne(c.Args().First())
					if err != nil {
						return err
					}
					fmt.Println(sum)

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}