package helpers

import (
	"fmt"

	"github.com/apolo96/metaudio/internal/interfaces"
)

type Parser struct {
	commands []interfaces.Command
}

func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string) error  {
	if len(args) < 1{
		help(p.commands)
		return nil
	}
	subcommand := args[0]
	for _, cmd := range p.commands{
		if cmd.Name() == subcommand{
			if err := cmd.ParseFlags(args[1:]); err != nil{
				return err
			}
			return cmd.Run()
		}
	}
	defer help(p.commands)
	return fmt.Errorf("error unknown subcommand: %s", subcommand)
}