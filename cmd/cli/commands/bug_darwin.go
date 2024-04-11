//go:build darwin

package commands

import "os/exec"

func openBrowser(url string) error {
	return exec.Command("open", url).Start()
}
