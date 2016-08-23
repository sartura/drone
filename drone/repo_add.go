package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

var repoAddCmd = cli.Command{
	Name:  "add",
	Usage: "add a repository",
	Action: func(c *cli.Context) {
		if err := repoAdd(c); err != nil {
			log.Fatalln(err)
		}
	},
}

func repoAdd(c *cli.Context) error {
	repo := c.Args().First()
	owner, name, err := parseRepo(repo)
	if err != nil {
		return err
	}

	client, err := newClient(c)
	if err != nil {
		return err
	}

	r, err := client.RepoPost(owner, name);
	if err != nil {
		return err
	}

	fmt.Printf("\n\nPlease add the hook manually in the repository %s, for event `Repository push`:\n%s\n\n", r.FullName, r.HookURI)
	fmt.Printf("Successfully activated repository %s/%s\n", owner, name)
	return nil
}
