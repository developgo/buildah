package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	imagesFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "display only container IDs",
		},
		cli.BoolFlag{
			Name:  "noheading, n",
			Usage: "do not print column headings",
		},
	}
	imagesDescription = "Lists locally stored images."
	imagesCommand     = cli.Command{
		Name:        "images",
		Usage:       "List images in local storage",
		Description: imagesDescription,
		Flags:       imagesFlags,
		Action:      imagesCmd,
		ArgsUsage:   " ",
	}
)

func imagesCmd(c *cli.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	quiet := false
	if c.IsSet("quiet") {
		quiet = c.Bool("quiet")
	}
	noheading := false
	if c.IsSet("noheading") {
		noheading = c.Bool("noheading")
	}

	images, err := store.Images()
	if err != nil {
		return fmt.Errorf("error reading images: %v", err)
	}

	if len(images) > 0 && !noheading && !quiet {
		fmt.Printf("%-64s %s\n", "IMAGE ID", "IMAGE NAME")
	}
	for _, image := range images {
		if quiet {
			fmt.Printf("%s\n", image.ID)
		} else {
			names := []string{""}
			if len(image.Names) > 0 {
				names = image.Names
			}
			for _, name := range names {
				fmt.Printf("%-64s %s\n", image.ID, name)
			}
		}
	}

	return nil
}
