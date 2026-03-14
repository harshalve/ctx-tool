package context

import (
	"os"
	"os/exec"
	"strings"
)

func GetCurrentState() (string, string) {
	cwd, _ := os.Getwd()

	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	branchStr := "none"
	if err == nil {
		branchStr = strings.TrimSpace(string(branch))
	}

	return cwd, branchStr
}
