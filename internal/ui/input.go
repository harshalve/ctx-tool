package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/harshalve/ctx-tool/internal/context"
)

var reader = bufio.NewReader(os.Stdin)

func CollectProjectData(name string) context.ProjectContext {
	ctx := context.ProjectContext{
		Name:      name,
		CreatedAt: time.Now(),
	}

	// --- Links Loop ---
	fmt.Println("\n--- Step 1: Add Links (Jira, GitHub, Docs) ---")
	fmt.Println("(Press Enter on an empty line to move to Terminals)")
	ctx.Links = collectResources("Link URL")

	// --- Terminals Loop ---
	fmt.Println("\n--- Step 2: Add Terminal Context (Commands/Tasks) ---")
	fmt.Println("(Press Enter on an empty line to finish)")
	ctx.Terminals = collectResources("Terminal Task/Command")

	return ctx
}

func collectResources(promptText string) []context.Resource {
	var resources []context.Resource
	for {
		fmt.Printf("%s: ", promptText)

		// 1. Read the input
		val, _ := reader.ReadString('\n')

		// 2. Remove the newline/spaces (CRITICAL)
		val = strings.TrimSpace(val)

		// 3. Check if empty BEFORE asking for notes
		if val == "" {
			break // This exits the loop and moves to the next section
		}

		fmt.Printf("Notes for '%s': ", val)
		note, _ := reader.ReadString('\n')
		note = strings.TrimSpace(note)

		resources = append(resources, context.Resource{Value: val, Notes: note})
	}
	return resources
}
