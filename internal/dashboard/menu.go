package dashboard

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func commands_suggest(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "help", Description: "TODO"},
		{Text: "sessions", Description: "TODO"},
		{Text: "exit", Description: "TODO"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func commands_parser(command string) {
	// TODO
	// Not the best way I think !
	switch command {
	case "help":
		helper()
	case "sessions":
		sessions()
	case "exit":
		exit()
	default:
		helper()
	}
}

func Menu() {
	hostname := "Gorsh"
	prompt_name := fmt.Sprintf("[%s]> ", hostname)
	command := prompt.Input(prompt_name, commands_suggest)
	commands_parser(command)
	Menu()

	// TODO
	// Add CTRL + C leave menu on trigger
}
