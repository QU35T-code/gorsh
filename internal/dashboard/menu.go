package dashboard

import (
	"github.com/desertbit/grumble"
)

func Menu() {
	var App = grumble.New(&grumble.Config{
		Name:                  "gorsh",
		Description:           "TODO",
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
	})

	App.AddCommand(&grumble.Command{
		Name:  "sessions",
		Help:  "Show all sessions",
		Usage: "sessions",
		Run: func(c *grumble.Context) error {
			sessions()
			return nil
		},
	})

	App.AddCommand(&grumble.Command{
		Name:  "help",
		Help:  "Show helper",
		Usage: "help",
		Run: func(c *grumble.Context) error {
			helper()
			return nil
		},
	})

	grumble.Main(App)
}
