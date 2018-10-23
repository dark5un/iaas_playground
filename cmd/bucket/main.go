package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	IAAS       string
	BucketName string
)

func main() {

	app := cli.NewApp()
	app.Usage = "Multi IAAS bucket handling"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "iaas",
			Destination: &IAAS,
			EnvVar:      "IAAS",
			Usage:       "IAAS choice of [gcp, aws, azure]",
		},
		cli.StringFlag{
			Name:        "name",
			Destination: &BucketName,
			EnvVar:      "BUCKET_NAME",
			Usage:       "of the bucket",
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
	app.Before = Before
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
