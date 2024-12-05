package main

import (
	"context"
	"fmt"
	"log"
	"lukaszwisniewski88/aoc2024/fifth"
	"lukaszwisniewski88/aoc2024/first"
	"lukaszwisniewski88/aoc2024/fourth"
	"lukaszwisniewski88/aoc2024/second"
	"lukaszwisniewski88/aoc2024/third"
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
			{
				Name:    "second",
				Aliases: []string{"2"},
				Usage:   "Second day of Advent of Code, supply the input file path",
				Action: func(ctx context.Context, c *cli.Command) error {
					sum, err := second.ProcessDayTwo(c.Args().First())
					if err != nil {
						return err
					}
					fmt.Println(sum)

					return nil
				},
			},
			{
				Name:    "third",
				Aliases: []string{"3"},
				Usage:   "third day of Advent of Code, supply the input file path",
				Action: func(ctx context.Context, c *cli.Command) error {
					sum, err := third.ProcessDayThree(c.Args().First())
					if err != nil {
						return err
					}
					fmt.Println(sum)

					return nil
				},
			},
			{
				Name:    "fourth",
				Aliases: []string{"4"},
				Usage:   "fourth day of Advent of Code, supply the input file path",
				Action: func(ctx context.Context, c *cli.Command) error {
					sum, err := fourth.ProcessDayFour(c.Args().First())
					if err != nil {
						return err
					}
					fmt.Println(sum)

					return nil
				},
			},
			{
				Name:    "fifth",
				Aliases: []string{"5"},
				Usage:   "fifth day of Advent of Code, supply the input file path",
				Action: func(ctx context.Context, c *cli.Command) error {
					sum, err := fifth.ProcessDayFive(c.Args().First())
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
