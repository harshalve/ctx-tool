package main

import (
	"fmt"
	"os"

	"github.com/harshalve/ctx-tool/internal/config"
	"github.com/harshalve/ctx-tool/internal/context"
	"github.com/harshalve/ctx-tool/internal/storage"
	"github.com/harshalve/ctx-tool/internal/ui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ctx <save|restore> <project_name>")
		os.Exit(1)
	}
	projectName := ""
	if len(os.Args) > 2 {
		projectName = os.Args[2]
	}

	command := os.Args[1]

	storagePath := config.GetStorageDir()
	store := &context.Store{StorageDir: storagePath}

	newStorage, err := storage.NewStorage()
	if err != nil {
		fmt.Printf("❌ Storage error: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "list":
		ui.HandleList(newStorage)
	case "save":
		executeSave(projectName, store)
	case "restore":
		ctx, err := store.Load(projectName)
		if err != nil {
			fmt.Printf("❌ Error: Could not find context '%s'\n", projectName)
			return
		}
		ui.RestoreProjectData(ctx)
	case "path":
		ctx, err := store.Load(projectName)
		if err != nil {
			fmt.Println(".")
			return
		}
		fmt.Print(ctx.Directory)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func executeSave(name string, store *context.Store) {
	fmt.Printf("Starting context snapshot for: %s\n", name)

	ctx := ui.CollectProjectData(name)

	err := store.Save(ctx)
	if err != nil {
		fmt.Printf("Error saving context: %v\n", err)
		return
	}
	fmt.Printf("✅ Context saved to %s/%s.json\n", store.StorageDir, name)
}
