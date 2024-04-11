//go:build windows

package commands

import "os/exec"

func openBrowser(url string) error {
	return exec.Command("cmd", "/C", "start", "msedge", url).Start()
}
