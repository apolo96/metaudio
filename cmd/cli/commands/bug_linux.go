//go:build linux

package commands

import "os/exec"

func openBrowser(url string) error {
	return exec.Command("xdg-open", url).Start()
}
