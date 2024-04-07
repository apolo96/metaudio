package format

import "github.com/fatih/color"

const (
	EmojiCheck = "\U00002705"
)

var ErrorColor = color.New(color.BgRed, color.FgWhite).SprintFunc()
