package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var (
		iaas string
	)
	app := cli.NewApp()
	app.Usage = "Multi IAAS bucket handling"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "iaas",
			Destination: &iaas,
			EnvVar:      "IAAS",
			Usage:       "IAAS choice of [gcp, aws, azure]",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "create",
			ShortName:   "c",
			Usage:       "a bucket in the target IAAS",
			Description: "create a bucket in the target IAAS",
			Action:      Create,
		},
		cli.Command{
			Name:        "delete",
			ShortName:   "d",
			Usage:       "a bucket in the target IAAS",
			Description: "delete a bucket in the target IAAS",
			Action:      Delete,
		},
		cli.Command{
			Name:        "list",
			ShortName:   "l",
			Usage:       "all buckets in the target IAAS",
			Description: "list all buckets in the target IAAS",
			Action:      List,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
