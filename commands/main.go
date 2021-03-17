package commands

import (
	"github.com/mix-go/cli"
)

var Commands = []*cli.Command{
	{
		Name:  "hello",
		Usage: "\tEcho demo",
		Options: []*cli.Option{
			{
				Names: []string{"n", "name"},
				Usage: "Your name",
			},
			{
				Names: []string{"say"},
				Usage: "\tSay ...",
			},
		},
		RunI: &HelloCommand{},
	},
}
