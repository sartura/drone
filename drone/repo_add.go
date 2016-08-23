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

	fmt.Printf("\n* Easy access URL for adding a hook manually in the %s can be found below:\n", repo)
	fmt.Printf("\t https://bitbucket.org/%s/admin/addon/admin/bitbucket-webhooks/bb-webhooks-repo-admin", repo)
	fmt.Printf("\n* Add the following hook manually in the repository %s, for event `Repository push`:\n\t%s\n\n", r.FullName, r.HookURI)
	return nil
}
