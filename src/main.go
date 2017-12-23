package main

import (
	"os"

	"github.com/urfave/cli"
)

func buildFunc(c *cli.Context) error {

	p := NewPackageConfig("")

	if len(c.Args()) == 0 {
		for _, s := range p.getProjects() {
			pb := NewProjectBuilder(s, p)
			pb.build()
		}
	} else {
		for _, s := range c.Args() {
			pb := NewProjectBuilder(s, p)
			pb.build()
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "CPM"
	app.Version = "0.01"
	app.EnableBashCompletion = true
	app.Copyright = "(C) 2017 Marius Messerschmidt"

	app.Commands = []cli.Command{
		{
			Category:  "Local",
			Name:      "build",
			Aliases:   []string{"b"},
			Usage:     "Builds all projects or a specified project",
			ArgsUsage: "[Project]...",
			Action:    buildFunc,
		},
	}

	app.Run(os.Args)
}
