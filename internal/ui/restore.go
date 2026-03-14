package ui

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/harshalve/ctx-tool/internal/context"
)

func RestoreProjectData(ctx context.ProjectContext) {
	fmt.Printf("\n🚀 Restoring Context: %s\n", ctx.Name)
	fmt.Printf("🌿 Branch:    %s\n", ctx.Branch)
	fmt.Printf("📍 Directory: %s\n", ctx.Directory)
	fmt.Println("------------------------------------------")

	// 1. Open Browser Links
	if len(ctx.Links) > 0 {
		fmt.Println("\n🔗 Opening Links:")
		for _, link := range ctx.Links {
			fmt.Printf("  - Opening: %s\n", link.Value)
			if link.Notes != "" {
				fmt.Printf("    📝 Note: %s\n", link.Notes)
			}
			openBrowser(link.Value)
		}
	}

	// 2. Display Terminal Context
	if len(ctx.Terminals) > 0 {
		fmt.Println("\n💻 Terminal Context & Tasks:")
		for _, term := range ctx.Terminals {
			fmt.Printf("  - Command: %s\n", term.Value)
			if term.Notes != "" {
				fmt.Printf("    📝 Note: %s\n", term.Notes)
			}
		}
	}
	fmt.Println("\n✅ Welcome back! Your context is rehydrated.")
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
