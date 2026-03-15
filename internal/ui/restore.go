package ui

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/harshalve/ctx-tool/internal/context"
)

func RestoreProjectData(ctx context.ProjectContext) {
	fmt.Printf("\n🚀 Restoring Context: %s\n", ctx.Name)
	fmt.Println("------------------------------------------")
	fmt.Printf("📍 Directory:  file://%s\n", ctx.Directory)
	fmt.Println("------------------------------------------")

	// 1. System/Git Setup
	handleBranchSwitch(ctx.Branch, ctx.Directory)

	// 2. Browser Links
	restoreLinks(ctx.Links)

	// 3. Terminal/Task Context
	displayTerminalContext(ctx.Terminals)

	fmt.Println("\n✅ Welcome back! Your context is rehydrated.")
}

func handleBranchSwitch(branch, dir string) {
	if branch == "" || branch == "none" || branch == "n/a (not a git repo)" {
		return
	}

	fmt.Printf("🌿 Target Branch: %s\n", branch)
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		fmt.Printf("⚠️  Could not auto-switch branch (check for uncommitted changes)\n")
	} else {
		fmt.Printf("✅ Switched to branch: %s\n", branch)
	}
}

func restoreLinks(links []context.Resource) {
	if len(links) == 0 {
		return
	}
	fmt.Println("\n🔗 Opening Links:")
	for _, link := range links {
		fmt.Printf("  - %s\n", link.Value)
		if link.Notes != "" {
			fmt.Printf("    📝 %s\n", link.Notes)
		}
		openBrowser(link.Value)
	}
}

func displayTerminalContext(terms []context.Resource) {
	if len(terms) == 0 {
		return
	}
	fmt.Println("\n💻 Terminal Context & Tasks:")
	for _, term := range terms {
		fmt.Printf("  - Command: %s\n", term.Value)
		if term.Notes != "" {
			fmt.Printf("    📝 %s\n", term.Notes)
		}
	}
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "rundll32"
		args = append(args, "url.dll,FileProtocolHandler")
	default: // linux
		cmd = "xdg-open"
	}
	args = append(args, url)
	_ = exec.Command(cmd, args...).Start()
}
