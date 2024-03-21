package helpers

import (
	"fmt"

	"github.com/apolo96/metaudio/internal/interfaces"
)

func help(cmds []interfaces.Command) {
	fmt.Printf("usage: metaudio <command> [<flags>] \n \n")
	fmt.Printf("These are a few Audiofile commands: \n \n")
	for _, c := range cmds {
		fmt.Printf("%s (%s) \n", c.Name(), c.Description())
	}
}
